package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Mohamadreza-shad/notepad/logger"
	"gorm.io/gorm"
)

var (
	ErrNotepadAlreadyExists = errors.New("notepad already exists")
	ErrSomethingWentWrong   = errors.New("something went wrong")
)

type Repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

type Notepad struct {
	Title     string
	Content   string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

type CreateNotepadParams struct {
	Title   string
	Content string
}

func (r *Repository) CreateNotepad(ctx context.Context, params CreateNotepadParams) error {
	err := r.db.Where("title = ?", params.Title).Find(&Notepad{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrSomethingWentWrong
	}
	if err == nil {
		return ErrNotepadAlreadyExists
	}
	err = r.db.Create(
		Notepad{
			Title:     params.Title,
			Content:   params.Content,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: 0,
			DeletedAt: 0,
		}).Error
	return err
}

func New(db *gorm.DB, l *logger.Logger) *Repository {
	return &Repository{db, l}
}
