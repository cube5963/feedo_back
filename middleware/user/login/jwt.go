package login

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("your_secret_key")

func init() {
	// Ensure 'secret' is referenced to avoid unused variable error.
	_ = secret
}

func generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "issuer",
		Subject:   "subject",
		Audience:  []string{"audience"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7*24*time.Hour)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        "id",
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verify(tokenString string) bool {
    token, err := jwt.ParseWithClaims(
        tokenString,
        &jwt.RegisteredClaims{},
        func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secret, nil
        },
    )
    if err != nil {
        return false
    }

    if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
        fmt.Println(claims.Issuer)
        return true
    }
    return false
}
