package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/davidwang/go-finance-api/go-finance-api/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// JWTConfig holds configuration for JWT tokens
type JWTConfig struct {
	Secret     string
	Issuer     string
	Expiration time.Duration // in minutes
}

// DefaultConfig returns a default JWT configuration
func DefaultConfig() JWTConfig {
	return JWTConfig{
		Secret:     "your-jwt-secret-key", // Will be overridden by config
		Issuer:     "finance-api",
		Expiration: 60, // 60 minutes
	}
}

// Claims represents the JWT claims
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(user *models.User, config JWTConfig) (string, error) {
	// Create the JWT claims
	expirationTime := time.Now().Add(time.Minute * config.Expiration)
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.Issuer,
			Subject:   user.ID.String(),
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString string, config JWTConfig) (*Claims, error) {
	// Parse the token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// GenerateRefreshToken creates a refresh token with longer expiration
func GenerateRefreshToken(userID uuid.UUID, config JWTConfig) (string, error) {
	// Create a longer expiration time for refresh tokens
	expirationTime := time.Now().Add(time.Hour * 24 * 7) // 7 days

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    config.Issuer,
		Subject:   userID.String(),
		ID:        uuid.New().String(), // Use a unique ID for each refresh token
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
