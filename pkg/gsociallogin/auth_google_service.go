package gsociallogin

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	OAuthStateCookieExpire time.Duration = 20 * time.Minute
	JwtTokenHolderExpire   time.Duration = 60 * time.Second
	JwtTokenExpire         time.Duration = 72 * time.Hour
)

// AuthGoogleService is center related logic [mock]
type AuthGoogleService interface {
	VerifyCallback(ce echo.Context, jwtSecret string) (string, error)
	GetAuthCodeURL(ce echo.Context, cookieSecure bool) string
}

// AuthGoogleServiceImpl implementation of AuthGoogleService
type AuthGoogleServiceImpl struct {
	dig.In
	Config
	Oauth2Config *oauth2.Config
}

// NewAuthGoogleService return new instance of AuthGoogleService
func NewAuthGoogleService(impl AuthGoogleServiceImpl) AuthGoogleService {
	return &impl
}

// NewOauth2Config return new instance of oauth2.Config
func NewOauth2Config(config Config) *oauth2.Config {
	c := oauth2.Config{
		RedirectURL:  config.Callback,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"}, // TODO: put to module
		Endpoint:     google.Endpoint,
	}
	return &c
}

func (c *AuthGoogleServiceImpl) GetAuthCodeURL(ce echo.Context, cookieSecure bool) (authCodeURL string) {
	// Create oauthState cookie
	oauthState := c.setRandomCookie(ce, "oauthstate", time.Now().Add(OAuthStateCookieExpire), cookieSecure)

	// AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	// validate that it matches the the state query parameter on your redirect callback.
	urlAuthCode := c.Oauth2Config.AuthCodeURL(oauthState)

	return urlAuthCode
}

// VerifyCallback to add metaTag
func (c *AuthGoogleServiceImpl) VerifyCallback(ce echo.Context, jwtSecret string) (string, error) {
	oauthState, err := ce.Cookie("oauthstate")
	if err != nil {
		return "", errors.Trace(err)
	}

	if ce.QueryParam("state") != oauthState.Value {
		return "", errors.New("invalid oauth google state")
	}

	userInfoResp, err := c.getUserInfoFromGoogle(ce.QueryParam("code"))
	if err != nil {
		return "", errors.Trace(err)
	}

	err = c.validateUserInfoResp(userInfoResp)
	if err != nil {
		return "", errors.Trace(err)
	}

	jwtToken, err := c.generateJwtToken(userInfoResp, jwtSecret)
	if err != nil {
		return "", errors.Trace(err)
	}

	return jwtToken, nil
}

func (c *AuthGoogleServiceImpl) setRandomCookie(ce echo.Context, cookieName string, expiration time.Time, cookieSecure bool) string {
	randomVal := generateRandomBase64(64)
	cookie := &http.Cookie{Name: cookieName, Value: randomVal, Expires: expiration, HttpOnly: true, Secure: cookieSecure}
	ce.SetCookie(cookie)
	return randomVal
}

func (c *AuthGoogleServiceImpl) getUserInfoFromGoogle(code string) (userInfoResp repository.GoogleOauth2UserInfoResp, err error) {
	// Use code to get token and get user info from Google.
	token, err := c.Oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, errors.Trace(err)
	}

	if !token.Valid() {
		return nil, errors.New("invalid token")
	}

	response, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken))
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&userInfoResp)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return userInfoResp, nil
}

func (c *AuthGoogleServiceImpl) validateUserInfoResp(userInfoResp repository.GoogleOauth2UserInfoResp) error {
	if verifiedEmail, ok := userInfoResp["verified_email"]; !ok || !verifiedEmail.(bool) {
		return errors.New("invalid or empty verified_email")
	}

	if c.Config.HostedDomain != "" {
		if hd, ok := userInfoResp["hd"]; !ok || hd != c.Config.HostedDomain {
			return errors.New("invalid or empty hd")
		}
	}
	return nil
}

func (c *AuthGoogleServiceImpl) generateJwtToken(userInfoResp repository.GoogleOauth2UserInfoResp, jwtSecret string) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userInfoResp["email"]
	claims["picture"] = userInfoResp["picture"]
	claims["exp"] = time.Now().Add(JwtTokenExpire).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.Trace(err)
	}

	return t, nil
}

func generateRandomBase64(keyLength int) string {
	b := make([]byte, keyLength)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
