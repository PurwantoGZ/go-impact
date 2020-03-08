package repository

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/purwantogz/go-impact/domain/user/interfaces"
	"github.com/purwantogz/go-impact/migrations"
	"github.com/purwantogz/go-impact/models"
)

type userRepository struct {
	DB *gorm.DB
}

//Init Init(db *gorm.DB) interfaces.IRepository
func Init(db *gorm.DB) interfaces.IRepository {
	return &userRepository{
		DB: db,
	}
}

//Store (repo *userRepository) Store(ctx context.Context, a *models.User) error
func (repo *userRepository) Store(ctx context.Context, a *models.User) error {
	tx := repo.DB.BeginTx(ctx, &sql.TxOptions{})
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&migrations.User{
		Email:     a.Email,
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
