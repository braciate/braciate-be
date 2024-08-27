package jwt

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func Sign(Data map[string]interface{}, secretEnvKey string, ExpiredAt time.Duration) (string, error) {
	expiredAt := time.Now().Add(ExpiredAt).Unix()

	JWTSecretKey := os.Getenv(secretEnvKey)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(JWTSecretKey))

	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}
