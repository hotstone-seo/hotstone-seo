package controller

import (
	"context"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
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
		*infra.App
		service.AuthService
	}
)

// Oauth2GoogleCallback is called after google auth flow has been successfully finished
func (c *AuthCntrl) Oauth2GoogleCallback(ce echo.Context, gUser oauth2google.GoogleUser) error {
	ctx := ce.Request().Context()
	jwtClaims, err := c.AuthService.BuildJwtClaims(ctx, gUser)
	if err != nil {
		return err
	}

	jwtToken, err := c.AuthService.GenerateJwtToken(jwtClaims, c.App.JWTSecret)
	if err != nil {
		return err
	}

	secureTokenCookie := &http.Cookie{
		Name: "secure_token", Value: string(jwtToken),
		Expires:  time.Now().Add(CookieExpiration),
		Path:     "/",
		HttpOnly: true, Secure: c.App.CookieSecure,
	}
	ce.SetCookie(secureTokenCookie)

	tokenCookie := &http.Cookie{
		Name: "token", Value: string(jwtToken),
		Expires:  time.Now().Add(CookieExpiration),
		Path:     "/",
		HttpOnly: false, Secure: c.App.CookieSecure,
	}
	ce.SetCookie(tokenCookie)
	return nil
}

// Middleware for google social login
func (c *AuthCntrl) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ce echo.Context) error {
			cfg := middleware.DefaultJWTConfig
			cfg.SigningKey = []byte(c.App.JWTSecret)
			cfg.TokenLookup = "cookie:secure_token"

			if err := middleware.JWTWithConfig(cfg)(next)(ce); err != nil {
				cleanCookie(ce)
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
				context.WithValue(currCtx, service.TokenCtxKey, token))
			// log.Warnf("# TOKEN: %+v", token)

			c.SetRequest(modifiedReq)
			return next(c)
		}
	}
}

func cleanCookie(ce echo.Context) (err error) {
	ce.SetCookie(&http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"})
	ce.SetCookie(&http.Cookie{Name: "token", MaxAge: -1, Path: "/"})
	return
}

// Logout by invalidating cookies
func (c *AuthCntrl) Logout(ce echo.Context) (err error) {
	cleanCookie(ce)
	return ce.Redirect(http.StatusSeeOther, c.LogoutRedirect)
}

// CheckAuthModules for check auth module access
func (c *AuthCntrl) CheckAuthModules() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ce echo.Context) error {
			// get user data from cookie
			user := ce.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			path := ce.Path()
			if !IsRoleAllow(path, claims) {
				log.Errorf("CheckAuthModules. Invalid Access ", path)
				cleanCookie(ce)
			}
			return next(ce)
		}
	}
}

// IsRoleAllow check is path allow by role
func IsRoleAllow(path string, claims jwt.MapClaims) bool {
	rolePaths := claims["paths"].([]interface{})
	for _, v := range rolePaths {
		matched, _ := regexp.MatchString(v.(string), path)
		if matched {
			return true
		}
	}
	return false
}

func urlWithQueryParams(rawurl string, values url.Values) (s string, err error) {
	var u *url.URL
	if u, err = url.Parse(rawurl); err != nil {
		return
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}
