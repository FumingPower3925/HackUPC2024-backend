package airport

import (
	"errors"
	"math/rand"
)

func Void() {}

var AirportMap [][]int
var copyAirportMap [][]int
var bays []Point
var currTarget Point

func CreateEmptyMatrix(dim int) [][]int {
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

func init() {
	AirportMap, bays = GenerateRandomMapWithWallsAndBays(10, 10, 10, 5)

	copyAirportMap = AirportMap
}

func GetTarget() Point {
	currTarget = bays[rand.Intn(len(bays))]
	return currTarget
}

func NextStep(pos, target Point) (Point, error) {
	path := AStar(copyAirportMap, pos, target)

	if path == nil {
		return Point{}, errors.New("no path found")
	}

	copyAirportMap[pos.X][pos.Y] = Wall // Reset previous pos

	// Update pos to the next point in the path
	next := path[len(path)-2] // Index 0 is the current pos

	// Update pos
	pos.X = next.P.X
	pos.Y = next.P.Y

	copyAirportMap[pos.X][pos.Y] = Person // Set new pos

	PrintResults(copyAirportMap, pos, target)

	return pos, nil
}
