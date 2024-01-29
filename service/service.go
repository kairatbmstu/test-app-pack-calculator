package service

import (
	"fmt"
	"sort"
	"test-app-repartners/model"
)

var PackServiceBean = PackService{
	PackSizes: []int{250, 500, 1000},
}

type Packs []model.Pack

// Len is the number of elements in the collection.
func (p Packs) Len() int {
	return len(p)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (p Packs) Less(i, j int) bool {
	return p[i].Size < p[j].Size
}

// Swap swaps the elements with indexes i and j.
func (p Packs) Swap(i, j int) {
	tmp := p[j]
	p[j] = p[i]
	p[i] = tmp
}

type PackService struct {
	PackSizes []int
}

func (p *PackService) SubmitPackSettings(packSizeSettings []int) []int {
	p.PackSizes = packSizeSettings
	return p.PackSizes
}

func (p *PackService) CalculatePacks(TotalNumberOfPacks int) ([]model.Pack, error) {
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
		return algo.Stack, nil
	}

	algo.completeMinStack()

	return algo.MinStack, nil
}

type Algorithm struct {
	Height             int
	Width              int
	PackSizes          []int
	TotalNumberOfPacks int
	MinSum             int
	Stack              Packs
	MinStack           Packs
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
			a.MinStack = make([]model.Pack, len(a.Stack))
			copy(a.MinStack, a.Stack)
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

func (a *Algorithm) completeMinStack() {
	total := 0
	for _, item := range a.MinStack {
		total += item.Num * item.Size
	}

	diff := a.TotalNumberOfPacks - total

	sort.Sort(a.MinStack)

	for i, _ := range a.MinStack {
		if a.MinStack[i].Size >= diff {
			a.MinStack[i].Num += 1
			break
		}
	}

	sort.Reverse(a.MinStack)
}

func (a *Algorithm) shrinkPacks() {
	for _, pack := range a.MinStack {
		itemsCount := pack.Num * pack.Size

	}
}
