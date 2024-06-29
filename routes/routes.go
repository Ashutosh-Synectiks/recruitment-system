package routes

import (
	"recruitment_system/controllers"
	"recruitment_system/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)

	api.HandleFunc("/uploadResume", controllers.UploadResume).Methods("POST")
	api.HandleFunc("/admin/job", controllers.CreateJob).Methods("POST")
	api.HandleFunc("/admin/job/{job_id}", controllers.GetJob).Methods("GET")
	api.HandleFunc("/admin/applicants", controllers.GetApplicants).Methods("GET")
	api.HandleFunc("/admin/applicant/{applicant_id}", controllers.GetApplicant).Methods("GET")
	api.HandleFunc("/jobs", controllers.GetJobs).Methods("GET")
	api.HandleFunc("/jobs/apply", controllers.ApplyToJob).Methods("GET")

	return r
}
