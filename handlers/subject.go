package handlers

import (
	"encoding/json"
	"net/http"
	"sample-project/models"
)

// CreateUser godoc
// @Summary Create a new subject
// @Description Create a new subject with the input payload
// @Tags subjects
// @Accept  json
// @Produce  json
// @Param subject body models.Subject true "Create subject"
// @Success 201 {object} models.Subject
// @Router /subjects [post]
func CreateSubject(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	json.NewDecoder(r.Body).Decode(&subject)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subject)
}
