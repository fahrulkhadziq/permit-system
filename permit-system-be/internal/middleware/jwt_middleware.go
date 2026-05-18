package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHandler := c.Request().Header.Get("Authorization")

		if authHandler == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "missing token",
			},
			)
		}
		tokenString := strings.Replace(
			authHandler,
			"Bearer ",
			"",
			1,
		)

		token, err := jwt.Parse(tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid token",
			})
		}

		c.Set("user", token)
		return next(c)
	}

}
