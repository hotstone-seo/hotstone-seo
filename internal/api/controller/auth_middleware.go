package controller

import (
	"context"
	"net/url"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// AuthMiddleware is middleware for authentication
type AuthMiddleware struct {
	dig.In
	*infra.App
}

// Middleware for google social login
func (c *AuthMiddleware) Middleware() echo.MiddlewareFunc {
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
func (c *AuthMiddleware) SetTokenCtxMiddleware() echo.MiddlewareFunc {
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

// CheckAuthModules for check auth module access
func (c *AuthMiddleware) CheckAuthModules() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ce echo.Context) error {
			// get user data from cookie
			user := ce.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			path := ce.Path()
			if !IsRoleAllow(path, claims) {
				log.Error("CheckAuthModules. Invalid Access " + path)
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
