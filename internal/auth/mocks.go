package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TestAuthenticator struct{}

const testSecret = "test-secret"

var testClaims = jwt.MapClaims{
	"aud": "test-audience",
	"exp": time.Now().Add(time.Hour).Unix(),
	"iss": "test-issuer",
	"sub": int64(12345),
}

func (a *TestAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	if claims == nil {
		claims = testClaims
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(testSecret))
}

func (a *TestAuthenticator) ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(testSecret), nil
	})
}
