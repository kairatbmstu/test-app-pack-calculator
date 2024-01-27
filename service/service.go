package service

import (
	"sort"
	"test-app-repartners/model"
)

var PackServiceBean = PackService{
	PackSizes: []int{23, 51, 26},
}

type PackService struct {
	PackSizes []int
}

func (p *PackService) SubmitPackSettings(packSizeSettings []int) []int {
	p.PackSizes = packSizeSettings
	return p.PackSizes
}

func (p *PackService) CalculatePacks(TotalNumberOfPacks int) ([]model.Pack, error) {
	var packs []model.Pack
	sort.Sort(sort.IntSlice(p.PackSizes))
	width := TotalNumberOfPacks / p.PackSizes[0]
	height := len(p.PackSizes)

	dp := make([][]int, height)
	for i := range dp {
		dp[i] = make([]int, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			dp[i][j] = p.PackSizes[i] * i
		}
	}

	return packs, nil
}
