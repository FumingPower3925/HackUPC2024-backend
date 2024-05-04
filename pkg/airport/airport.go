package airport

import (
	"fmt"
)

func createEmptyMatrix(dim int) [][]int {
	matrix := make([][]int, dim)
	for i := range matrix {
		matrix[i] = make([]int, dim)
	}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = Empty
		}
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, cell := range row {
			switch cell {
			case Empty:
				fmt.Print("-1")
			case Walkable:
				fmt.Print("W ")
			case Wall:
				fmt.Print("# ")
			case Person:
				fmt.Print("P ")
			case Bay:
				fmt.Print("B ")
			}
		}
		fmt.Println()
	}
}

func Test() {
	// var dim int = 10
	// matrix := generateRandomMap(dim, dim, 33)
	matrix4x4 := [][]int{
		{Person, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, Bay},
	}

	printResults(matrix4x4)
}

func printResults(matrix [][]int) {
	start := Point{0, 0}
	goal := Point{len(matrix) - 1, len(matrix[0]) - 1}

	fmt.Printf("Matrix: %dx%d\n", len(matrix), len(matrix[0]))
	printMatrix(matrix)

	path := AStar(matrix, start, goal)
	if path == nil {
		fmt.Println("No path found")
	} else {
		for i := len(path) - 1; i >= 0; i-- {
			fmt.Printf("(%d,%d) -> ", path[i].point.x, path[i].point.y)
		}
		fmt.Println()
	}
}
