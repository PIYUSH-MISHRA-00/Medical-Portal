package handlers

import (
	"encoding/json"
	"net/http"

	"medical-portal/internal/models"
	"medical-portal/internal/repository"
)

type DoctorHandler struct {
	patientRepo *repository.PatientRepository
}

func NewDoctorHandler(patientRepo *repository.PatientRepository) *DoctorHandler {
	return &DoctorHandler{patientRepo: patientRepo}
}

func (h *DoctorHandler) GetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := h.patientRepo.GetAllPatients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(patients)
}

func (h *DoctorHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
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