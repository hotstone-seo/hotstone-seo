package controller

import (
	"fmt"
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
	result, err := c.Svc.Login()
	if err != nil {
		return err
	}
	ce.SetCookie(result.Cookie)
	return ce.Redirect(http.StatusTemporaryRedirect, result.Redirect)
}

// Callback for google auth
func (c *AuthCntrl) Callback(ce echo.Context) (err error) {
	ctx := ce.Request().Context()

	failureURL, err := urlWithQueryParams(c.RedirectFailure, url.Values{"oauth_error": {"true"}})
	if err != nil {
		return fmt.Errorf("AuthCallback: %s", err.Error())
	}

	oauthState, _ := ce.Cookie("oauthstate")

	callbackRes, err := c.Svc.Callback(ctx, &service.CallbackRequest{
		OAuthState: oauthState,
		StateParam: ce.QueryParam("state"),
		CodeParam:  ce.QueryParam("code"),
	})
	if err != nil {
		errMsg := err.Error()
		if strings.HasPrefix(errMsg, "callback-failed:") {
			log.Error(errMsg)
			return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
		}
		return err
	}

	jwtToken := callbackRes.JWTToken

	c.setCookie(ce, jwtToken, c.Auth.CookieSecure)

	successURL, err := urlWithQueryParams(c.RedirectSuccess, url.Values{})
	if err != nil {
		log.Errorf("AuthCallback: %s", err.Error())
		return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
	}
	return ce.Redirect(http.StatusTemporaryRedirect, successURL)
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

func urlWithQueryParams(rawurl string, values url.Values) (s string, err error) {
	var u *url.URL
	if u, err = url.Parse(rawurl); err != nil {
		return
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}
