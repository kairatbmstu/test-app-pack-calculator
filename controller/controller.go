// Package controller defines HTTP handlers for managing package settings and calculations.
package controller

import (
	"encoding/json"
	"net/http"
	"test-app-repartners/model"
	"test-app-repartners/service"
	"text/template"
)

// GetPackageSettings handles HTTP GET requests to retrieve current package settings.
func GetPackageSettings(w http.ResponseWriter, r *http.Request) {
	// Convert the struct to JSON
	jsonData, err := json.Marshal(service.PackServiceBean.PackSizes)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}

// PostPackageSettings handles HTTP POST requests to update package settings.
func PostPackageSettings(w http.ResponseWriter, r *http.Request) {
	var packSizeSettings []int
	if err := json.NewDecoder(r.Body).Decode(&packSizeSettings); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	result := service.PackServiceBean.SubmitPackSettings(packSizeSettings)
	//fmt.Fprint(w, "Pack settings submitted successfully")

	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}

// CalculatePacks handles HTTP POST requests to calculate optimal pack distribution.
func CalculatePacks(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request struct {
		TotalNumberOfPacks int `json:"totalNumberOfPacks"`
	}

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	packs, err := service.PackServiceBean.CalculatePacks(request.TotalNumberOfPacks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packs)
}

type PageVariables struct {
	PackSizes []int
	Packs     []model.Pack
}

func GetIndexPage(w http.ResponseWriter, r *http.Request) {
	// Define data to be passed to the template
	data := PageVariables{
		PackSizes: service.PackServiceBean.PackSizes,
		Packs:     []model.Pack{},
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, passing in the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
