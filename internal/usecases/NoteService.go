package usecases

import (
	"MyNote/internal/domain/models"
	"MyNote/internal/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type NoteService struct {
	repo repository.INoteRepository
}

func NewNoteService(r repository.INoteRepository) *NoteService {
	return &NoteService{repo: r}
}

func (s *NoteService) Create(note *models.Note) error {
	if note.UserID <= 0 {
		return errors.New("incorrect User ID")
	}
	if note.Title == "" {
		return errors.New("title cannot be empty")
	}
	if note.Content == "" {
		return errors.New("content cannot be empty")
	}
	note.CreatedAt = time.Now()
	return s.repo.Create(note)
}

func (s *NoteService) GetByID(id uint) (*models.Note, error) {
	if id <= 0 {
		return nil, errors.New("ID cannot be less or equal to 0")
	}
	note, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("note not found")
		}
		return nil, err
	}
	if note == nil {
		return nil, errors.New("note not found")
	}
	return note, nil
}

func (s *NoteService) GetAllByUserID(userID uint) ([]models.Note, error) {
	if userID <= 0 {
		return nil, errors.New("userID cannot be less or equal to 0")
	}
	return s.repo.GetAllByUserID(userID)
}

func (s *NoteService) Update(note *models.Note) error {
	if note.ID <= 0 {
		return errors.New("ID cannot be less or equal to 0")
	}
	if note.Title == "" {
		return errors.New("title cannot be empty")
	}
	if note.Content == "" {
		return errors.New("content cannot be empty")
	}
	note.UpdatedAt = time.Now()
	return s.repo.Update(note)
}

func (s *NoteService) Delete(id uint) error {
	if id <= 0 {
		return errors.New("ID cannot be less or equal to 0")
	}
	return s.repo.Delete(id)
}
