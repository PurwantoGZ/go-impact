package usecase

import (
	"context"
	"time"

	"github.com/purwantogz/go-impact/config"
	"github.com/purwantogz/go-impact/domain/account/interfaces"
	"github.com/purwantogz/go-impact/domain/token/factory"
)

type accountUsecase struct {
	accountRepo    interfaces.IRepository
	contextTimeout time.Duration
}

//Init constructor account usecase
func Init(ap interfaces.IRepository, timeout time.Duration) interfaces.IUseCase {
	return &accountUsecase{
		accountRepo:    ap,
		contextTimeout: timeout,
	}
}

//Login Login(ctx context.Context, username, password string) (map[string]string, error)
func (au *accountUsecase) Login(c context.Context, username, password string) (map[string]string, error) {
	_, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	roles, err := au.accountRepo.Login(username, password)
	if err != nil {
		return map[string]string{
			"code":    "0",
			"message": err.Error(),
		}, err
	}

	tokenFactory := factory.New(config.JwtConf.Key, config.JwtConf.Issuer, 30)
	token, err := tokenFactory.Build(username, roles)

	if err != nil {
		return map[string]string{
			"code":    "0",
			"message": err.Error(),
		}, err
	}

	return token, nil

}
