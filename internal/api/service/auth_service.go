package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"

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
		BuildJwtClaims(ctx context.Context, gUser oauth2google.GoogleUser) (JwtClaims, error)
		GenerateJwtToken(jwtClaim JwtClaims, jwtSecret string) (string, error)
	}
	// AuthServiceImpl implementation of AuthService
	AuthServiceImpl struct {
		dig.In
		UserRepo     repository.UserRepo
		UserRoleRepo repository.UserRoleRepo
		SettingSvc   SettingSvc
	}
	// JwtClaims holds JWT claims information
	JwtClaims struct {
		email         string
		picture       string
		userID        int64
		userRole      string
		modules       string
		simulationKey string
		paths         []string
		menus         []string
	}
	CtxKey int
)

// NewAuthService return new instance of AuthGoogleService
// @ctor
func NewAuthService(impl AuthServiceImpl) AuthService {
	return &impl
}

// BuildJwtClaims build JWT claims based on given user
func (c *AuthServiceImpl) BuildJwtClaims(ctx context.Context, gUser oauth2google.GoogleUser) (jwtClaims JwtClaims, err error) {
	users, _ := c.UserRepo.Find(ctx, dbkit.Equal(repository.UserTable.Email, gUser.Email))
	if len(users) < 1 {
		return jwtClaims, fmt.Errorf("AuthVerifyCallback check user exists : %w", err)
	}

	role, err := c.UserRoleRepo.FindOne(ctx, users[0].UserRoleID)
	if err == sql.ErrNoRows {
		return jwtClaims, fmt.Errorf("AuthVerifyCallback get role modules: %w", err)
	}

	simulationKey := c.SettingSvc.GetValue(ctx, SimulationKey)
	return JwtClaims{
		email:         gUser.Email,
		picture:       gUser.Picture,
		userID:        users[0].ID,
		userRole:      role.Name,
		simulationKey: simulationKey,
		menus:         role.Menus,
		paths:         role.Paths,
	}, nil
}

// GenerateJwtToken generates and returns JWT token with additional claims
func (c *AuthServiceImpl) GenerateJwtToken(jwtClaim JwtClaims, jwtSecret string) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = jwtClaim.email
	claims["picture"] = jwtClaim.picture
	claims["exp"] = time.Now().Add(TokenEpiration).Unix()
	claims["user_id"] = jwtClaim.userID
	claims["user_role"] = jwtClaim.userRole
	claims["modules"] = jwtClaim.modules
	claims["simulation_key"] = jwtClaim.simulationKey
	claims["menus"] = jwtClaim.menus
	claims["paths"] = jwtClaim.paths

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
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
