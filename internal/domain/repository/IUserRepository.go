package repository

import "MyNote/internal/domain/models"

type IUserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}
