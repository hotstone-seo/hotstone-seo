package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/hotstone-seo/hotstone-seo/pkg/gauthkit"

	"go.uber.org/dig"
)

var (
	// TokenEpiration is JWT token expiration time
	TokenEpiration time.Duration = 72 * time.Hour
)

const (
	TokenCtxKey CtxKey = iota
)

type (
	// AuthService is center related logic
	// @mock
	AuthService interface {
		GenerateJwtToken(ctx context.Context, gUser *gauthkit.UserInfo) (string, error)
	}
	// AuthServiceImpl implementation of AuthService
	AuthServiceImpl struct {
		dig.In
		*infra.Auth
		UserRepo     repository.UserRepo
		UserRoleRepo repository.UserRoleRepo
		SettingSvc   SettingSvc
	}
	CtxKey int
)

// NewAuthService return new instance of AuthGoogleService
// @ctor
func NewAuthService(impl AuthServiceImpl) AuthService {
	return &impl
}

// GenerateJwtToken generates and returns JWT token with additional claims
func (c *AuthServiceImpl) GenerateJwtToken(ctx context.Context, gUser *gauthkit.UserInfo) (string, error) {

	users, _ := c.UserRepo.Find(ctx, dbkit.Equal(repository.UserTable.Email, gUser.Email))
	if len(users) < 1 {
		return "", errors.New("User is missing")
	}

	role, _ := c.UserRoleRepo.FindOne(ctx, users[0].UserRoleID)
	if role == nil {
		return "", errors.New("Role is missing")
	}

	simulationKey := c.SettingSvc.GetValue(ctx, SimulationKey)

	return c.signedKey(c.Auth.JWTSecret, map[string]interface{}{
		"email":          gUser.Email,
		"picture":        gUser.Picture,
		"exp":            time.Now().Add(TokenEpiration).Unix(),
		"user_id":        users[0].ID,
		"user_role":      role.Name,
		"simulation_key": simulationKey,
		"menus":          role.Menus,
		"paths":          role.Paths,
	})
}

func (c *AuthServiceImpl) signedKey(key string, data map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range data {
		claims[k] = v
	}
	return token.SignedString([]byte(c.JWTSecret))
}

func generateRandomBase64(keyLength int) string {
	b := make([]byte, keyLength)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
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
