package jwtutil

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	config "github.com/noyandey88/go-todo-app/configs"
)

// Claims is our custom JWT claims struct
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateAccessToken creates a short-lived JWT for authentication
func GenerateAccessToken(userID uint, secret string, expiresIn int) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresIn) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GenerateRefreshToken creates a long-lived JWT for refresh purposes
func GenerateRefreshToken(userID uint, secret string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GenerateResetToken creates a short-lived token for password resets
func GenerateResetToken(userID uint, secret string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), // reset link valid for 30 mins
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseAccessToken validates and returns the user ID from an access token
func ParseAccessToken(tokenString, secret string) (uint, error) {
	return parseToken(tokenString, secret)
}

// ParseResetToken validates and returns the user ID from a reset token
func ParseResetToken(tokenString, secret string) (uint, error) {
	return parseToken(tokenString, secret)
}

// Common parser function
func parseToken(tokenString, secret string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	return claims.UserID, nil
}

func ParseUserIdFromToken(token string) (uint, error) {
	// Ensure "Bearer " prefix
	if !strings.HasPrefix(strings.ToLower(token), "bearer ") {
		token = "Bearer " + token
	}

	parts := strings.Split(token, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		err := fmt.Errorf("invalid authorization header format")
		return 0, err
	}

	cfg := config.LoadConfig()
	id, err := ParseAccessToken(parts[1], cfg.JWT.Secret)
	return id, err
}
