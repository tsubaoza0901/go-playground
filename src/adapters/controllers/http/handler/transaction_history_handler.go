package handler

import (
	"go-playground/m/v1/src/adapters/controllers/http/handler/request"
	"go-playground/m/v1/src/adapters/controllers/http/handler/response"
	"go-playground/m/v1/src/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TransactionHistoryHandler ...
type TransactionHistoryHandler struct {
	usecase.ITransactionUsecase
}

// NewTransactionHistoryHandler ...
func NewTransactionHistoryHandler(p usecase.ITransactionUsecase) TransactionHistoryHandler {
	return TransactionHistoryHandler{p}
}

// RetrieveTransactionHistories ...
func (h TransactionHistoryHandler) RetrieveTransactionHistories(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.RetrieveTransactionHistories)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	payments, err := h.RetrieveTransactionHistoriesByUserID(ctx, req.UserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response.NewTransactionHistories(payments))
}
