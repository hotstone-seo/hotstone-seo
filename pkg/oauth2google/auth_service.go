package oauth2google

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/server/repository"

	"github.com/labstack/echo"
	"go.uber.org/dig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// AuthService is center related logic
// @mock
type AuthService interface {
	VerifyCallback(ce echo.Context, jwtSecret string) (string, error)
	GetAuthCodeURL(ce echo.Context, cookieSecure bool) string
}

// AuthServiceImpl implementation of AuthService
type AuthServiceImpl struct {
	dig.In
	*oauth2.Config
	cfg          *Config
	UserRepo     repository.UserRepo
	RoleTypeRepo repository.RoleTypeRepo
}

// NewService return new instance of AuthGoogleService
// @ctor
func NewService(cfg *Config, userRepo repository.UserRepo, roleTypeRepo repository.RoleTypeRepo) AuthService {
	return &AuthServiceImpl{
		cfg: cfg,
		Config: &oauth2.Config{
			RedirectURL:  cfg.Callback,
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"}, // TODO: put to module
			Endpoint:     google.Endpoint,
		},
		UserRepo:     userRepo,
		RoleTypeRepo: roleTypeRepo,
	}
}

func (c *AuthServiceImpl) GetAuthCodeURL(ce echo.Context, cookieSecure bool) (authCodeURL string) {
	// Create oauthState cookie
	oauthState := c.setRandomCookie(ce, "oauthstate", time.Now().Add(StateExpiration), cookieSecure)

	// AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	// validate that it matches the the state query parameter on your redirect callback.
	urlAuthCode := c.AuthCodeURL(oauthState)

	return urlAuthCode
}

// VerifyCallback to add metaTag
func (c *AuthServiceImpl) VerifyCallback(ce echo.Context, jwtSecret string) (string, error) {
	oauthState, err := ce.Cookie("oauthstate")
	if err != nil {
		return "", fmt.Errorf("AuthVerifyCallback: %w", err)
	}

	if ce.QueryParam("state") != oauthState.Value {
		return "", errors.New("invalid oauth google state")
	}

	userInfoResp, err := c.getUserInfoFromGoogle(ce.Request().Context(), ce.QueryParam("code"))
	if err != nil {
		return "", fmt.Errorf("AuthVerifyCallback: %w", err)
	}

	err = c.validateUserInfoResp(userInfoResp)
	if err != nil {
		return "", fmt.Errorf("AuthVerifyCallback: %w", err)
	}

	email := userInfoResp["email"].(string)
	user, err := c.UserRepo.FindUserByEmail(ce.Request().Context(), email)
	if user == nil || err == sql.ErrNoRows {
		return "", fmt.Errorf("AuthVerifyCallback check user exists : %w", err)
	}
	var roleAccess string
	var roleModule string
	if user != nil {
		roleType, err := c.RoleTypeRepo.FindOne(ce.Request().Context(), user.RoleTypeID)
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("AuthVerifyCallback get role: %w", err)
		}
		roleAccess = roleType.Name
		//TODO : still error. will re-check when get value
		// roleModule = roleType.Modules
		// res1D := &response1{roleModule}
		// res1B, _ := json.Marshal(res1D)
		//slcB, _ := json.Marshal(roleModule)
		//fmt.Print(slcB)
		roleModule = `{"modules": ["rules","datasources","mismatchrule","analytic","simulation","audittrail","user","roleType"]}`
	}
	jwtToken, err := c.generateJwtToken(userInfoResp, jwtSecret, user.ID, roleAccess, roleModule)
	if err != nil {
		return "", fmt.Errorf("AuthVerifyCallback: %w", err)
	}

	return jwtToken, nil
}

func (c *AuthServiceImpl) setRandomCookie(ce echo.Context, cookieName string, expiration time.Time, cookieSecure bool) string {
	randomVal := generateRandomBase64(64)
	cookie := &http.Cookie{Name: cookieName, Value: randomVal, Expires: expiration, HttpOnly: true, Secure: cookieSecure}
	ce.SetCookie(cookie)
	return randomVal
}

func (c *AuthServiceImpl) getUserInfoFromGoogle(ctx context.Context, code string) (userInfoResp repository.GoogleOauth2UserInfoResp, err error) {
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

func (c *AuthServiceImpl) validateUserInfoResp(userInfoResp repository.GoogleOauth2UserInfoResp) error {
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

func (c *AuthServiceImpl) generateJwtToken(userInfoResp repository.GoogleOauth2UserInfoResp, jwtSecret string, userID int64, roleAccess string, roleModule string) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userInfoResp["email"]
	claims["picture"] = userInfoResp["picture"]
	claims["exp"] = time.Now().Add(TokenEpiration).Unix()
	claims["user_id"] = userID
	claims["user_role"] = roleAccess

	// fmt.Println(roleModule)
	// str := `{"modules": ["rules","datasources","mismatchrule","analytic","simulation","audittrail","user","roleType"]}`
	res := responseModule{}
	json.Unmarshal([]byte(roleModule), &res)
	claims["modules"] = res

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

type responseModule struct {
	Modules []string `json:"modules"`
}
