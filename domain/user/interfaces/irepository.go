package interfaces

import (
	"context"

	"github.com/purwantogz/go-impact/models"
)

//IRepository for user repo
type IRepository interface {
	//Fetch(ctx context.Context, cursor string, num int64) (res []*models.User, nextCursor string, err error)
	//GetByID(ctx context.Context, id int64) (*models.User, error)
	//GetByTitle(ctx context.Context, title string) (*models.User, error)
	//Update(ctx context.Context, ar *models.User) error
	Store(ctx context.Context, a *models.User) error
	//Delete(ctx context.Context, id int64) error
}
