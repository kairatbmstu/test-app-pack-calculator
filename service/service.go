package service

import (
	"fmt"
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
		dp[i] = make([]int, width+1)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width+1; j++ {
			dp[i][j] = p.PackSizes[i] * j
		}
	}

	var stack = []int{}

	dfs(height-1, width, TotalNumberOfPacks, dp, stack)
	fmt.Println(stack)

	return packs, nil
}

func dfs(i int, j int, targetSum int, dp [][]int, stack []int) bool {
	fmt.Println("i : {}  j : {} ", i, j)
	stack = append(stack, dp[i][j]) //push item to stack
	if i == 0 {
		diff := targetSum - dp[i][j]
		fmt.Println("diff : %s", diff)
		if diff == 0 {
			return true
		}

		return false
	}

	for l := j; l >= 0; l-- {
		if dp[i][l] <= targetSum {
			var found = dfs(i-1, l, targetSum-dp[i][j], dp, stack)
			fmt.Println("found: ", found)
			if found {
				return true
			} else {
				stack = stack[:i-1] // pop item from stack
			}
		}
	}

	return false
}
