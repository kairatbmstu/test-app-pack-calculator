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

	sort.Sort(sort.Reverse(sort.IntSlice(p.PackSizes)))

	for _, size := range p.PackSizes {
		numPacks := TotalNumberOfPacks / size
		if numPacks > 0 {
			packs = append(packs, model.Pack{Size: size, Num: numPacks})
			TotalNumberOfPacks -= numPacks * size
		}
	}

	sort.Sort(sort.IntSlice(p.PackSizes))

	if TotalNumberOfPacks > 0 {
		for _, size := range p.PackSizes {
			if size > TotalNumberOfPacks {
				foundPackSize := false
				for _, pack := range packs {
					if pack.Size == size {
						pack.Num += 1
						foundPackSize = true
					}
				}

				if !foundPackSize {
					packs = append(packs, model.Pack{Size: size, Num: 1})
					break
				}
			}
		}
	}

	return packs, nil
}
