package handlers

import (
	"encoding/json"
	"net/http"

	"medical-portal/internal/models"
	"medical-portal/internal/repository"
)

type ReceptionistHandler struct {
	patientRepo *repository.PatientRepository
}

func NewReceptionistHandler(patientRepo *repository.PatientRepository) *ReceptionistHandler {
	return &ReceptionistHandler{patientRepo: patientRepo}
}

func (h *ReceptionistHandler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.patientRepo.CreatePatient(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ReceptionistHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.patientRepo.UpdatePatient(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ReceptionistHandler) DeletePatient(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if err := h.patientRepo.DeletePatient(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ReceptionistHandler) GetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := h.patientRepo.GetAllPatients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(patients)
}