package jwt

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Decode decodes an authorization header and returns a map of the encoded data inside
func Decode(authorizationHeader string, publicKey string) (map[string]interface{}, error) {
	headerParts := strings.Split(authorizationHeader, " ")

	if len(headerParts) != 2 {
		return nil, NewIncorrectHeaderFormat()
	}

	token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodRS256.Name {
			return nil, NewIncorrectAlgorithmError(token.Method.Alg())
		}

		rsaKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		if err != nil {
			return nil, err
		}

		return rsaKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, NewFailedToDecodeClaims()
}
