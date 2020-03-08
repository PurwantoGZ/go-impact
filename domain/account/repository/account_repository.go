package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/purwantogz/go-impact/domain/account/interfaces"
	"github.com/purwantogz/go-impact/migrations"
	"github.com/purwantogz/go-impact/models"
)

type accountRepository struct {
	Db *gorm.DB
}

//Init Init(db *gorm.DB) interfaces.IRepository
func Init(db *gorm.DB) interfaces.IRepository {
	return &accountRepository{
		Db: db,
	}
}

func (ac *accountRepository) Login(username, password string) (*models.Roles, error) {

	//Get user by email
	user := &migrations.User{}
	err := ac.Db.Model(&migrations.User{}).Where("email=?", username).Take(user).Error
	if err != nil {
		return &models.Roles{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &models.Roles{}, errors.New("User not found")
	}

	//Get Roles
	roles := &migrations.Role{}
	err = ac.Db.Model(&migrations.Role{}).Where("email=?", user.Email).Take(roles).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return &models.Roles{
			RoleType: "user",
			Scope: models.Scopes{
				Create: false,
				Delete: false,
				Edit:   false,
				Read:   true,
			},
		}, nil
	}

	//when user role found
	return &models.Roles{
		RoleType: roles.RoleType,
		Scope: models.Scopes{
			Create: roles.Create,
			Read:   roles.Read,
			Delete: roles.Delete,
			Edit:   roles.Edit,
		},
	}, nil

}
