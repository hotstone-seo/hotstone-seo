package controller

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/juju/errors"

	"net/http"
	"net/url"

	"github.com/hotstone-seo/hotstone-seo/pkg/gsociallogin"
	"github.com/hotstone-seo/hotstone-seo/server/config"

	"github.com/labstack/echo"
	"go.uber.org/dig"
)

var (
	// JWTCookieExpire is expiration for JWT cookie
	JWTCookieExpire time.Duration = 72 * time.Hour // TODO: put in config
)

// AuthCntrl is controller to handle authentication
type AuthCntrl struct {
	dig.In
	*config.Config
	gsociallogin.Service
}

// Login with google auth
func (c *AuthCntrl) Login(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/login] REQ:\n%s\n\n", requestDump)
	// }

	authCodeURL := c.Service.GetAuthCodeURL(ce, c.CookieSecure)
	return ce.Redirect(http.StatusTemporaryRedirect, authCodeURL)
}

// Callback for google auth
func (c *AuthCntrl) Callback(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/callback] REQ:\n%s\n\n", requestDump)
	// }
	failureURL, err := urlWithQueryParams(c.Oauth2GoogleRedirectFailure, url.Values{"oauth_error": {"true"}})
	if err != nil {
		return errors.Trace(err)
	}

	jwtToken, err := c.VerifyCallback(ce, c.JwtSecret)
	if err != nil {
		log.Error(errors.Details(err))
		return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
	}

	// successUrl, err := urlWithQueryParams(c.Oauth2GoogleRedirectSuccess, url.Values{"holder": {holder}})
	successURL, err := urlWithQueryParams(c.Oauth2GoogleRedirectSuccess, url.Values{})
	if err != nil {
		log.Error(errors.Details(err))
		return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
	}

	secureTokenCookie := &http.Cookie{
		Name: "secure_token", Value: string(jwtToken),
		Expires:  time.Now().Add(JWTCookieExpire),
		Path:     "/",
		HttpOnly: true, Secure: c.Config.CookieSecure,
	}
	ce.SetCookie(secureTokenCookie)

	tokenCookie := &http.Cookie{
		Name: "token", Value: string(jwtToken),
		Expires:  time.Now().Add(JWTCookieExpire),
		Path:     "/",
		HttpOnly: false, Secure: c.Config.CookieSecure,
	}
	ce.SetCookie(tokenCookie)

	return ce.Redirect(http.StatusTemporaryRedirect, successURL)
}

// Logout by invalidating cookies
func (c *AuthCntrl) Logout(ce echo.Context) (err error) {
	secureTokenCookie := &http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"}
	ce.SetCookie(secureTokenCookie)

	tokenCookie := &http.Cookie{Name: "token", MaxAge: -1, Path: "/"}
	ce.SetCookie(tokenCookie)

	return ce.Redirect(http.StatusSeeOther, c.Config.AuthLogoutRedirect)
}

func urlWithQueryParams(urlStr string, queryParam url.Values) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", errors.Trace(err)
	}
	u.RawQuery = queryParam.Encode()
	return u.String(), nil
}
