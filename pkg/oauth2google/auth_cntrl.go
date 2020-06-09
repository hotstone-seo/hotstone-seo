package oauth2google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
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

type DataModule struct {
	Module []Module `json:"modules"`
}
type Module struct {
	Label   string     `json:"label"`
	Name    string     `json:"name"`
	Path    string     `json:"path"`
	APIPath []APIPathS `json:"api_path"`
}

type APIPathS struct {
	Path string `json:"path"`
}

// Login with google auth
func (c *AuthCntrl) Login(ce echo.Context) (err error) {
	oauthState := c.GenerateOauthState()
	c.SetState(ce, oauthState)

	authCodeURL := c.GetAuthCodeURL(oauthState)
	return ce.Redirect(http.StatusTemporaryRedirect, authCodeURL)

}

// Callback for google auth
func (c *AuthCntrl) Callback(cb Callback) func(echo.Context) error {
	return func(ce echo.Context) (err error) {
		ctx := ce.Request().Context()

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

		successURL, err := urlWithQueryParams(c.RedirectSuccess, url.Values{})
		if err != nil {
			log.Errorf("AuthCallback: %s", err.Error())
			return ce.Redirect(http.StatusTemporaryRedirect, failureURL)
		}
		return ce.Redirect(http.StatusTemporaryRedirect, successURL)
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

			isAllow := false
			for index, result := range modArray {
				for k, v := range result.APIPath {
					idxStr := strings.Index(currentAccessAPIPath, v.Path)
					if idxStr > -1 {
						log.Infof(currentAccessAPIPath, " was found at index", index, ";", k)
						isAllow = true
						break
					}
				}
				if isAllow {
					break
				}
			}
			if !isAllow {
				log.Errorf("CheckAuthModules. Invalid Access")
				ce.SetCookie(&http.Cookie{Name: "secure_token", MaxAge: -1, Path: "/"})
				ce.SetCookie(&http.Cookie{Name: "token", MaxAge: -1, Path: "/"})

				failureURL, err := urlWithQueryParams(c.RedirectFailure, url.Values{"oauth_error": {"true"}})
				if err != nil {
					return fmt.Errorf("CheckAuthModules: %s", err.Error())
				}
				return ce.Redirect(http.StatusSeeOther, failureURL)
			}
			return next(ce)
		}
	}
}
