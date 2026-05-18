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
	var unitID *string

	if role != "DIRECTOR" {
		unit := claims["unit_id"].(string)
		unitID = &unit
	}

	data, err := h.Service.GetStatistics(unitID)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}
	return c.JSON(http.StatusOK, data)
}
