package controller

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/juju/errors"

	"net/http"
	"net/url"

	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

var (
	JwtTokenCookieExpire time.Duration = 72 * time.Hour
)

// AuthCntrl is controller to handle authentication
type AuthCntrl struct {
	dig.In
	config.Config
	service.AuthGoogleService
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

// AuthGoogleCallback handle Google auth callback
func (c *AuthCntrl) AuthGoogleCallback(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/callback] REQ:\n%s\n\n", requestDump)
	// }
	failureURL, err := urlWithQueryParams(c.Oauth2GoogleRedirectFailure, url.Values{"oauth_error": {"true"}})
	if err != nil {
		return errors.Trace(err)
	}

	jwtToken, err := c.AuthGoogleService.VerifyCallback(ce)
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
		Expires:  time.Now().Add(JwtTokenCookieExpire),
		Path:     "/",
		HttpOnly: true, Secure: c.Config.CookieSecure,
	}
	ce.SetCookie(secureTokenCookie)

	tokenCookie := &http.Cookie{
		Name: "token", Value: string(jwtToken),
		Expires:  time.Now().Add(JwtTokenCookieExpire),
		Path:     "/",
		HttpOnly: false, Secure: c.Config.CookieSecure,
	}
	ce.SetCookie(tokenCookie)

	return ce.Redirect(http.StatusTemporaryRedirect, successURL)
}

func (c *AuthCntrl) AuthGoogleToken(ce echo.Context) (err error) {
	var (
		req      repository.TokenReq
		jwtToken []byte
		ctx      = ce.Request().Context()
	)
	if err = ce.Bind(&req); err != nil {
		return
	}
	if jwtToken, err = c.AuthGoogleService.GetThenDeleteJwtToken(ctx, req.Holder); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if req.SetCookie {
		secureTokenCookie := &http.Cookie{
			Name: "secure_token", Value: string(jwtToken),
			Expires:  time.Now().Add(JwtTokenCookieExpire),
			Path:     "/",
			HttpOnly: true, Secure: c.Config.CookieSecure,
		}
		ce.SetCookie(secureTokenCookie)

		tokenCookie := &http.Cookie{
			Name: "token", Value: string(jwtToken),
			Expires:  time.Now().Add(JwtTokenCookieExpire),
			Path:     "/",
			HttpOnly: false, Secure: c.Config.CookieSecure,
		}
		ce.SetCookie(tokenCookie)
	}

	return ce.JSON(http.StatusOK, repository.TokenResp{Token: string(jwtToken)})
}

// AuthLogout handle logout by invalidating cookies
func (c *AuthCntrl) AuthLogout(ce echo.Context) (err error) {
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