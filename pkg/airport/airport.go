package airport

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
)

const (
	X_min float64 = 10
	X_max float64 = 70
	Y_min float64 = -30
	Y_max float64 = 30
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
	AirportMap, bays = GenerateRandomMapWithWallsAndBays(60, 60, 10, 5)
	copyAirportMap = AirportMap
}

func GetTarget(uid string) (Point, error) {
	currTarget = bays[rand.Intn(len(bays))]
	fmt.Printf("%d %d\n", currTarget.X, currTarget.Y)
	return currTarget, nil
}

func extractLatitudeAndLongitude(jsonData string) (Location, error) {
	var loc Location
	err := json.Unmarshal([]byte(jsonData), &loc)
	if err != nil {
		return Location{}, &json.UnmarshalTypeError{}
	}

	return loc, nil
}

func Gps2D(jsonData string) (Point, error) {
	loc, _ := extractLatitudeAndLongitude(jsonData)

	if X_min > loc.Latitude || X_max < loc.Latitude || Y_min > loc.Longitude || Y_max < loc.Longitude {
		loc.Latitude = 11
		loc.Longitude = -29
	}

	return Point{int((loc.Latitude - 10) / 3.0 * 60.0), int((loc.Longitude + 30) / 3.0 * 60.0)}, nil
}

func NextStep(pos, target, lastPos Point) (string, error) {
	path := AStar(copyAirportMap, pos, target)

	if path == nil {
		return "", errors.New("no path found")
	}

	copyAirportMap[lastPos.X][lastPos.Y] = Wall // Reset previous pos

	next := path[len(path)-1]
	// Update pos to the next point in the path
	if len(path) != 1 {
		next = path[len(path)-2]
	}
	dirVect := Point{next.P.X - pos.X, next.P.Y - pos.Y}
	dirVectPrev := Point{pos.X - lastPos.X, pos.Y - lastPos.Y}

	// Update pos
	pos.X = next.P.X
	pos.Y = next.P.Y

	copyAirportMap[pos.X][pos.Y] = Person // Set new pos

	PrintResults(copyAirportMap, pos, target)
	dirString := dirVect.String()
	endPoint := Point{0, 0}
	if dirVect == dirVectPrev && dirVect != endPoint {
		dirString = "CONTINUE FORWARD"
	}
	if currTarget == pos {
		dirString = "ARRIVED"
	}
	return dirString, nil
}
