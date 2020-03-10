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
	//mailServe := mail.Build("purwanto.dev@gmail.com", "test_go_email", "asdhs shda shds dhdas")
	// mailServe := mail.BuildWithHtml("purwanto.dev@gmail.com", "test_go_email")
	// err = mailServe.Send()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	e := echo.New()
	routes.Endpoints(e, db)
	e.Logger.Fatal(e.Start(":8000"))

}
