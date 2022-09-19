package handler

import (
	"go-playground/m/v1/adapters/controllers/http/handler/request"
	"go-playground/m/v1/adapters/controllers/http/handler/response"
	"go-playground/m/v1/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BalanceControlHandler ...
type BalanceControlHandler struct {
	usecase.IBalanceControlUsecase
}

// NewBalanceControlHandler ...
func NewBalanceControlHandler(u usecase.IBalanceControlUsecase) BalanceControlHandler {
	return BalanceControlHandler{u}
}

// Pay ...
func (h BalanceControlHandler) Pay(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.Payment)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.PayMoney(ctx, req.UserID, req.ConvertToPaymentInput()); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "OK")
}

// TopUp ...
func (h BalanceControlHandler) TopUp(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.PuttingMoney)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.PutMoney(ctx, req.UserID, req.ConvertToPuttingMoneyInput()); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "OK")
}

// GetRemainingBalance ...
func (h BalanceControlHandler) GetRemainingBalance(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.RetrieveRemainingBalance)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	balance, err := h.RetrieveRemainingBalanceByUserID(ctx, req.UserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response.NewBalance(balance))
}
