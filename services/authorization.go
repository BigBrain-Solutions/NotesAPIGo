package services

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Id    string
	Scope string
}

const KEY = "qwertyuiopASDFGHJKL1234567890"

func IsAuthorized(value string) bool {
	tokenStr := strings.Split(value, "Bearer ")[1]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})

	if err != nil {
		return false
	}

	// Prevents from hacking the JWT and removing the algorithm in header
	if token.Method.Alg() == "None" || token.Method.Alg() == "" {
		return false
	}

	token.Signature = KEY

	claims := JwtParse(value)

	if claims.ExpiresAt < time.Now().Unix() {
		return false
	}

	return token.Valid
}

func JwtParse(value string) *Claims {
	claims := &Claims{}
	token := strings.Split(value, "Bearer ")[1]

	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})

	return claims
}
