package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"medical-portal/internal/auth"
	"medical-portal/internal/handlers"
	"medical-portal/internal/middleware"
	"medical-portal/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	patientRepo := repository.NewPatientRepository(db)

	authService := auth.NewAuthService(userRepo)
	doctorHandler := handlers.NewDoctorHandler(patientRepo)
	receptionistHandler := handlers.NewReceptionistHandler(patientRepo)

	http.HandleFunc("/login", authService.LoginHandler)
	http.HandleFunc("/doctor/patients", middleware.RBAC(doctorHandler.GetPatients, "doctor"))
	http.HandleFunc("/doctor/patient/update", middleware.RBAC(doctorHandler.UpdatePatient, "doctor"))
	http.HandleFunc("/receptionist/patient/create", middleware.RBAC(receptionistHandler.CreatePatient, "receptionist"))
	http.HandleFunc("/receptionist/patient/update", middleware.RBAC(receptionistHandler.UpdatePatient, "receptionist"))
	http.HandleFunc("/receptionist/patient/delete", middleware.RBAC(receptionistHandler.DeletePatient, "receptionist"))
	http.HandleFunc("/receptionist/patients", middleware.RBAC(receptionistHandler.GetPatients, "receptionist"))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}