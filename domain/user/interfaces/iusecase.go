package interfaces

import (
	"context"

	"github.com/purwantogz/go-impact/models"
)

//IUseCase user use case
type IUseCase interface {
	//Fetch(ctx context.Context, cursor string, num int64) ([]*models.User, string, error)
	//GetByID(ctx context.Context, id int64) (*models.User, error)
	//Update(ctx context.Context, ar *models.User) error
	//GetByTitle(ctx context.Context, title string) (*models.User, error)
	Store(context.Context, *models.User) error
	//Delete(ctx context.Context, id int64) error
}
