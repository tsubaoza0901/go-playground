package rest

import (
	"go-playground/m/v1/controllers"
	"go-playground/m/v1/controllers/request"
	"go-playground/m/v1/presenters"
	"go-playground/m/v1/usecase/rule"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BalanceControlHandler ...
type BalanceControlHandler struct {
	balanceControlController *controllers.BalanceControl
	*presenters.Balance
}

// NewBalanceControlHandler ...
func NewBalanceControlHandler(bcc *controllers.BalanceControl, bp *presenters.Balance) *BalanceControlHandler {
	return &BalanceControlHandler{bcc, bp}
}

// Pay ...
func (h *BalanceControlHandler) Pay(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.Payment)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.balanceControlController.Pay(ctx, req)
	if h.Error != nil {
		switch *h.Error {
		case rule.BadRequest, rule.ShortBalance:
			return c.JSON(http.StatusBadRequest, h.Error.Text())
		case rule.NotFound:
			return c.JSON(http.StatusNotFound, h.Error.Text())
		case rule.InternalServerError:
			return c.JSON(http.StatusInternalServerError, h.Error.Text())
		}
	}
	return c.JSON(http.StatusOK, "OK")
}

// TopUp ...
func (h *BalanceControlHandler) TopUp(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.PuttingMoney)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.balanceControlController.TopUp(ctx, req)
	if h.Error != nil {
		return c.JSON(http.StatusBadRequest, h.Error.Text())
	}
	return c.JSON(http.StatusOK, "OK")
}

// GetRemainingBalance ...
func (h *BalanceControlHandler) GetRemainingBalance(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(request.RetrieveRemainingBalance)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.balanceControlController.GetRemainingBalance(ctx, req)
	if h.Error != nil {
		return c.JSON(http.StatusBadRequest, h.Error.Text())
	}
	return c.JSON(http.StatusOK, h.AppResponse.Data)
}
