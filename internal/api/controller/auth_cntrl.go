package controller

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var (
	// CookieExpiration is expiration for JWT cookie
	CookieExpiration time.Duration = 72 * time.Hour

	StateExpiration time.Duration = 20 * time.Minute
)

type (
	// AuthCntrl is controller to handle authentication
	AuthCntrl struct {
		dig.In
		*infra.Auth
		Svc service.AuthSvc
	}
)

// Login with google auth
func (c *AuthCntrl) Login(ce echo.Context) error {
	loginRes, err := c.Svc.Login()
	if err != nil {
		return err
	}
	ce.SetCookie(&http.Cookie{
		Name:     "oauthstate",
		Value:    loginRes.State,
		Expires:  time.Now().Add(StateExpiration),
		HttpOnly: true,
		Secure:   c.Auth.CookieSecure,
	})
	return ce.Redirect(http.StatusTemporaryRedirect, loginRes.Redirect)
}

// Callback for google auth
func (c *AuthCntrl) Callback(ce echo.Context) (err error) {
	ctx := ce.Request().Context()

	oauthState, _ := ce.Cookie("oauthstate")

	callbackRes, err := c.Svc.Callback(ctx, &service.CallbackRequest{
		OAuthState:      oauthState,
		StateParam:      ce.QueryParam("state"),
		CodeParam:       ce.QueryParam("code"),
		TokenExpiration: CookieExpiration,
	})
	if err != nil {
		errMsg := err.Error()
		if strings.HasPrefix(errMsg, "callback-failed:") {
			log.Error(errMsg)
			u, _ := url.Parse(c.RedirectFailure)
			u.Query().Set("oauth_error", "true")
			return ce.Redirect(http.StatusTemporaryRedirect, u.String())
		}
		return err
	}

	c.setCookie(ce, callbackRes.JWTToken, c.Auth.CookieSecure)
	u, _ := url.Parse(c.RedirectSuccess)
	return ce.Redirect(http.StatusTemporaryRedirect, u.String())
}

// Logout by invalidating cookies
func (c *AuthCntrl) Logout(ce echo.Context) (err error) {
	cleanCookie(ce)
	return ce.Redirect(http.StatusSeeOther, c.LogoutRedirect)
}

func cleanCookie(ce echo.Context) {
	ce.SetCookie(&http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"})
	ce.SetCookie(&http.Cookie{Name: "token", MaxAge: -1, Path: "/"})
}

func (c *AuthCntrl) setCookie(ce echo.Context, jwtToken string, secure bool) {
	ce.SetCookie(&http.Cookie{
		Name:     "secure_token",
		Value:    jwtToken,
		Expires:  time.Now().Add(CookieExpiration),
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
	})

	ce.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Expires:  time.Now().Add(CookieExpiration),
		Path:     "/",
		HttpOnly: false,
		Secure:   secure,
	})
}
