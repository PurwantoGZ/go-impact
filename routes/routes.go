package routes

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_accountHanlder "github.com/purwantogz/go-impact/domain/account/delivery/http"
	_accountRepo "github.com/purwantogz/go-impact/domain/account/repository"
	_accountUsecase "github.com/purwantogz/go-impact/domain/account/usecase"
	_userHanlder "github.com/purwantogz/go-impact/domain/user/delivery/http"
	_userRepo "github.com/purwantogz/go-impact/domain/user/repository"
	_userUsecase "github.com/purwantogz/go-impact/domain/user/usecase"
	validator "gopkg.in/go-playground/validator.v9"
)

//CustomValidator struct custom Validator
type CustomValidator struct {
	validator *validator.Validate
}

//Validate (cv *CustomValidator) Validate(i interface{})error
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

//Endpoints EndPoints(e *echo.Echo)
func Endpoints(e *echo.Echo, db *gorm.DB) {

	//Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Routes
	g := e.Group("/v1")
	// Middleware
	g.Use(middleware.Logger())
	g.Use(middleware.Recover())

	//CORS
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	timeoutContext := time.Duration(1000) * time.Second

	//User Repo & User UseCase
	userRepo := _userRepo.Init(db)
	userUsecase := _userUsecase.Init(userRepo, timeoutContext)
	_userHanlder.NewUserHandler(g, userUsecase)

	//Account Repo & Usecase
	accountRepo := _accountRepo.Init(db)
	accountUsecase := _accountUsecase.Init(accountRepo, timeoutContext)
	_accountHanlder.NewAccountHandler(e, accountUsecase)

}
