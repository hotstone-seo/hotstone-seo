package oauth2google

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"go.uber.org/dig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// CookieExpiration is expiration for oauthstate cookie
	CookieExpiration time.Duration = 72 * time.Hour

	StateExpiration time.Duration = 20 * time.Minute

	TokenEpiration time.Duration = 72 * time.Hour
)

type (
	// AuthService is center related logic
	// @mock
	AuthService interface {
		Login() (*LoginResult, error)
		VerifyState(ce echo.Context, state string) bool
		VerifyUser(ctx context.Context, code string) (GoogleUser, error)
	}
	// AuthServiceImpl implementation of AuthService
	AuthServiceImpl struct {
		dig.In
		Cfg *infra.Auth
	}
	googleOauth2UserInfoResp map[string]interface{}
	// GoogleUser holds Google user information
	GoogleUser struct {
		Email   string
		Picture string
	}
	// LoginResult is result of login
	LoginResult struct {
		Redirect string
		Cookie   *http.Cookie
	}
)

// NewService return new instance of AuthGoogleService
// @ctor
func NewService(impl AuthServiceImpl) AuthService {
	return &impl
}

func (c *AuthServiceImpl) oauthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  c.Cfg.Callback,
		ClientID:     c.Cfg.ClientID,
		ClientSecret: c.Cfg.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func (c *AuthServiceImpl) Login() (*LoginResult, error) {
	b := make([]byte, 64)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return &LoginResult{
		Redirect: c.oauthConfig().AuthCodeURL(state),
		Cookie: &http.Cookie{
			Name:     "oauthstate",
			Value:    state,
			Expires:  time.Now().Add(StateExpiration),
			HttpOnly: true,
			Secure:   c.Cfg.CookieSecure,
		},
	}, nil
}

// VerifyState verify oauthstate
func (c *AuthServiceImpl) VerifyState(ce echo.Context, state string) bool {
	oauthState, err := ce.Cookie("oauthstate")
	if err != nil {
		log.Warnf("VerifyState: %+v", err)
		return false
	}

	if ce.QueryParam("state") != oauthState.Value {
		log.Warnf("VerifyState: state not matched")
		return false
	}

	return true
}

// VerifyUser verify and return legitimate username
func (c *AuthServiceImpl) VerifyUser(ctx context.Context, code string) (gUser GoogleUser, err error) {
	userInfoResp, err := c.getUserInfoFromGoogle(ctx, code)
	if err != nil {
		return gUser, fmt.Errorf("AuthVerifyCallback: %w", err)
	}

	err = c.validateUserInfoResp(userInfoResp)
	if err != nil {
		return gUser, fmt.Errorf("AuthVerifyCallback: %w", err)
	}

	return GoogleUser{Email: userInfoResp["email"].(string), Picture: userInfoResp["picture"].(string)}, nil
}

func (c *AuthServiceImpl) getUserInfoFromGoogle(ctx context.Context, code string) (userInfoResp googleOauth2UserInfoResp, err error) {

	config := c.oauthConfig()

	// Use code to get token and get user info from Google.
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("AuthGetUserInfo: %w", err)
	}

	if !token.Valid() {
		return nil, errors.New("AuthGetUserInfo: invalid token")
	}

	response, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken))
	if err != nil {
		return nil, fmt.Errorf("AuthGetUserInfo: %w", err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&userInfoResp)
	if err != nil {
		return nil, fmt.Errorf("AuthGetUserInfo: %w", err)
	}

	return userInfoResp, nil
}

func (c *AuthServiceImpl) validateUserInfoResp(userInfoResp googleOauth2UserInfoResp) error {
	if verifiedEmail, ok := userInfoResp["verified_email"]; !ok || !verifiedEmail.(bool) {
		return errors.New("AuthUserInfo: invalid or empty verified_email")
	}

	if c.Cfg.HostedDomain != "" {
		if hd, ok := userInfoResp["hd"]; !ok || hd != c.Cfg.HostedDomain {
			return errors.New("AuthUserInfo: invalid or empty hd")
		}
	}
	return nil
}
