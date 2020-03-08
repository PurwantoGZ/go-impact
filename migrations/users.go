package migrations

import "github.com/jinzhu/gorm"

//User migrations for user
type User struct {
	gorm.Model
	Email     string `gorm:"unique;size:100;not null"`
	FirstName string `gorm:"not null;size:30"`
	LastName  string `gorm:"not null;size:30"`
}
