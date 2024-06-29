package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"recruitment_system/models"
	"recruitment_system/utils"
)

func UploadResume(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)

	if user.UserType != "Applicant" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	file, header, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	if header.Header.Get("Content-Type") != "application/pdf" && header.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	part, err := writer.CreateFormFile("resume", header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(part, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Close()

	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("apikey", "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to process resume", http.StatusInternalServerError)
		return
	}

	var result models.ResumeData
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile := models.Profile{
		Applicant:  user.Email,
		Resume:     header.Filename,
		Skills:     result.Skills,
		Education:  result.Education,
		Experience: result.Experience,
	}

	if err := profile.SaveProfile(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
