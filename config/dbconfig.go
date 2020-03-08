package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

//DbConfig config DB Configuration
type DbConfig struct {
	Driver   string
	Host     string
	Port     int
	Catalog  string
	Username string
	Password string
}

//Load Func LoadDriver(string driver)
func Load(driver string) (*DbConfig, error) {
	if driver == "" {
		return &DbConfig{}, errors.New("driver must be initialized")
	}

	err := godotenv.Load()
	if err != nil {
		return &DbConfig{}, errors.New("error loading .env files")
	}

	JwtConf = &JWTConfig{
		Issuer: os.Getenv("APP_NAME"),
		Key:    os.Getenv("JWT_SECRET"),
	}

	if driver == "mysql" {
		port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

		dbConf := &DbConfig{
			Driver:   os.Getenv("DB_CONNECTION"),
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			Catalog:  os.Getenv("DB_DATABASE"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		}
		return dbConf, nil
	}
	return &DbConfig{}, nil
}

//InitGormDB func InitDB(config)
func (cf *DbConfig) InitGormDB() (*gorm.DB, error) {

	connString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cf.Username, cf.Password, cf.Host, cf.Catalog)

	db, err := gorm.Open(cf.Driver, connString)

	if err != nil {
		log.Fatal(err.Error())
	}
	return db, nil
}
