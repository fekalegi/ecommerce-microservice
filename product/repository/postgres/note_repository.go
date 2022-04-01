package postgres

import (
	"ecommerce-microservice/order/domain"
	"ecommerce-microservice/order/repository"
	"errors"
	"gorm.io/gorm"
)

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) repository.NoteRepository {
	return noteRepository{
		db,
	}
}

func (n noteRepository) CreateNote(note domain.Note) (*domain.Note, error) {
	err := n.db.Create(&note).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (n noteRepository) FetchNoteByID(id int) (*domain.Note, error) {
	var note domain.Note
	err := n.db.Preload("User").First(&note, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &note, nil
}

func (n noteRepository) FetchAllNote() ([]domain.Note, error) {
	var notes []domain.Note
	err := n.db.Preload("User").Find(&notes).Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (n noteRepository) UpdateNote(note domain.Note) error {
	err := n.db.Where("id = ?", note.ID).Updates(&note).Error
	if err != nil {
		return err
	}
	return nil
}

func (n noteRepository) DeleteNote(id int) error {
	err := n.db.Delete(&domain.Note{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
