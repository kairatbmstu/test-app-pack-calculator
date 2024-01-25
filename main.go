package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var packSizes = []int{23, 51, 26}

type PageVariables struct {
	PackSizes []int
	Packs     []Pack
}

type Pack struct {
	Size int
	Num  int
}

type PackService struct {
}

func (p PackService) SubmitPackSettings(packSizeSettings []int) {
	packSizes = packSizeSettings
}

func (p PackService) CalculatePacks(TotalNumberOfPackes int, packSizeSettings []int) ([]Pack, error) {
	return nil, nil
}

type AppController struct {
}

func (a AppController) GetIndexPage(w http.ResponseWriter, r *http.Request) {
	// Define data to be passed to the template
	data := PageVariables{
		PackSizes: packSizes,
		Packs:     []Pack{},
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("web/Index.html")
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

func (a AppController) SubmitPackSettings(w http.ResponseWriter, r *http.Request) {
}

func (a AppController) CalculatePacks(w http.ResponseWriter, r *http.Request) {
}

func main() {
	controller := AppController{}
	http.HandleFunc("/", controller.GetIndexPage)
	http.HandleFunc("/hello", controller.GetIndexPage)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
