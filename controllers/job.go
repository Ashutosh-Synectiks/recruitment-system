package controllers

import (
	"encoding/json"
	"net/http"
	"recruitment_system/models"
	"recruitment_system/utils"

	"github.com/gorilla/mux"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)

	if user.UserType != "Admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var job models.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	job.PostedBy = user.Email

	if err := job.CreateJob(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)

	if user.UserType != "Admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	jobID := vars["job_id"]

	job, err := models.GetJob(jobID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(job)
}

func GetApplicants(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)

	if user.UserType != "Admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	applicants, err := models.GetApplicants()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(applicants)
}

func GetApplicant(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)

	if user.UserType != "Admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	applicantID := vars["applicant_id"]

	applicant, err := models.GetApplicant(applicantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(applicant)
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := models.GetJobs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(jobs)
}

func ApplyToJob(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)

	if user.UserType != "Applicant" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jobID := r.URL.Query().Get("job_id")
	if jobID == "" {
		http.Error(w, "Job ID is required", http.StatusBadRequest)
		return
	}

	err := models.ApplyToJob(user.Email, jobID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
