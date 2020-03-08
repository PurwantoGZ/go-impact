package migrations

import "github.com/jinzhu/gorm"

//Role migration for user roles
type Role struct {
	gorm.Model
	Email    string `gorm:"unique;not null;size:100"`
	RoleType string `gorm:"not null;size:15"`
	Create   bool   `gorm:"not null"`
	Read     bool   `gorm:"not null"`
	Edit     bool   `gorm:"not null"`
	Delete   bool   `gorm:"not null"`
}
