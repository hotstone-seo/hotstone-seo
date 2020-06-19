package oauth2google

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/pkg/gauthkit"
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
		Callback(context.Context, *CallbackRequest) (*CallbackResult, error)
		VerifyUser(ctx context.Context, code string) (*gauthkit.UserInfo, error)
	}
	// AuthServiceImpl implementation of AuthService
	AuthServiceImpl struct {
		dig.In
		Cfg *infra.Auth
	}

	// LoginResult is result of login
	LoginResult struct {
		Redirect string
		Cookie   *http.Cookie
	}
	// CallbackRequest is request of callback
	CallbackRequest struct {
		OAuthState *http.Cookie
		StateParam string
	}
	// CallbackResult is result of callback
	CallbackResult struct{}
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

// Login process
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

// Callback process
func (c *AuthServiceImpl) Callback(ctx context.Context, req *CallbackRequest) (*CallbackResult, error) {
	if req.OAuthState == nil || req.StateParam != req.OAuthState.Value {
		return nil, errors.New("bad-state")
	}
	return nil, nil
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
func (c *AuthServiceImpl) VerifyUser(ctx context.Context, code string) (*gauthkit.UserInfo, error) {

	config := c.oauthConfig()

	// Use code to get token and get user info from Google.
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("verify-user: %w", err)
	}
	if !token.Valid() {
		return nil, errors.New("verify-user: invalid token")
	}

	userInfo, err := gauthkit.RetrieveUserInfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("verify-user: %w", err)
	}

	if !userInfo.VerifiedEmail {
		return nil, errors.New("verify-user: invalid or empty verified_email")
	}

	if userInfo.Hd != c.Cfg.HostedDomain && c.Cfg.HostedDomain != "" {
		return nil, errors.New("verify-user: invalid or empty hd")
	}

	return userInfo, nil
}
