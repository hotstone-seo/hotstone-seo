package controller

import (
	log "github.com/sirupsen/logrus"

	"github.com/juju/errors"

	"net/http"
	"net/url"

	"github.com/hotstone-seo/hotstone-seo/app/config"
	"github.com/hotstone-seo/hotstone-seo/app/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// AuthCntrl is controller to handle authentication
type AuthCntrl struct {
	dig.In
	config.Config
	service.AuthGoogleService
}

// Route to define API Route
func (c *AuthCntrl) Route(e *echo.Echo) {
	e.GET("auth/google/login", c.AuthGoogleLogin)
	e.GET("auth/google/callback", c.AuthGoogleCallback)
}

// AuthGoogleLogin handle Google auth login
func (c *AuthCntrl) AuthGoogleLogin(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/login] REQ:\n%s\n\n", requestDump)
	// }

	authCodeURL := c.AuthGoogleService.GetAuthCodeURL(ce)
	return ce.Redirect(http.StatusTemporaryRedirect, authCodeURL)
}

// AuthGoogleLogin handle Google auth callback
func (c *AuthCntrl) AuthGoogleCallback(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/callback] REQ:\n%s\n\n", requestDump)
	// }

	holder, err := c.AuthGoogleService.VerifyCallback(ce)
	if err != nil {
		log.Error(errors.Details(err))
		return ce.Redirect(http.StatusTemporaryRedirect, c.Oauth2GoogleRedirectFailure)
	}

	queryParam := url.Values{}
	queryParam.Set("holder", holder)
	successUrl, err := urlWithQueryParams(c.Oauth2GoogleRedirectSuccess, queryParam)
	if err != nil {
		log.Error(errors.Details(err))
		return ce.Redirect(http.StatusTemporaryRedirect, c.Oauth2GoogleRedirectFailure)
	}

	return ce.Redirect(http.StatusTemporaryRedirect, successUrl)
}

func urlWithQueryParams(urlStr string, queryParam url.Values) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", errors.Trace(err)
	}
	u.RawQuery = queryParam.Encode()
	return u.String(), nil
}
