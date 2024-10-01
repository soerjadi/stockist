package token

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(userID int64, tokenSecret string) (string, error) {
	tokenClaims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	}

	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return accessTokenClaims.SignedString([]byte(tokenSecret))
}

func GenerateRefreshToken(userID int64, tokenSecret string) (string, error) {
	refreshClaims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
	}

	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	return refreshTokenClaims.SignedString([]byte(tokenSecret))
}
