package migrations

import "github.com/jinzhu/gorm"

//Token gorm table for Token
type Token struct {
	gorm.Model
	Email       string `gorm:"type:varchar(100);not null;unique"`
	Type        string `gorm:"not null;size:10"`
	AccessToken string `gorm:"type:varchar(1000);not null"`
	Refresh     string `gorm:"type:varchar(1000);not null"`
	ExpiresIn   int64  `gorm:"not null"`
}
