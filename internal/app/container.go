package app

import (
	"MyNote/internal/domain/repository"
	"MyNote/internal/handlers"
	"MyNote/internal/infrastructure/config"
	"MyNote/internal/infrastructure/db"
	_repo "MyNote/internal/infrastructure/db/repository"
	"MyNote/internal/usecases"
	"MyNote/pkg/validation"
	"time"

	"github.com/go-playground/validator/v10"
)

type Container struct {
	AuthHandler *handlers.AuthHandler
	Validator   *validator.Validate
}

func BuildContainer(cfg *config.Config) (*Container, error) {

	db.ConnectDB(cfg)
	var userRepo repository.IUserRepository = _repo.NewUserRepository(db.DB)
	jwtService := usecases.NewJWTService(
		cfg.JWT.Secret,
		time.Duration(cfg.JWT.AccesTokenTTLMinute),
		time.Duration(cfg.JWT.RefreshTokenTTLMinute),
	)
	authService := usecases.NewAuthService(userRepo, jwtService)
	v := validator.New()
	v.RegisterValidation("password", validation.PasswordValidator)

	authHandler := handlers.NewAuthHandler(authService, v)

	return &Container{
		AuthHandler: authHandler,
		Validator:   v,
	}, nil

}
