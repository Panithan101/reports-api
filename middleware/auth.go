package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// TokenData represents the token structure
type TokenData struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Exp      int64  `json:"exp"`
}

// AuthMiddleware checks for valid token in Authorization header
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "❌ Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "❌ Invalid authorization format", http.StatusUnauthorized)
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate the token
		tokenData, err := parseToken(tokenString)
		if err != nil {
			http.Error(w, "❌ Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Check if token is expired
		if time.Now().Unix() > tokenData.Exp {
			http.Error(w, "❌ Token expired", http.StatusUnauthorized)
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), "user", tokenData)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// AdminMiddleware checks if the user has admin role
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// First check authentication
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "❌ Authorization header required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "❌ Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tokenData, err := parseToken(tokenString)
		if err != nil {
			http.Error(w, "❌ Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Check if token is expired
		if time.Now().Unix() > tokenData.Exp {
			http.Error(w, "❌ Token expired", http.StatusUnauthorized)
			return
		}

		// Check if user has admin role
		if tokenData.Role != "admin" {
			http.Error(w, "❌ Admin access required", http.StatusForbidden)
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), "user", tokenData)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// parseToken validates and parses a token
func parseToken(tokenString string) (*TokenData, error) {
	// Decode base64
	tokenBytes, err := base64.StdEncoding.DecodeString(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token format")
	}

	// Parse JSON
	var tokenData TokenData
	if err := json.Unmarshal(tokenBytes, &tokenData); err != nil {
		return nil, fmt.Errorf("invalid token content")
	}

	return &tokenData, nil
}

// GetUserFromContext extracts user info from request context
func GetUserFromContext(r *http.Request) (*TokenData, bool) {
	user, ok := r.Context().Value("user").(*TokenData)
	return user, ok
}
