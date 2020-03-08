package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/purwantogz/go-impact/config"
	"github.com/purwantogz/go-impact/domain/account/interfaces"
	"github.com/purwantogz/go-impact/helpers"
	_logRequest "github.com/purwantogz/go-impact/request"
)

type accountHandler struct {
	AccountUsecase interfaces.IUseCase
}

//NewAccountHandler NewAccountHandler(e *echo.Echo, ausecase interfaces.IUseCase)
func NewAccountHandler(e *echo.Echo, ausecase interfaces.IUseCase) {
	handler := &accountHandler{
		AccountUsecase: ausecase,
	}

	e.POST("/account/login", handler.Login)
}

//Login login handler
func (h *accountHandler) Login(c echo.Context) error {
	var acRequest _logRequest.LoginRequest
	err := c.Bind(&acRequest)
	if err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	if err = c.Validate(acRequest); err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	resp, err := h.AccountUsecase.Login(ctx, acRequest.Username, acRequest.Password)
	if err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	successResponse := map[string]interface{}{
		"code":    1,
		"message": "success get data",
		"data":    resp,
	}

	return c.JSONPretty(http.StatusCreated, successResponse, config.JSONTab)
}
