package interfaces

import (
	"context"

	"github.com/purwantogz/go-impact/models"
)

//IUploadUsecase ...
type IUploadUsecase interface {
	Save(c context.Context, m *models.ImageFile) error
}
