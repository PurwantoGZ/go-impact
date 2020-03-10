package usecase

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/purwantogz/go-impact/domain/img/interfaces"
	"github.com/purwantogz/go-impact/models"
)

type uploadUsecase struct {
	contextTimeout time.Duration
}

//Init ...
func Init(timeout time.Duration) interfaces.IUploadUsecase {
	return &uploadUsecase{
		contextTimeout: timeout,
	}
}

//Save Save(context.Context, *models.ImageFile) error
func (uc *uploadUsecase) Save(c context.Context, m *models.ImageFile) error {
	_, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	err := ioutil.WriteFile(m.Path, m.Data, 0666)
	if err != nil {
		return err
	}
	return nil
}
