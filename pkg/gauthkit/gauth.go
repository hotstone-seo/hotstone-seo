package gauthkit

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

// UserInfo from google auth
type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Hd            string `json:"hd"`
}

// RetrieveUserInfo to retrieve user info from google auth
func RetrieveUserInfo(ctx context.Context, token *oauth2.Token) (*UserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken))
	if err != nil {
		return nil, fmt.Errorf("gauth: %w", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	var userInfo UserInfo
	if err = json.Unmarshal(b, &userInfo); err != nil {
		return nil, fmt.Errorf("gauth: %w", err)
	}

	return &userInfo, nil
}
