package rest

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/presenters"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GradeHandler ...
type GradeHandler struct {
	gradeController *controllers.Grade
	*presenters.Grade
}

// NewGradeHandler ...
func NewGradeHandler(gc *controllers.Grade, gp *presenters.Grade) *GradeHandler {
	return &GradeHandler{gc, gp}
}

// GetGradeList ...
func (h *GradeHandler) GetGradeList(c echo.Context) error {
	ctx := c.Request().Context()

	h.gradeController.GetGradeList(ctx)
	if h.Error != nil {
		return c.JSON(http.StatusBadGateway, h.Error.Text())
	}
	return c.JSON(http.StatusOK, h.AppResponse.Data)
}
