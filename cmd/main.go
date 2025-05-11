package main

import (
	"MyNote/internal/app"
	"MyNote/internal/infrastructure/config"
	"MyNote/pkg/validation"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	v := validator.New()
	v.RegisterValidation("password", validation.PasswordValidator)

	cfg := config.Load("/media/Akunamatata/Pet-project/MyNote/config/config.yaml")
	container, err := app.BuildContainer(cfg)
	if err != nil {
		fmt.Printf("Failed to build container: %v", err)
	}
	e := echo.New()
	e.POST("/register", container.AuthHandler.Register)
	e.POST("/login", container.AuthHandler.Login)

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
