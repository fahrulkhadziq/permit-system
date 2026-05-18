package handler

import (
	"net/http"
	"permit-license/internal/dto"
	"permit-license/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ApprovalHandler struct {
	Service service.ApprovalService
}

func (h *ApprovalHandler) Approve(c echo.Context) error {
	id := c.Param("id")

	var req dto.ApprovalRequest
	if err := c.Bind(&req); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "invalid request",
			},
		)
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)
	unitID := claims["unit_id"].(string)

	err := h.Service.Approve(id, userID, role, unitID, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK,
		map[string]interface{}{
			"message": "success approve document",
		},
	)
}

func (h *ApprovalHandler) Reject(c echo.Context) error {
	id := c.Param("id")

	var req dto.ApprovalRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "invalid request",
			},
		)
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)
	unitID := claims["unit_id"].(string)

	err := h.Service.Reject(id, userID, role, unitID, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK,
		map[string]interface{}{
			"message": "success reject document",
		},
	)
}
