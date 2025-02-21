package subject

import (
	"sample-project/internal/entities"
	"sample-project/repositories"
)

type SubjectUseCase struct {
	SubjectRepo repositories.SubjectRepository
}

func NewSubjectUseCase(subjectRepo repositories.SubjectRepository) *SubjectUseCase {
	return &SubjectUseCase{SubjectRepo: subjectRepo}
}

func (uc *SubjectUseCase) CreateSubject(subject *entities.Subject) error {
	return uc.SubjectRepo.Create(subject)
}
