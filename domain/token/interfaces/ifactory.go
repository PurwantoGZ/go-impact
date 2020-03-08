package interfaces

import "github.com/purwantogz/go-impact/models"

//ITokenFactory interfaces fo token factory
type ITokenFactory interface {
	Build(email string, role *models.Roles) (map[string]string, error)
	Refresh(refreshToken string) (map[string]string, error)
}
