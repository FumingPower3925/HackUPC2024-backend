package airport

import "math/rand"

func TestLogic() bool {
	originalAirportMap, bays := GenerateRandomMapWithWallsAndBays(10, 10, 10, 5)

	modifiedMap := originalAirportMap
	position := Point{X: 0, Y: 0}
	target := bays[rand.Intn(len(bays))]

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
