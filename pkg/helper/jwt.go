package helper

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/racibaz/go-arch/pkg/config"
)

var jwtKey []byte

func InitJWT(key string) {
	jwtKey = []byte(key)
}

type CustomClaims struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Platform string `json:"X-Platform"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, name, platform string) (string, error) {
	if len(jwtKey) == 0 {
		return "", errors.New("jwt key not initialized")
	}

	if platform != "web" && platform != "mobile" {
		return "", errors.New("invalid platform for token")
	}

	timeout := time.Duration(0)
	config := config.Get()

	if platform == PlatformWeb {
		timeout = time.Duration(config.App.JWTWebTimeout)
	} else if platform == PlatformMobile {
		timeout = time.Duration(config.App.JWTMobileTimeout)
	}

	claims := &CustomClaims{
		UserID:   userID,
		Name:     name,
		Platform: platform,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(timeout * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprint(userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyJWT(tokenStr string) (string, string, string, error) {
	if len(jwtKey) == 0 {
		return "", "", "", errors.New("jwt key not initialized")
	}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&CustomClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtKey, nil
		},
	)
	if err != nil {
		return "", "", "", err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return "", "", "", errors.New("invalid claims type")
	}

	if claims.UserID == "" || claims.Name == "" {
		return "", "", "", errors.New("invalid user claims")
	}

	if claims.Platform != "web" && claims.Platform != "mobile" {
		return "", "", "", errors.New("invalid platform claim")
	}

	return claims.UserID, claims.Name, claims.Platform, nil
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
