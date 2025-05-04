package usecases

import (
	"MyNote/internal/domain/models"
	"MyNote/internal/domain/repository"
	"errors"
	"time"
)

type NoteService struct {
	repo repository.INoteRepository
}

func NewNoteRepository(r repository.INoteRepository) *NoteService {
	return &NoteService{repo: r}
}

func (s *NoteService) Create(note *models.Note) error {
	if note.UserID <= 0 {
		return errors.New("Incorrect User ID")
	}
	if note.Title == "" {
		return errors.New("Title cannot be empty")
	}
	note.CreatedAt = time.Now()
	return s.repo.Create(note)
}

func (s *NoteService) GetByID(id uint) (*models.Note, error) {
	if id <= 0 {
		return nil, errors.New("Id cannot be less or equal to 0")
	}
	note, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if note == nil {
		return nil, errors.New("Note not found")
	}
	return note, nil
}
