package oauth2google

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	log "github.com/sirupsen/logrus"

	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"
)

// AuthCntrl is controller to handle authentication
type (
	AuthCntrl struct {
		dig.In
		*Config
		AuthService
	}

	Callback func(string) error
)

// Login with google auth
func (c *AuthCntrl) Login(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/login] REQ:\n%s\n\n", requestDump)
	// }

	authCodeURL := c.GetAuthCodeURL(ce, c.CookieSecure)
	return ce.Redirect(http.StatusTemporaryRedirect, authCodeURL)
}

func (c *AuthCntrl) CallbackX(cb Callback) func(echo.Context) error {
	return func(ce echo.Context) (err error) {
		return nil
	}
}

// Callback for google auth
func (c *AuthCntrl) Callback(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/callback] REQ:\n%s\n\n", requestDump)
	// }
	failureURL, err := urlWithQueryParams(c.RedirectFailure, url.Values{"oauth_error": {"true"}})
	if err != nil {
		return fmt.Errorf("AuthCallback: %s", err.Error())
	}

	jwtToken, err := c.VerifyCallback(ce, c.JWTSecret)
	if err != nil {
		log.Errorf("AuthCallback: %s", err.Error())
		return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
	}

	// successUrl, err := urlWithQueryParams(c.Oauth2GoogleRedirectSuccess, url.Values{"holder": {holder}})
	successURL, err := urlWithQueryParams(c.RedirectSuccess, url.Values{})
	if err != nil {
		log.Errorf("AuthCallback: %s", err.Error())
		return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
	}

	secureTokenCookie := &http.Cookie{
		Name: "secure_token", Value: string(jwtToken),
		Expires:  time.Now().Add(CookieExpiration),
		Path:     "/",
		HttpOnly: true, Secure: c.Config.CookieSecure,
	}
	ce.SetCookie(secureTokenCookie)

	tokenCookie := &http.Cookie{
		Name: "token", Value: string(jwtToken),
		Expires:  time.Now().Add(CookieExpiration),
		Path:     "/",
		HttpOnly: false, Secure: c.Config.CookieSecure,
	}
	ce.SetCookie(tokenCookie)

	return ce.Redirect(http.StatusTemporaryRedirect, successURL)
}

func (c *AuthCntrl) clean(ce echo.Context) (err error) {
	ce.SetCookie(&http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"})
	ce.SetCookie(&http.Cookie{Name: "token", MaxAge: -1, Path: "/"})
	return
}

// Logout by invalidating cookies
func (c *AuthCntrl) Logout(ce echo.Context) (err error) {
	c.clean(ce)
	return ce.Redirect(http.StatusSeeOther, c.LogoutRedirect)
}

// Middleware for google social login
func (c *AuthCntrl) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ce echo.Context) error {
			cfg := middleware.DefaultJWTConfig
			cfg.SigningKey = []byte(c.JWTSecret)
			cfg.TokenLookup = "cookie:secure_token"

			if err := middleware.JWTWithConfig(cfg)(next)(ce); err != nil {
				c.clean(ce)
				return err
			}

			return nil
		}
	}
}

// SetTokenCtxMiddleware re-set token to request context for informational purpose (getting username, etc)
func (c *AuthCntrl) SetTokenCtxMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user")
			currCtx := c.Request().Context()
			modifiedReq := c.Request().Clone(
				context.WithValue(currCtx, repository.TokenCtxKey, token))
			// log.Warnf("# TOKEN: %+v", token)

			c.SetRequest(modifiedReq)
			return next(c)
		}
	}
}

func urlWithQueryParams(rawurl string, values url.Values) (s string, err error) {
	var u *url.URL
	if u, err = url.Parse(rawurl); err != nil {
		return
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}

type DataModule struct {
	Module []Module `json:"modules"`
}
type Module struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	Path  string `json:"path"`
}

// Middleware for check auth module access
func (c *AuthCntrl) CheckAuthModules() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ce echo.Context) error {
			// get user data from cookie
			user := ce.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			modules := claims["modules"]

			currentURLPath := ce.Path() // get API Path

			in := []byte(modules.(string))
			var raw DataModule
			if err := json.Unmarshal(in, &raw); err != nil {
				log.Warnf("JWT Error CheckAuthModules: %s", err.Error())
			}
			modArray := raw.Module

			isAllow := false
			for index, result := range modArray {
				fmt.Println(index)

				// TODO: compare currentURLPath with regex pattern module
				if "/api/"+result.Path == currentURLPath { //compare current URL with user privilege user
					isAllow = true
					break
				}
			}
			fmt.Println(isAllow)
			/* if !isAllow {
				// return c.Logout(ce)
				return echo.ErrUnauthorized
			} */
			return next(ce)
		}
	}
}
