package impl

import (
	//	"base/internal/usecases/model"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var (
	blacklistedTokens = make(map[string]bool)
	keyOrders         = []string{}
)

func (u *usecase) LogoutToken(ctx context.Context, token string) (err error) {
	blacklistedTokens[token] = true

	keyOrders = append(keyOrders, token)

	// If the number of keys exceeds maxKeys, remove the oldest key
	if len(keyOrders) > maxKeys {
		oldestKey := keyOrders[0]
		delete(blacklistedTokens, oldestKey)
		keyOrders = keyOrders[1:]
	}

	return
}

func (u *usecase) IsTokenBlacklisted(token string) bool {
	_, blacklisted := blacklistedTokens[token]
	return blacklisted
}

func (u *usecase) CheckToken(ctx context.Context, reqToken string) (claims jwt.MapClaims, isBool bool, err error) {
	if reqToken == "" {
		return claims, false, errors.New("error, empty token")
	}

	if u.IsTokenBlacklisted(reqToken) {
		return claims, false, errors.New("token is blacklisted")
	}

	token, err := jwt.Parse(
		strings.Replace(reqToken, "Bearer ", "", 1),
		func(token *jwt.Token) (interface{}, error) {
			return []byte(u.cfg.TokenSecret), nil
		})

	if err != nil {
		return claims, false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims.VerifyExpiresAt(0, true)
		return claims, true, nil
	}
	return claims, false, errors.New("invalid JWT Token")
}
func (u *usecase) CheckTokenReset(ctx context.Context, reqToken string) (claims jwt.MapClaims, isBool bool, err error) {
	if reqToken == "" {
		return claims, false, errors.New("error, empty token")
	}
	secretKey := u.cfg.TokenSecret
	token, err := jwt.Parse(
		strings.Replace(reqToken, "Bearer ", "", 1),
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

	if err != nil {
		return claims, false, err
	}
	fmt.Println(token.Valid)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims.VerifyExpiresAt(0, true)
		return claims, true, nil
	}
	return claims, false, errors.New("invalid JWT Token")
}
