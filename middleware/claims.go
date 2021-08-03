package middleware

import (
	"go-microservices/jwt"
	"go-microservices/microservice"
	"strings"

	"github.com/labstack/echo"
)

// GetClaim retrieves the specified claim from the authorization header JWT.
/*
	Returns nil if the header is not present.
*/
type GetClaim func(c echo.Context, claim string, defaultValue interface{}) (interface{}, error)

// BuildGetClaim builds the GetClaim middleware.
func BuildGetClaim(proxy microservice.APIProxy) GetClaim {
	return func(c echo.Context, claim string, defaultValue interface{}) (interface{}, error) {
		publicKey, err := proxy.PublicKey()
		if err != nil {
			return nil, err
		}

		authorizationHeader := c.Request().Header.Get(echo.HeaderAuthorization)

		claims, err := jwt.Decode(authorizationHeader, publicKey)
		if err != nil {
			return nil, err
		}

		paths := strings.Split(claim, ".")

		var value interface{}
		for i, path := range paths {
			if i == 0 {
				value = claims[path]
			} else {
				mapValue := value.(map[string]interface{})

				value = mapValue[path]
			}
		}

		if value == nil {
			return defaultValue, nil
		}

		return value, nil
	}
}
