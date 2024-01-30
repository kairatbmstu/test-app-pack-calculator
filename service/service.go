// Package service provides functionality related to pack settings and calculations.
package service

import (
	"fmt"
	"sort"
	"test-app-repartners/model"
)

// PackServiceBean is an instance of PackService with default PackSizes.
var PackServiceBean = PackService{
	PackSizes: []int{250, 500, 1000},
}

// Packs is a type representing a collection of model.Pack instances.
type Packs []model.Pack

// Len is the number of elements in the collection.
func (p Packs) Len() int {
	return len(p)
}

// Less reports whether the element with index i must sort before the element with index j.
func (p Packs) Less(i, j int) bool {
	return p[i].Size < p[j].Size
}

// Swap swaps the elements with indexes i and j.
func (p Packs) Swap(i, j int) {
	tmp := p[j]
	p[j] = p[i]
	p[i] = tmp
}

// PackService is a service for handling pack-related operations.
type PackService struct {
	PackSizes []int
}

// SubmitPackSettings updates the PackSizes in PackService with the given packSizeSettings.
func (p *PackService) SubmitPackSettings(packSizeSettings []int) []int {
	p.PackSizes = packSizeSettings
	return p.PackSizes
}

// CalculatePacks calculates the optimal distribution of packs for a given TotalNumberOfPacks.
// It uses a dynamic programming approach to find the solution.
func (p *PackService) CalculatePacks(TotalNumberOfPacks int) ([]model.Pack, error) {
	sort.Sort(sort.IntSlice(p.PackSizes))
	width := TotalNumberOfPacks / p.PackSizes[0]
	height := len(p.PackSizes)

	dp := make([][]int, height)
	for i := range dp {
		dp[i] = make([]int, width+2)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width+2; j++ {
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

// Algorithm is a struct representing the state and parameters of the dynamic programming algorithm.
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

// Start initiates the dfs algorithm and returns true if a valid solution is found.
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

// Dfs is a recursive function used  to explore possible solutions.
func (a *Algorithm) Dfs(i int, j int, targetSum int) bool {
	a.Stack = append(a.Stack, model.Pack{Size: a.PackSizes[i], Num: j, DpCoords: model.DpCoords{I: i, J: j}}) //push item to stack
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

// completeMinStack updates the MinStack to represent the optimal solution after the algorithm completes.
func (a *Algorithm) completeMinStack() {
	sort.Sort(a.MinStack)

	rightVal := a.Dp[a.MinStack[0].DpCoords.I][a.MinStack[0].DpCoords.J+1]
	bottomVal := a.Dp[a.MinStack[0].DpCoords.I+1][a.MinStack[0].DpCoords.J]

	if rightVal == bottomVal {
		a.MinStack[0].DpCoords.I += 1
	} else if rightVal < bottomVal {
		a.MinStack[0].DpCoords.J += 1
	} else if rightVal > bottomVal {
		a.MinStack[0].DpCoords.I += 1
	}

	a.MinStack[0].Num = a.MinStack[0].DpCoords.J
	a.MinStack[0].Size = a.PackSizes[a.MinStack[0].DpCoords.I]

	sort.Reverse(a.MinStack)
}
