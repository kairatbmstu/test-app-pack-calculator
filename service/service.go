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

	var algo = Algorithm{
		Height:             height,
		Width:              width,
		TotalNumberOfPacks: TotalNumberOfPacks,
		Stack:              stack,
		Dp:                 dp,
	}

	var result = algo.Start(TotalNumberOfPacks)

	fmt.Println("result : ", result)

	return packs, nil
}

type Algorithm struct {
	Height             int
	Width              int
	TotalNumberOfPacks int
	Stack              []int
	Dp                 [][]int
}

func (a Algorithm) Start(targetSum int) bool {

	for l := a.Width; l >= 0; l-- {
		if a.Dp[a.Height-1][l] <= a.TotalNumberOfPacks {
			var found = a.Dfs(a.Height-1, l, targetSum-a.Dp[a.Height-1][l])
			fmt.Println("found: ", found)
			if found {
				return true
			}
		}
	}

	return false
}

func (a Algorithm) Dfs(i int, j int, targetSum int) bool {
	//fmt.Printf("i : %d  j : %d \n", i, j)
	a.Stack = append(a.Stack, a.Dp[i][j]) //push item to stack
	fmt.Println("stack: ", a.Stack)
	if i == 0 {
		diff := targetSum - a.Dp[i][j]
		//fmt.Println("diff : %s", diff)
		if diff == 0 {
			return true
		}

		return false
	}

	for l := a.Width; l >= 0; l-- {
		if a.Dp[i][l] <= a.TotalNumberOfPacks {
			var found = a.Dfs(i-1, l, targetSum-a.Dp[i][l])
			//fmt.Println("found: ", found)
			if found {
				return true
			} else {
				if len(a.Stack) > 0 {
					a.Stack = a.Stack[:len(a.Stack)] // pop item from stack
				}
			}
		}
	}

	return false
}
