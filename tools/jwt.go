package tools

import (
	"dms/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims struct represents the payload of the JWT token.
// As the JWT are for User and Device, then claims includes standard fields such as Id and AdditionalInfo,
// along with the embedded RegisteredClaims struct that provides standard JWT claims.
type Claims struct {
	Id             string `json:"user_id"`
	AdditionalInfo string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token for a authentication with a specific ID, Additional Info, and Expired Duration.
// The token includes a payload (Claims) with information and expiration time.
// The token is signed using a secret key and the HS256 signing method.
func GenerateToken(id string, additionalInfo string, expirationDuration time.Duration) (string, error) {
	cfg := config.LoadConfig()
	expirationTime := time.Now().Add(expirationDuration)

	claims := &Claims{
		Id:             id,
		AdditionalInfo: additionalInfo,
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
