package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/pkg/gauthkit"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
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
		service.AuthService
		Svc2 oauth2google.AuthService
	}
)

// Login with google auth
func (c *AuthCntrl) Login(ce echo.Context) error {
	result, err := c.Svc2.Login()
	if err != nil {
		return err
	}
	ce.SetCookie(result.Cookie)
	return ce.Redirect(http.StatusTemporaryRedirect, result.Redirect)
}

func (c *AuthCntrl) Callback(ce echo.Context) (err error) {
	ctx := ce.Request().Context()

	failureURL, err := urlWithQueryParams(c.RedirectFailure, url.Values{"oauth_error": {"true"}})
	if err != nil {
		return fmt.Errorf("AuthCallback: %s", err.Error())
	}

	oauthState, _ := ce.Cookie("oauthstate")

	_, err = c.Svc2.Callback(ctx, &oauth2google.CallbackRequest{
		OAuthState: oauthState,
		StateParam: ce.QueryParam("state"),
	})
	if err != nil {
		errMsg := err.Error()
		if strings.HasPrefix(errMsg, "callback-failed:") {
			log.Error(errMsg)
			return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
		}
		return err
	}

	gUser, err := c.Svc2.VerifyUser(ctx, ce.QueryParam("code"))
	if err != nil {
		log.Errorf("AuthCallback verify user: %s", err.Error())
		return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
	}

	if err = c.callback(ce, gUser); err != nil {
		log.Errorf("AuthCallback callback: %s", err.Error())
		return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
	}

	successURL, err := urlWithQueryParams(c.RedirectSuccess, url.Values{})
	if err != nil {
		log.Errorf("AuthCallback: %s", err.Error())
		return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
	}
	return ce.Redirect(http.StatusTemporaryRedirect, successURL)
}

func (c *AuthCntrl) callback(ce echo.Context, gUser *gauthkit.UserInfo) error {
	ctx := ce.Request().Context()

	jwtToken, err := c.AuthService.GenerateJwtToken(ctx, gUser)
	if err != nil {
		return err
	}

	ce.SetCookie(&http.Cookie{
		Name:     "secure_token",
		Value:    string(jwtToken),
		Expires:  time.Now().Add(CookieExpiration),
		Path:     "/",
		HttpOnly: true,
		Secure:   c.Auth.CookieSecure,
	})

	ce.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    string(jwtToken),
		Expires:  time.Now().Add(CookieExpiration),
		Path:     "/",
		HttpOnly: false,
		Secure:   c.Auth.CookieSecure,
	})
	return nil
}

// Logout by invalidating cookies
func (c *AuthCntrl) Logout(ce echo.Context) (err error) {
	cleanCookie(ce)
	return ce.Redirect(http.StatusSeeOther, c.LogoutRedirect)
}

func cleanCookie(ce echo.Context) (err error) {
	ce.SetCookie(&http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"})
	ce.SetCookie(&http.Cookie{Name: "token", MaxAge: -1, Path: "/"})
	return
}

func urlWithQueryParams(rawurl string, values url.Values) (s string, err error) {
	var u *url.URL
	if u, err = url.Parse(rawurl); err != nil {
		return
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}
