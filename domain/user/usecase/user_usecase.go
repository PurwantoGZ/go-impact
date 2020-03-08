package usecase

import (
	"context"
	"time"

	"github.com/purwantogz/go-impact/domain/user/interfaces"
	"github.com/purwantogz/go-impact/models"
)

var test string

type userUsecase struct {
	userRepo       interfaces.IRepository
	contextTimeout time.Duration
}

//Init constructor user use case
func Init(u interfaces.IRepository, timeout time.Duration) interfaces.IUseCase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

//Store Store(context.Context, *models.User) error
func (uc *userUsecase) Store(c context.Context, m *models.User) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	err := uc.userRepo.Store(ctx, m)
	if err != nil {
		return err
	}
	return nil
}
