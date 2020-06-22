package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/pkg/gauthkit"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"go.uber.org/dig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	TokenCtxKey CtxKey = iota
)

var (
	// SimulationKey is HotStone client key for Simulation
	SimulationKey string = "simulation_key"
)

type (
	// AuthSvc is center related logic
	// @mock
	AuthSvc interface {
		Login() (*LoginResult, error)
		Callback(context.Context, *CallbackRequest) (*CallbackResult, error)
	}
	// AuthSvcImpl implementation of AuthService
	AuthSvcImpl struct {
		dig.In
		*infra.Auth
		UserRepo     repository.UserRepo
		UserRoleRepo repository.UserRoleRepo
		SettingSvc   SettingSvc
	}
	// LoginResult is result of login
	LoginResult struct {
		Redirect string
		State    string
	}
	// CallbackRequest is request of callback
	CallbackRequest struct {
		OAuthState      *http.Cookie
		StateParam      string
		CodeParam       string
		TokenExpiration time.Duration
	}
	// CallbackResult is result of callback
	CallbackResult struct {
		JWTToken string
	}
	CtxKey int
)

// NewService return new instance of AuthGoogleService
// @ctor
func NewService(impl AuthSvcImpl) AuthSvc {
	return &impl
}

func (c *AuthSvcImpl) oauthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  c.Auth.Callback,
		ClientID:     c.Auth.ClientID,
		ClientSecret: c.Auth.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// Login process
func (c *AuthSvcImpl) Login() (*LoginResult, error) {
	b := make([]byte, 64)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return &LoginResult{
		Redirect: c.oauthConfig().AuthCodeURL(state),
		State:    state,
	}, nil
}

// Callback process
func (c *AuthSvcImpl) Callback(ctx context.Context, req *CallbackRequest) (*CallbackResult, error) {
	if req.OAuthState == nil || req.StateParam != req.OAuthState.Value {
		return nil, errors.New("bad-state")
	}

	token, err := c.exchange(ctx, req.CodeParam)
	if err != nil {
		return nil, fmt.Errorf("callback-failed: %w", err)
	}

	userInfo, err := c.retrieveUserInfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("callback-failed: %w", err)
	}

	users, _ := c.UserRepo.Find(ctx, dbkit.Equal(repository.UserTable.Email, userInfo.Email))
	if len(users) < 1 {
		return nil, errors.New("callback-failed: user is missing")
	}

	role, _ := c.UserRoleRepo.FindOne(ctx, users[0].UserRoleID)
	if role == nil {
		return nil, errors.New("callback-failed: role is missing")
	}

	simulationKey := c.SettingSvc.GetValue(ctx, SimulationKey)

	jwtToken, err := c.signedKey(c.Auth.JWTSecret, map[string]interface{}{
		"email":          userInfo.Email,
		"picture":        userInfo.Picture,
		"exp":            time.Now().Add(req.TokenExpiration).Unix(),
		"user_id":        users[0].ID,
		"user_role":      role.Name,
		"simulation_key": simulationKey,
		"menus":          role.Menus,
		"paths":          role.Paths,
	})
	if err != nil {
		return nil, fmt.Errorf("callback-failed: %w", err)
	}

	return &CallbackResult{
		JWTToken: jwtToken,
	}, nil
}

func (c *AuthSvcImpl) exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.oauthConfig().Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if !token.Valid() {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

func (c *AuthSvcImpl) retrieveUserInfo(ctx context.Context, token *oauth2.Token) (*gauthkit.UserInfo, error) {
	userInfo, err := gauthkit.RetrieveUserInfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("verify-user: %w", err)
	}
	if !userInfo.VerifiedEmail {
		return nil, errors.New("verify-user: invalid or empty verified_email")
	}
	if userInfo.Hd != c.Auth.HostedDomain && c.Auth.HostedDomain != "" {
		return nil, errors.New("verify-user: invalid or empty hd")
	}
	return userInfo, nil
}

func (c *AuthSvcImpl) signedKey(key string, data map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range data {
		claims[k] = v
	}
	return token.SignedString([]byte(c.Auth.JWTSecret))
}

func GetUsername(ctx context.Context) string {
	token, ok := ctx.Value(TokenCtxKey).(*jwt.Token)
	if !ok {
		return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}
	email, ok := claims["email"].(string)
	if !ok {
		return ""
	}
	return email
}
