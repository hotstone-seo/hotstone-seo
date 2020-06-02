package oauth2google

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
type AuthCntrl struct {
	dig.In
	*Config
	AuthService
}

// Login with google auth
func (c *AuthCntrl) Login(ce echo.Context) (err error) {
	// requestDump, err := httputil.DumpRequest(ce.Request(), true)
	// if err == nil {
	// 	log.Warnf("[auth/google/login] REQ:\n%s\n\n", requestDump)
	// }

	authCodeURL := c.GetAuthCodeURL(ce, c.CookieSecure)
	return ce.Redirect(http.StatusTemporaryRedirect, authCodeURL)
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

// Logout by invalidating cookies
func (c *AuthCntrl) Logout(ce echo.Context) (err error) {
	ce.SetCookie(&http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"})
	ce.SetCookie(&http.Cookie{Name: "token", MaxAge: -1, Path: "/"})
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
				log.Warnf("JWT Error: %s", err.Error())
				return c.Logout(ce)
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
	Label   string `json:"label"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	APIPath string `json:"api_path"`
}

// Middleware for check auth module access
func (c *AuthCntrl) CheckAuthModules() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ce echo.Context) error {
			// get user data from cookie
			user := ce.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			modules := claims["modules"]

			currentAccessAPIPath := ce.Path() // get current API Path

			in := []byte(modules.(string))
			var raw DataModule
			if err := json.Unmarshal(in, &raw); err != nil {
				log.Warnf("JWT Error CheckAuthModules: %s", err.Error())
			}
			modArray := raw.Module

			//TODO: save these paths to modules (in array)
			var newModule = new(Module)
			newModule.Name = "center"
			newModule.Label = "api center for add local business, etc"
			newModule.Path = "-"
			newModule.APIPath = "api/center"
			modArray = append(modArray, *newModule)

			newModule = new(Module)
			newModule.Name = "structureddata"
			newModule.Label = "api center for get structured-data"
			newModule.Path = "-"
			newModule.APIPath = "api/structured-data"
			modArray = append(modArray, *newModule)

			newModule = new(Module)
			newModule.Name = "tags"
			newModule.Label = "api tags"
			newModule.Path = "-"
			newModule.APIPath = "api/tags"
			modArray = append(modArray, *newModule)

			newModule = new(Module)
			newModule.Name = "fetchtags"
			newModule.Label = "fetch tags"
			newModule.Path = "-"
			newModule.APIPath = "p/fetch-tags"
			modArray = append(modArray, *newModule)
			//end TODO

			isAllow := false
			for index, result := range modArray {
				idxStr := strings.Index(currentAccessAPIPath, "/"+result.APIPath)
				if idxStr > -1 {
					log.Infof(currentAccessAPIPath, " was found at index", index)
					isAllow = true
					break
				}
			}

			if !isAllow {
				log.Errorf("CheckAuthModules. Invalid Access")
				failureURL, err := urlWithQueryParams(c.RedirectFailure, url.Values{"oauth_error": {"true"}})
				if err != nil {
					return fmt.Errorf("CheckAuthModules: %s", err.Error())
				}
				ce.SetCookie(&http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"})
				ce.SetCookie(&http.Cookie{Name: "token", MaxAge: -1, Path: "/"})
				return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
			}
			return next(ce)
		}
	}
}
