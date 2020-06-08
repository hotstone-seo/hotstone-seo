package oauth2google

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"go.uber.org/dig"
)

type (
	// AuthCntrl is controller to handle authentication
	AuthCntrl struct {
		dig.In
		*Config
		AuthService
	}

	// Callback is called after google auth flow has successfully finished
	Callback func(ce echo.Context, gUser GoogleUser) error
)

// Login with google auth
func (c *AuthCntrl) Login(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/login] REQ:\n%s\n\n", requestDump)
	// }

	oauthState := c.GenerateOauthState()
	c.SetState(ce, oauthState)

	authCodeURL := c.GetAuthCodeURL(oauthState)
	return ce.Redirect(http.StatusTemporaryRedirect, authCodeURL)

}

// Callback for google auth
func (c *AuthCntrl) Callback(cb Callback) func(echo.Context) error {
	return func(ce echo.Context) (err error) {
		ctx := ce.Request().Context()

		// requestDump, err := httputil.DumpRequest(ce.Request(), true)
		// if err == nil {
		// 	log.Warnf("[auth/google/callback] REQ:\n%s\n\n", requestDump)
		// }
		failureURL, err := urlWithQueryParams(c.RedirectFailure, url.Values{"oauth_error": {"true"}})
		if err != nil {
			return fmt.Errorf("AuthCallback: %s", err.Error())
		}

		if !c.VerifyState(ce, ce.QueryParam("state")) {
			return fmt.Errorf("AuthCallback state not valid")
		}

		gUser, err := c.VerifyUser(ctx, ce.QueryParam("code"))
		if err != nil {
			log.Errorf("AuthCallback verify user: %s", err.Error())
			return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
		}

		if err = cb(ce, gUser); err != nil {
			log.Errorf("AuthCallback callback: %s", err.Error())
			return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
		}

		// secureTokenCookie := &http.Cookie{
		// 	Name: "secure_token", Value: string(jwtToken),
		// 	Expires:  time.Now().Add(CookieExpiration),
		// 	Path:     "/",
		// 	HttpOnly: true, Secure: c.Config.CookieSecure,
		// }
		// ce.SetCookie(secureTokenCookie)

		// tokenCookie := &http.Cookie{
		// 	Name: "token", Value: string(jwtToken),
		// 	Expires:  time.Now().Add(CookieExpiration),
		// 	Path:     "/",
		// 	HttpOnly: false, Secure: c.Config.CookieSecure,
		// }
		// ce.SetCookie(tokenCookie)

		// successUrl, err := urlWithQueryParams(c.Oauth2GoogleRedirectSuccess, url.Values{"holder": {holder}})
		successURL, err := urlWithQueryParams(c.RedirectSuccess, url.Values{})
		if err != nil {
			log.Errorf("AuthCallback: %s", err.Error())
			return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
		}
		return ce.Redirect(http.StatusTemporaryRedirect, successURL)
	}
}

// func (c *AuthCntrl) clean(ce echo.Context) (err error) {
// 	ce.SetCookie(&http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"})
// 	ce.SetCookie(&http.Cookie{Name: "token", MaxAge: -1, Path: "/"})
// 	return
// }

// // Logout by invalidating cookies
// func (c *AuthCntrl) Logout(ce echo.Context) (err error) {
// 	c.clean(ce)
// 	return ce.Redirect(http.StatusSeeOther, c.LogoutRedirect)
// }

// // Middleware for google social login
// func (c *AuthCntrl) Middleware() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(ce echo.Context) error {
// 			cfg := middleware.DefaultJWTConfig
// 			cfg.SigningKey = []byte(c.JWTSecret)
// 			cfg.TokenLookup = "cookie:secure_token"

// 			if err := middleware.JWTWithConfig(cfg)(next)(ce); err != nil {
// 				c.clean(ce)
// 				return err
// 			}

// 			return nil
// 		}
// 	}
// }

// // SetTokenCtxMiddleware re-set token to request context for informational purpose (getting username, etc)
// func (c *AuthCntrl) SetTokenCtxMiddleware() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			token := c.Get("user")
// 			currCtx := c.Request().Context()
// 			modifiedReq := c.Request().Clone(
// 				context.WithValue(currCtx, repository.TokenCtxKey, token))
// 			// log.Warnf("# TOKEN: %+v", token)

// 			c.SetRequest(modifiedReq)
// 			return next(c)
// 		}
// 	}
// }

func urlWithQueryParams(rawurl string, values url.Values) (s string, err error) {
	var u *url.URL
	if u, err = url.Parse(rawurl); err != nil {
		return
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}
