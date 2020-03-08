package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/purwantogz/go-impact/config"
	"github.com/purwantogz/go-impact/domain/user/interfaces"
	"github.com/purwantogz/go-impact/helpers"
	"github.com/purwantogz/go-impact/models"
)

//UserHandler merepresentasikan the httphandler
type userHandler struct {
	UsUsecase interfaces.IUseCase
}

//NewUserHandler construktor user handler
func NewUserHandler(e *echo.Group, us interfaces.IUseCase) {
	handler := &userHandler{
		UsUsecase: us,
	}

	e.POST("/users", handler.Store)
}

//Store asasasa
func (h *userHandler) Store(c echo.Context) error {
	var newUser models.User
	err := c.Bind(&newUser)
	if err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	if err = c.Validate(newUser); err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = h.UsUsecase.Store(ctx, &newUser)
	if err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	successResponse := map[string]interface{}{
		"code":    1,
		"message": "success store data",
		"data":    newUser,
	}

	return c.JSONPretty(http.StatusCreated, successResponse, config.JSONTab)
}
