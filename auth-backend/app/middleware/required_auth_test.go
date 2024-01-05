package middleware

import (
	"auth-backend/app/domain/dco"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TestDecodeToken(t *testing.T) {
	// Define a valid token
	validToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	}).SignedString(dco.JWTSigningKey)
	if err != nil {
		t.Fatalf("Failed to create valid token: %v", err)
	}

	// Define an invalid token
	invalidToken := "invalidToken"

	t.Run("valid token", func(t *testing.T) {
		claim, err := DecodeToken(validToken)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if claim == nil {
			t.Errorf("Expected claim to be not nil")
		}
	})

	t.Run("invalid token", func(t *testing.T) {
		_, err := DecodeToken(invalidToken)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestRequiredAuth(t *testing.T) {
	// Create a new gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Add the middleware to the router
	router.Use(RequiredAuth())

	// Add a test route
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Test with no Authorization header
	req, _ := http.NewRequest("GET", "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Test with valid token
	validToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, dco.JWTClaim{
		Username: "test",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dco.JWTExpirationTime)),
		},
	}).SignedString(dco.JWTSigningKey)
	if err != nil {
		t.Fatalf("Failed to create valid token: %v", err)
	}
	req, _ = http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", validToken)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test with expired token
	expiredToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, dco.JWTClaim{
		Username: "test",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-dco.JWTExpirationTime)),
		},
	}).SignedString(dco.JWTSigningKey)
	if err != nil {
		t.Fatalf("Failed to create valid token: %v", err)
	}
	req, _ = http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", expiredToken)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}
