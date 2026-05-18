package handler

import (
	"net/http"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/service"

	"github.com/labstack/echo/v4"
)

type MasterDocumentHandler struct {
	Service service.MasterDocumentService
}

func (h *MasterDocumentHandler) Create(c echo.Context) error {
	var req dto.MasterDocumentRequest

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

	err = h.Service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]interface{}{
				"message": err.Error(),
			})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Master document created successfully",
	})
}

func (h *MasterDocumentHandler) FindAll(c echo.Context) error {
	var params dto.QueryParams

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			})
	}
	masterDocs, totalRows, err := h.Service.FindAll(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]interface{}{
				"message": err.Error(),
			})
	}

	var response []dto.MasterDocumentResponse
	for _, item := range masterDocs {
		response = append(response, item)
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

func (h *MasterDocumentHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	masterDoc, err := h.Service.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			})
	}
	return c.JSON(http.StatusOK, masterDoc)
}

func (h *MasterDocumentHandler) Update(
	c echo.Context,
) error {

	id := c.Param("id")

	var req dto.UpdateMasterDocumentRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "invalid request",
			},
		)
	}

	if err := helper.Validate.Struct(req); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	err := h.Service.Update(id, req)

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
			"message": "success update master document",
		},
	)
}

func (h *MasterDocumentHandler) Delete(
	c echo.Context,
) error {

	id := c.Param("id")

	err := h.Service.Delete(id)

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
			"message": "success delete master document",
		},
	)
}
