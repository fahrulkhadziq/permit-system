package handler

import (
	"net/http"
	"permit-license/internal/constants"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type PermitLicenseHandler struct {
	Service service.PermitLicenseService
}

func (h *PermitLicenseHandler) Create(c echo.Context) error {
	var req dto.CreatePermitLicenseRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
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

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": "File is required",
			})
	}

	user := c.Get("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)

	unitID := claims["unit_id"].(string)

	err = h.Service.Create(
		req,
		fileHeader,
		userID,
		unitID,
	)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			})
	}

	return c.JSON(http.StatusCreated,
		map[string]interface{}{
			"message": "permit created",
		},
	)
}

func (h *PermitLicenseHandler) FindAll(c echo.Context) error {

	var params dto.QueryParams
	c.Bind(&params)

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	userID := claims["user_id"].(string)
	unitID := claims["unit_id"].(string)

	// USER UNIT
	if role == "USER_UNIT" {

		params.UploadedBy =
			userID
	}

	// HEAD UNIT
	if role == "HEAD_UNIT" {

		params.UnitID =
			unitID

		params.StatusCode =
			constants.StatusWaitingApproval
	}

	// DIRECTOR
	if role == "DIRECTOR" {

		params.StatusCode =
			constants.StatusWaitingDirectorApproval
	}

	data, totalRows, err := h.Service.FindAll(params)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			})
	}

	var response []dto.PermitLicenseResponse

	for _, item := range data {
		response = append(response, helper.ToPermitLicenseResponse(item))
	}

	page, limit := helper.NormalizePagination(params.Page, params.Limit)

	result := helper.Paginate(
		page,
		limit,
		totalRows,
		response,
	)

	return c.JSON(http.StatusOK, result)
}

func (h *PermitLicenseHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	data, err := h.Service.FindByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			})
	}

	response := helper.ToPermitLicenseDetailsResponse(data)

	return c.JSON(http.StatusOK, response)
}

func (h *PermitLicenseHandler) Download(c echo.Context) error {
	id := c.Param("id")

	data, err := h.Service.FindByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			})
	}

	return c.File(data.FileURL)
}
