package oauth2google

import (
	"fmt"
	"net/http"
	"net/url"

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
