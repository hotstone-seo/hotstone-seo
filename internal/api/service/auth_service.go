package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"

	"go.uber.org/dig"
)

var (
	// TokenEpiration is JWT token expiration time
	TokenEpiration time.Duration = 72 * time.Hour
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
)

// NewService return new instance of AuthGoogleService
// @ctor
func NewService(userRepo repository.UserRepo, UserRoleRepo repository.UserRoleRepo, settingSvc SettingSvc) AuthService {
	return &AuthServiceImpl{
		UserRepo:     userRepo,
		UserRoleRepo: UserRoleRepo,
		SettingSvc:   settingSvc,
	}
}

// BuildJwtClaims build JWT claims based on given user
func (c *AuthServiceImpl) BuildJwtClaims(ctx context.Context, gUser oauth2google.GoogleUser) (jwtClaims JwtClaims, err error) {
	user, err := c.UserRepo.FindUserByEmail(ctx, gUser.Email)
	if user == nil || err == sql.ErrNoRows {
		return jwtClaims, fmt.Errorf("AuthVerifyCallback check user exists : %w", err)
	}
	var roleAccess string
	var roleMenus []string
	var rolePaths []string
	if user != nil {
		UserRole, err := c.UserRoleRepo.FindOne(ctx, user.UserRoleID)
		if err == sql.ErrNoRows {
			return jwtClaims, fmt.Errorf("AuthVerifyCallback get role modules: %w", err)
		}
		roleAccess = UserRole.Name
		roleMenus = UserRole.Menus
		rolePaths = UserRole.Paths

	}
	simulationKey := c.SettingSvc.GetValue(ctx, SimulationKey)
	return JwtClaims{
		email:         gUser.Email,
		picture:       gUser.Picture,
		userID:        user.ID,
		userRole:      roleAccess,
		simulationKey: simulationKey,
		menus:         roleMenus,
		paths:         rolePaths,
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
