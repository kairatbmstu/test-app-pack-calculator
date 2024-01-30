package main

import (
	"fmt"
	"net/http"
	"test-app-repartners/controller"
)

func main() {

	http.Handle("/static", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/", controller.GetIndexPage)

	http.HandleFunc("/submitPackSettings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controller.GetPackageSettings(w, r)
			return
		}

		if r.Method == http.MethodPost {
			controller.PostPackageSettings(w, r)
			return
		}

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/calculatePacks", controller.CalculatePacks)

	// Start the server
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
