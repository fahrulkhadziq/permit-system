package handler

import (
	"net/http"
	"permit-license/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	Service service.DashboardService
}

func (h *DashboardHandler) GetStatistics(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	if role == "DIRECTOR" {

		data, err :=
			h.Service.GetStatisticsAll()

		if err != nil {

			return c.JSON(
				http.StatusBadRequest,
				map[string]interface{}{
					"message": err.Error(),
				},
			)
		}

		return c.JSON(
			http.StatusOK,
			data,
		)
	}
	// USER UNIT / HEAD UNIT

	unitID :=
		claims["unit_id"].(string)

	data, err :=
		h.Service.GetStatistics(
			&unitID,
			role,
		)

	if err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		data,
	)
}
