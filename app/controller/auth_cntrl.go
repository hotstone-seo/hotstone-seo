package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/hotstone-seo/hotstone-seo/app/config"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

type GoogleOauth2UserInfoResp map[string]interface{}

// NewOauth2Config return new instance of oauth2.Config [constructor]
func NewOauth2Config(config config.Config) *oauth2.Config {

	c := oauth2.Config{
		RedirectURL:  config.Oauth2GoogleCallback,
		ClientID:     config.Oauth2GoogleClientID,
		ClientSecret: config.Oauth2GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return &c
}

// AuthCntrl is controller to handle authentication
type AuthCntrl struct {
	dig.In
	config.Config
	Oauth2Config *oauth2.Config
}

// Route to define API Route
func (c *AuthCntrl) Route(e *echo.Echo) {
	e.Any("auth/google/login", c.AuthGoogleLogin)
	e.Any("auth/google/callback", c.AuthGoogleCallback)

}

// AuthGoogleLogin handle Google auth login
func (c *AuthCntrl) AuthGoogleLogin(ce echo.Context) (err error) {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(ce.Response())

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	url := c.Oauth2Config.AuthCodeURL(oauthState)

	return ce.Redirect(http.StatusTemporaryRedirect, url)
}

// AuthGoogleLogin handle Google auth callback
func (c *AuthCntrl) AuthGoogleCallback(ce echo.Context) (err error) {
	// Read oauthState from Cookie
	oauthState, _ := ce.Cookie("oauthstate")

	if ce.FormValue("state") != oauthState.Value {
		return errors.New("invalid oauth google state")
	}

	userInfoResp, err := c.getUserInfoFromGoogle(ce.FormValue("code"))
	if err != nil {
		return
	}

	return ce.String(http.StatusOK, fmt.Sprintf("%+v", userInfoResp))
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func (c *AuthCntrl) getUserInfoFromGoogle(code string) (userInfoResp GoogleOauth2UserInfoResp, err error) {
	// Use code to get token and get user info from Google.
	token, err := c.Oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	if !token.Valid() {
		return nil, fmt.Errorf("invalid token")
	}

	response, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken))
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&userInfoResp)
	if err != nil {
		return nil, err
	}

	return userInfoResp, nil
}
