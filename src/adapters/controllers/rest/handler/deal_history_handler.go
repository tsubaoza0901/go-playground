package handler

import (
	"go-playground/m/v1/adapters/controllers/rest/handler/request"
	"go-playground/m/v1/adapters/controllers/rest/handler/response"
	"go-playground/m/v1/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DealHistoryHandler ...
type DealHistoryHandler struct {
	usecase.IDealUsecase
}

// NewDealHistoryHandler ...
func NewDealHistoryHandler(p usecase.IDealUsecase) DealHistoryHandler {
	return DealHistoryHandler{p}
}

// GetDealHistoryList ...
func (h DealHistoryHandler) GetDealHistoryList(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.RetrieveDealHistories)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	dealHistories, err := h.RetrieveDealHistoriesByUserID(ctx, req.UserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response.NewDealHistories(dealHistories))
}
