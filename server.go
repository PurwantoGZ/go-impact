package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/purwantogz/go-impact/config"
	"github.com/purwantogz/go-impact/routes"
)

func main() {

	dbConfig, err := config.Load("mysql")

	db, err := dbConfig.InitGormDB()

	if err != nil {
		fmt.Println(err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//db.AutoMigrate(&migrations.User{}, &migrations.Token{}, &migrations.Role{})

	e := echo.New()
	routes.Endpoints(e, db)
	e.Logger.Fatal(e.Start(":8000"))

}
