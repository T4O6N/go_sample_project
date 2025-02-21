package repositories

import (
	"sample-project/internal/entities"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	Create(subject *entities.Subject) error
}

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) Create(subject *entities.Subject) error {
	return r.db.Create(subject).Error
}
