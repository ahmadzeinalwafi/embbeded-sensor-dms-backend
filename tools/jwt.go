package tools

import (
	"dms/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims struct represents the payload of the JWT token.
// It includes standard fields like UserId and Email,
// along with the embedded RegisteredClaims struct that provides standard JWT claims.
type Claims struct {
	UserId string `json:"user_id"`
	Email  string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token for a user with a specific ID and email.
// The token includes a payload (Claims) with user-specific information and expiration time.
// The token is signed using a secret key and the HS256 signing method.
func GenerateToken(userID string, email string) (string, error) {
	cfg := config.LoadConfig()
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		UserId: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "device-management-system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.GetString("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
