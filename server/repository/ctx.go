package repository

import (
	"context"

	"github.com/dgrijalva/jwt-go"
)

type CtxKey int

const (
	TokenCtxKey CtxKey = iota
)

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
