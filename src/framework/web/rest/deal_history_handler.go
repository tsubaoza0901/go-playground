package rest

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/presenters"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DealHistoryHandler ...
type DealHistoryHandler struct {
	dealHistoryController *controllers.DealHistory
	*presenters.DealHistory
}

// NewDealHistoryHandler ...
func NewDealHistoryHandler(dc *controllers.DealHistory, dhp *presenters.DealHistory) *DealHistoryHandler {
	return &DealHistoryHandler{dc, dhp}
}

// GetDealHistoryList ...
func (h *DealHistoryHandler) GetDealHistoryList(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.RetrieveDealHistories)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.dealHistoryController.GetDealHistoryList(ctx, req)
	if h.Error != nil {
		return c.JSON(http.StatusBadGateway, h.Error.Text())
	}
	return c.JSON(http.StatusOK, h.AppResponse.Data)
}
