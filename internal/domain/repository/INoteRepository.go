package repository

import "MyNote/internal/domain/models"

type INoteRepository interface {
	Create(note *models.Note) error
	GetByID(id uint) (*models.Note, error)
	Update(note *models.Note) error
	Delete(id uint) error
}
