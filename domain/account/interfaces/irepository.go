package interfaces

import (
	"github.com/purwantogz/go-impact/models"
)

//IRepository IRepository interface
type IRepository interface {
	Login(username, password string) (*models.Roles, error)
}
