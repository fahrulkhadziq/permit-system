package handler

import (
	"net/http"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/service"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service service.AuthService
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := helper.Validate.Struct(req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	err := h.Service.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "register success",
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request",
		})
	}

	if err := helper.Validate.Struct(req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	response, err := h.Service.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {

	var req dto.RefreshTokenRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "invalid request",
			},
		)
	}

	token, err :=
		h.Service.RefreshToken(
			req.RefreshToken,
		)

	if err != nil {

		return c.JSON(
			http.StatusUnauthorized,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"access_token": token,
		},
	)
}

func (h *AuthHandler) Logout(c echo.Context) error {

	var req dto.LogoutRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "invalid request",
			},
		)
	}

	err := h.Service.Logout(
		req.RefreshToken,
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
		map[string]interface{}{
			"message": "logout success",
		},
	)
}
