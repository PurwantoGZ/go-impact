package http

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/purwantogz/go-impact/config"
	"github.com/purwantogz/go-impact/domain/img/interfaces"
	"github.com/purwantogz/go-impact/helpers"
	"github.com/purwantogz/go-impact/models"
)

type imgHandler struct {
	ImgUsecase interfaces.IUploadUsecase
}

//NewImgHandler ...
func NewImgHandler(e *echo.Echo, us interfaces.IUploadUsecase) {
	handler := &imgHandler{
		ImgUsecase: us,
	}

	e.POST("/upload", handler.Save)
}

//Save(context.Context, *models.ImageFile) error
func (h *imgHandler) Save(c echo.Context) error {

	file, header, err := c.Request().FormFile("file")
	if err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	defer file.Close()

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	wd, err := os.Getwd()
	if err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	path := wd + "/uploads/" + header.Filename

	randomValue, _ := randomHex(12)
	imgData := &models.ImageFile{
		ID:   randomValue,
		Data: data,
		Name: header.Filename,
		Path: path,
	}

	h.ImgUsecase.Save(ctx, imgData)
	if err != nil {
		response, responsecode := helpers.ResponseErr(err, "ERR_GENERAL")
		return c.JSONPretty(responsecode, response, config.JSONTab)
	}

	successResponse := map[string]interface{}{
		"id": imgData.ID,
	}

	return c.JSONPretty(http.StatusCreated, successResponse, config.JSONTab)
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
