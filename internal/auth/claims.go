package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

func getClaims(ctx context.Context) (jwt.MapClaims, bool) {
	claims, ok := ctx.Value(ContextUserKey).(jwt.MapClaims)

	return claims, ok
}

func GetClaimString (ctx context.Context, key string) (string, bool) {
	claims, ok := getClaims(ctx)
	if !ok {
		return "", false
	}

	val, exist := claims[key]
	if !exist {
		return "", false
	}

	strVal, ok := val.(string)

	return strVal, ok
}