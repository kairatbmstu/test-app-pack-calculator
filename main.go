package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type PageVariables struct {
	PackSizes []int
	Packs     []Pack
}

type Pack struct {
	Size int `json:"size"`
	Num  int `json:"num"`
}

type PackService struct {
	packSizes []int
}

func (p *PackService) SubmitPackSettings(packSizeSettings []int) {
	p.packSizes = packSizeSettings
}

func (p *PackService) CalculatePacks(TotalNumberOfPacks int) ([]Pack, error) {
	var packs []Pack

	sort.Sort(sort.Reverse(sort.IntSlice(p.packSizes)))

	for _, size := range p.packSizes {
		numPacks := TotalNumberOfPacks / size
		if numPacks > 0 {
			packs = append(packs, Pack{Size: size, Num: numPacks})
			TotalNumberOfPacks -= numPacks * size
		}
	}

	if TotalNumberOfPacks != 0 {
		return nil, fmt.Errorf("unable to fulfill the order completely with available pack sizes")
	}

	return packs, nil
}

func main() {
	packService := &PackService{
		packSizes: []int{23, 51, 26},
	}

	http.HandleFunc("/submitPackSettings", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method == http.MethodGet {
		// 	json.NewEncoder().Encode(packService.packSizes)
		// }

		if r.Method == http.MethodPost {
			var packSizeSettings []int
			if err := json.NewDecoder(r.Body).Decode(&packSizeSettings); err != nil {
				http.Error(w, "Invalid JSON format", http.StatusBadRequest)
				return
			}

			packService.SubmitPackSettings(packSizeSettings)
			fmt.Fprint(w, "Pack settings submitted successfully")
		}

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/calculatePacks", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var request struct {
			TotalNumberOfPacks int `json:"totalNumberOfPacks"`
		}

		err := decoder.Decode(&request)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		packs, err := packService.CalculatePacks(request.TotalNumberOfPacks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(packs)
	})

	// Start the server
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
