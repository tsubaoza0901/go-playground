package handler

import (
	"go-playground/m/v1/src/adapters/controllers/http/handler/response"
	"go-playground/m/v1/src/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GradeHandler ...
type GradeHandler struct {
	usecase.IGradeUsecase
}

// NewGradeHandler ...
func NewGradeHandler(u usecase.IGradeUsecase) GradeHandler {
	return GradeHandler{u}
}

// RetrieveAllGrades ...
func (h GradeHandler) RetrieveAllGrades(c echo.Context) error {
	ctx := c.Request().Context()

	grades, err := h.RetrieveGrades(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response.NewGrades(*grades))
}
