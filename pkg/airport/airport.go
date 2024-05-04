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

func PrintMatrix(matrix [][]int) {
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

func TestLogic() bool {
	originalAirportMap := GenerateRandomMap(5, 5, 10)

	modifiedMap := originalAirportMap
	position := Point{X: 0, Y: 0}
	target := Point{X: len(modifiedMap) - 1, Y: len(modifiedMap[0]) - 1}

	modifiedMap[0][0] = Person
	modifiedMap[len(modifiedMap)-1][len(modifiedMap[0])-1] = Bay

	PrintMatrix(modifiedMap)
	PrintResults(modifiedMap, position, target)

	for position != target {
		path := AStar(modifiedMap, position, target)

		if path == nil {
			return false
		}

		modifiedMap[position.X][position.Y] = Wall // Reset previous position

		// Update position to the next point in the path
		next := path[len(path)-2] // Index 0 is the current position

		// Update position
		position.X = next.P.X
		position.Y = next.P.Y

		modifiedMap[position.X][position.Y] = Person // Set new position

		PrintResults(modifiedMap, position, target)
	}

	return true
}

func TestAlgos() {
	matrix1x5 := [][]int{
		{Person, 0, 0, 0, Bay},
	}

	PrintResults(matrix1x5, Point{0, 0}, Point{0, 4})

	matrix1x5[0][1] = 1

	PrintResults(matrix1x5, Point{0, 0}, Point{0, 4})

	matrix1x5[0][0] = 1
	matrix1x5[0][2] = Person
	PrintResults(matrix1x5, Point{0, 2}, Point{0, 4})

	matrix4x4 := [][]int{
		{Person, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, Bay},
	}

	PrintResults(matrix4x4, Point{0, 0}, Point{4, 4})

	matrix4x4[0][2] = 1

	PrintResults(matrix4x4, Point{0, 0}, Point{4, 4})

	matrix4x4[0][0] = 0
	matrix4x4[0][3] = Person
	PrintResults(matrix4x4, Point{0, 3}, Point{4, 4})
}

func PrintResults(matrix [][]int, start Point, goal Point) {
	fmt.Printf("Matrix: %dx%d\n", len(matrix), len(matrix[0]))
	PrintMatrix(matrix)

	path := AStar(matrix, start, goal)
	if path == nil {
		fmt.Println("No path found")
	} else {
		for i := len(path) - 1; i >= 0; i-- {
			fmt.Printf("(%d,%d) -> ", path[i].P.X, path[i].P.Y)
		}
	}
	fmt.Println()
	fmt.Println()
}
