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

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"go.uber.org/dig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type (

	// AuthService is center related logic
	// @mock
	AuthService interface {
		GenerateOauthState() string
		SetState(ce echo.Context, state string)
		VerifyState(ce echo.Context, state string) bool
		GetAuthCodeURL(coauthState string) string
		VerifyUser(ctx context.Context, code string) (GoogleUser, error)
	}

	// AuthServiceImpl implementation of AuthService
	AuthServiceImpl struct {
		dig.In
		*oauth2.Config
		cfg *Config
	}

	googleOauth2UserInfoResp map[string]interface{}

	// GoogleUser holds Google user information
	GoogleUser struct {
		Email   string
		Picture string
	}
)

// NewService return new instance of AuthGoogleService
// @ctor
func NewService(cfg *Config) AuthService {
	return &AuthServiceImpl{
		cfg: cfg,
		Config: &oauth2.Config{
			RedirectURL:  cfg.Callback,
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"}, // TODO: put to module
			Endpoint:     google.Endpoint,
		},
	}
}

// GenerateOauthState generate oauth state
func (c *AuthServiceImpl) GenerateOauthState() (oauthState string) {
	return generateRandomBase64(64)
}

// SetState set oauthstate to cookie
func (c *AuthServiceImpl) SetState(ce echo.Context, state string) {
	expire := time.Now().Add(StateExpiration)
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expire, HttpOnly: true, Secure: c.cfg.CookieSecure}
	ce.SetCookie(cookie)
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

// GetAuthCodeURL returns URL to validate and protect user from CSRF attacks
func (c *AuthServiceImpl) GetAuthCodeURL(oauthState string) (authCodeURL string) {
	urlAuthCode := c.AuthCodeURL(oauthState)

	return urlAuthCode
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
	// Use code to get token and get user info from Google.
	token, err := c.Exchange(ctx, code)
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

	if c.cfg.HostedDomain != "" {
		if hd, ok := userInfoResp["hd"]; !ok || hd != c.cfg.HostedDomain {
			return errors.New("AuthUserInfo: invalid or empty hd")
		}
	}
	return nil
}

func generateRandomBase64(keyLength int) string {
	b := make([]byte, keyLength)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
