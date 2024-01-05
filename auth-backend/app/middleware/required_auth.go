package middleware

import (
	"auth-backend/app/constant"
	"auth-backend/app/domain/dco"
	"auth-backend/app/pkg"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequiredAuth() gin.HandlerFunc {
	// This middleware will be used for routes that require authentication
	// It must be implemented in the /me route
	return func(c *gin.Context) {
		defer pkg.PanicHandler(c)

		// Get the token from the header
		tokenString := c.Request.Header.Get("Authorization")

		// If the token is empty, return with status 401
		if tokenString == "" {
			pkg.PanicException(constant.Unauthorized)
		}

		// decode and validate the token
		token, err := DecodeToken(tokenString)
		if err != nil {
			pkg.PanicException(constant.Unauthorized)
		}

		// check if the token is expired
		if token.ExpiresAt.Time.Before(time.Now()) {
			slog.Error("Error happened: when parse token", "error", "token is expired")
			pkg.PanicException(constant.Unauthorized)
		}

		// Set the retrieved token to the context
		c.Set("retrievedToken", token)
		c.Next()
	}
}

func DecodeToken(tokenString string) (*dco.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dco.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			message := fmt.Sprintf("unexpected signing method %s", token.Header["alg"])
			return nil, errors.New(message)
		}

		return dco.JWTSigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*dco.JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claim, nil
}
