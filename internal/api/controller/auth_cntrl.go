package controller

import (
	"net/http"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/labstack/echo"
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
		Svc2 oauth2google.AuthService
	}
)

// Login with google auth
func (c *AuthCntrl) Login(ce echo.Context) (err error) {
	oauthState := c.Svc2.GenerateOauthState()
	c.Svc2.SetState(ce, oauthState)

	authCodeURL := c.Svc2.GetAuthCodeURL(oauthState)
	return ce.Redirect(http.StatusTemporaryRedirect, authCodeURL)
}

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
