package handler

import (
	"net/http"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/service"

	"github.com/labstack/echo/v4"
)

type UnitHandler struct {
	Service service.UnitService
}

func (h *UnitHandler) Create(
	c echo.Context,
) error {

	var req dto.CreateUnitRequest

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

	err := h.Service.Create(req)

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
			"message": "success create unit",
		},
	)
}

func (h *UnitHandler) FindAll(
	c echo.Context,
) error {

	var params dto.QueryParams

	c.Bind(&params)

	data, totalRows, err :=
		h.Service.FindAll(params)

	if err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	var response []dto.UnitResponse

	for _, item := range data {

		response = append(
			response,
			helper.ToUnitResponse(item),
		)
	}

	page, limit :=
		helper.NormalizePagination(
			params.Page,
			params.Limit,
		)

	result := helper.Paginate(
		page,
		limit,
		totalRows,
		response,
	)

	return c.JSON(
		http.StatusOK,
		result,
	)
}

func (h *UnitHandler) FindByID(
	c echo.Context,
) error {

	id := c.Param("id")

	data, err := h.Service.FindByID(id)

	if err != nil {

		return c.JSON(
			http.StatusNotFound,
			map[string]interface{}{
				"message": "unit not found",
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		helper.ToUnitResponse(*data),
	)
}

func (h *UnitHandler) Update(
	c echo.Context,
) error {

	id := c.Param("id")

	var req dto.UpdateUnitRequest

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
			"message": "success update unit",
		},
	)
}

func (h *UnitHandler) Delete(
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
			"message": "success delete unit",
		},
	)
}
