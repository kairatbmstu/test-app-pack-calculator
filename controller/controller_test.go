package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPackageSettings(t *testing.T) {
	req, err := http.NewRequest("GET", "/get-package-settings", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPackageSettings)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetPackageSettings returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// You can add more assertions here if needed.
}

func TestPostPackageSettings(t *testing.T) {
	// Prepare a sample request body
	requestBody := []int{1, 2, 3}

	// Convert the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/post-package-settings", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostPackageSettings)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("PostPackageSettings returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// You can add more assertions here if needed.
}

func TestCalculatePacks(t *testing.T) {
	// Prepare a sample request body
	requestBody := []byte(`{"totalNumberOfPacks": 10}`)

	req, err := http.NewRequest("POST", "/calculate-packs", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculatePacks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CalculatePacks returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// You can add more assertions here if needed.
}
