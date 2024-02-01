package service

import (
	"test-app-repartners/model"
	"testing"
)

func TestSubmitPackSettings(t *testing.T) {
	packService := &PackService{}
	packSizes := []int{250, 500, 1000, 500} // includes duplicates
	expectedResult := []int{250, 500, 1000}

	result := packService.SubmitPackSettings(packSizes)

	if len(result) != len(expectedResult) {
		t.Errorf("Expected length of %v, but got %v", len(expectedResult), len(result))
	}

	for i, v := range result {
		if v != expectedResult[i] {
			t.Errorf("Expected %v at index %v, but got %v", expectedResult[i], i, v)
		}
	}
}

func TestCalculatePacks(t *testing.T) {
	packService := &PackService{
		PackSizes: []int{250, 500, 1000},
	}

	totalNumberOfPacks := 2000
	expectedResult := []model.Pack{
		{Size: 250, Num: 4},
		{Size: 500, Num: 2},
		{Size: 1000, Num: 1},
	}

	result, err := packService.CalculatePacks(totalNumberOfPacks)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != len(expectedResult) {
		t.Errorf("Expected length of %v, but got %v", len(expectedResult), len(result))
	}

	for i, v := range result {
		if v != expectedResult[i] {
			t.Errorf("Expected %v at index %v, but got %v", expectedResult[i], i, v)
		}
	}
}
