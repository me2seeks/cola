package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/me2seeks/cola/config"
)

type Claims struct {
	jwt.RegisteredClaims
}

func GenerateToken(uid int64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Cfg.JWT.ExpiresTime) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Cfg.JWT.NotBefore) * time.Second)),
			Issuer:    config.Cfg.JWT.Issuer,
			Subject:   strconv.FormatInt(uid, 10),
			// ID:        "1",
			Audience: config.Cfg.JWT.Audience,
		},
	}).SignedString([]byte(config.Cfg.JWT.SigningKey))
}

func ParseToken(token string) (*Claims, error) {
	// Parse the token
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Cfg.JWT.SigningKey), nil
	})
	if !t.Valid || err != nil {
		return nil, fmt.Errorf("token invalid")
	}
	claims, ok := t.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, err
}
