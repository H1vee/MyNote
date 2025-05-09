package usecases

import (
	"MyNote/internal/domain/models"
	"MyNote/internal/domain/repository"
	"MyNote/pkg/utils"
	"errors"
)

type AuthService struct {
	userRepo   repository.IUserRepository
	jwtService *JWTService
}

func NewAuthService(userRepo repository.IUserRepository, jwtService *JWTService) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (a *AuthService) Register(email, password string) error {
	_, err := a.userRepo.GetByEmail(email)
	if err == nil {
		return errors.New("user already exist")
	}
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:         email,
		Hash_Password: hashed,
	}
	return a.userRepo.Create(user)
}

func (a *AuthService) Login(email, password string) (map[string]string, error) {
	user, err := a.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if err := utils.CheckHashPassword(password, user.Hash_Password); err != nil {
		return nil, errors.New("invalid credentials")
	}
	access, err := a.jwtService.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}
	refresh, err := a.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	}, nil
}
