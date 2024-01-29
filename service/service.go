package service

import (
	"fmt"
	"sort"
	"test-app-repartners/model"
)

var PackServiceBean = PackService{
	PackSizes: []int{23, 31, 53},
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

	var stack = []model.Pack{}

	var algo = Algorithm{
		Height:             height,
		Width:              width,
		PackSizes:          p.PackSizes,
		TotalNumberOfPacks: TotalNumberOfPacks,
		MinSum:             TotalNumberOfPacks,
		Stack:              stack,
		Dp:                 dp,
	}

	var result = algo.Start()

	fmt.Println("result : ", result)
	fmt.Println("stack : ", algo.Stack)
	fmt.Println("MinSum : ", algo.MinSum)
	fmt.Println("MinStack : ", algo.MinStack)

	if result {
		return algo.MinStack, nil
	}

	return packs, nil
}

type Algorithm struct {
	Height             int
	Width              int
	PackSizes          []int
	TotalNumberOfPacks int
	MinSum             int
	Stack              []model.Pack
	MinStack           []model.Pack
	Dp                 [][]int
}

func (a *Algorithm) Start() bool {

	for l := a.Width; l >= 0; l-- {
		var found = a.Dfs(a.Height-1, l, a.TotalNumberOfPacks-a.Dp[a.Height-1][l])
		fmt.Println("found: ", found)
		if found {
			return true
		}
	}

	return false
}

func (a *Algorithm) Dfs(i int, j int, targetSum int) bool {
	a.Stack = append(a.Stack, model.Pack{Size: a.PackSizes[i], Num: j}) //push item to stack
	fmt.Println("stack: ", a.Stack)
	fmt.Println("targetSum: ", targetSum)

	if targetSum < 0 {
		if len(a.Stack) > 0 {
			a.Stack = a.Stack[:len(a.Stack)-1] // pop item from stack
		}
		return false
	}

	if i == 0 {
		if targetSum <= a.MinSum {
			a.MinSum = targetSum
			a.MinStack = a.Stack[:]
		}
		if targetSum == 0 {
			return true
		}
		if len(a.Stack) > 0 {
			a.Stack = a.Stack[:len(a.Stack)-1] // pop item from stack
		}
		return false
	}

	for l := a.Width; l >= 0; l-- {
		var found = a.Dfs(i-1, l, targetSum-a.Dp[i-1][l])
		//fmt.Println("found: ", found)
		if found {
			return true
		}
	}

	if len(a.Stack) > 0 {
		a.Stack = a.Stack[:len(a.Stack)-1] // pop item from stack
	}

	return false
}
