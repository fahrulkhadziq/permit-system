package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RoleMiddleware(role string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			user := c.Get("user").(*jwt.Token)

			claims := user.Claims.(jwt.MapClaims)

			userRole := claims["role"].(string)

			if userRole != role {
				return c.JSON(http.StatusForbidden,
					map[string]interface{}{
						"message": "forbidden",
					},
				)
			}

			return next(c)
		}
	}
}
