package dco

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var JWTSigningKey = []byte(os.Getenv("GATEWAY_JWT_SIGNING_KEY"))
var JWTExpirationTime = 12 * time.Hour

type JWTClaim struct {
	Username    string
	Department  string
	Permissions []string
	IsAdmin     bool

	jwt.RegisteredClaims
}

type AuthResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
