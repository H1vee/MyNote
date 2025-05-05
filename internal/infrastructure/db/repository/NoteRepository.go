package repository

import (
	"MyNote/internal/domain/models"
	"MyNote/internal/domain/repository"

	"gorm.io/gorm"
)

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) repository.INoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) Create(note *models.Note) error {
	return r.db.Create(note).Error
}

func (r *noteRepository) GetByID(id uint) (*models.Note, error) {
	var note models.Note
	// if err := r.db.Where("id =?",id).First(&note).Error; err!=nil{
	// 	return nil, err
	// }
	err := r.db.First(&note, id).Error
	return &note, err
}

func (r *noteRepository) GetAllByUserID(userID uint) ([]models.Note, error) {
	var notes []models.Note
	err := r.db.Where("user_id =?", userID).Find(&notes).Error
	return notes, err
}

func (r *noteRepository) Update(note *models.Note) error {
	return r.db.Save(note).Error
}

func (r *noteRepository) Delete(id uint) error {
	return r.db.Delete(&models.Note{}, id).Error
}
