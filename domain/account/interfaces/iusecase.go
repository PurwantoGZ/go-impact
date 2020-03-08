package interfaces

import "context"

//IUseCase IUsaCase interface
type IUseCase interface {
	Login(c context.Context, username, password string) (map[string]string, error)
}
