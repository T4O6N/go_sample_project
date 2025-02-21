package handlers

import (
	"encoding/json"
	"net/http"
	"sample-project/internal/entities"
	"sample-project/internal/usecases/subject"
)

type SubjectHandler struct {
	SubjectUseCase *subject.SubjectUseCase
}

func NewSubjectHandler(subjectUseCase *subject.SubjectUseCase) *SubjectHandler {
	return &SubjectHandler{SubjectUseCase: subjectUseCase}
}

// CreateUser godoc
// @Summary Create a new subject
// @Description Create a new subject with the input payload
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param subject body models.Subject true "Create subject"
// @Success 201 {object} models.Subject
// @Router /subjects [post]
func (h *SubjectHandler) CreateSubject(w http.ResponseWriter, r *http.Request) {
	var subject entities.Subject
	json.NewDecoder(r.Body).Decode(&subject)
	err := h.SubjectUseCase.CreateSubject(&subject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subject)
}
