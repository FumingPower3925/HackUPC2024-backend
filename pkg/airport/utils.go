package airport

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
)

// Int mapping
const (
	Empty    int = -1
	Walkable int = 0
	Wall     int = 1
	Person   int = 2
	Bay      int = 3
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Point struct {
	X, Y int
}

type Node struct {
	P          Point
	f, g, h    float64
	parent     *Node
	entityType int
}

type PriorityQueue []*Node

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

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func distance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func (p Point) neighbors(matrix [][]int) []Point {
	neighbors := make([]Point, 0)
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		nx, ny := p.X+dir[0], p.Y+dir[1]
		if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) {
			if matrix[nx][ny] != Wall {
				neighbors = append(neighbors, Point{nx, ny})
			}
		}
	}
	return neighbors
}

func AStar(matrix [][]int, start, goal Point) []*Node {
	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)

	startNode := &Node{P: start, f: 0, g: 0, h: 0}
	heap.Push(&openSet, startNode)

	closedSet := make(map[Point]bool)
	cameFrom := make(map[Point]*Node)

	for len(openSet) > 0 {
		current := heap.Pop(&openSet).(*Node)

		if current.P == goal {
			path := make([]*Node, 0)
			for current != nil {
				path = append(path, current)
				current = current.parent
			}
			return path
		}

		closedSet[current.P] = true

		for _, neighbor := range current.P.neighbors(matrix) {
			if closedSet[neighbor] {
				continue
			}

			g := current.g + 1 // Assuming each step has a cost of 1

			newPath := false
			if _, ok := cameFrom[neighbor]; !ok {
				newPath = true
			} else if g < cameFrom[neighbor].g {
				newPath = true
			}

			if newPath {
				neighborNode := &Node{
					P:          neighbor,
					g:          g,
					h:          distance(neighbor, goal),
					parent:     current,
					entityType: matrix[neighbor.X][neighbor.Y],
				}
				heap.Push(&openSet, neighborNode)
				cameFrom[neighbor] = neighborNode
			}
		}
	}

	return nil
}

func Dijkstra(matrix [][]int, start, goal Point) []*Node {
	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)

	startNode := &Node{P: start, f: 0, g: 0, h: 0} // f, g, h are not used in Dijkstra's
	heap.Push(&openSet, startNode)

	closedSet := make(map[Point]bool)
	cameFrom := make(map[Point]*Node)
	costs := make(map[Point]float64) // Keep track of the cost to reach each node

	for len(openSet) > 0 {
		current := heap.Pop(&openSet).(*Node)

		if current.P == goal {
			path := make([]*Node, 0)
			for current != nil {
				path = append(path, current)
				current = current.parent
			}
			return path
		}

		closedSet[current.P] = true

		for _, neighbor := range current.P.neighbors(matrix) {
			if closedSet[neighbor] {
				continue
			}

			g := current.g + 1 // Assuming each step has a cost of 1

			newCost := g
			if _, ok := costs[neighbor]; !ok || newCost < costs[neighbor] {
				costs[neighbor] = newCost
				neighborNode := &Node{
					P:      neighbor,
					g:      newCost,
					parent: current,
				}
				heap.Push(&openSet, neighborNode)
				cameFrom[neighbor] = neighborNode
			}
		}
	}

	return nil
}

func GenerateRandomMapWithWallsAndBays(rows, cols, numWalls, numBays int) ([][]int, []Point) {
	// Generate an empty map
	airportMap := make([][]int, rows)
	for i := range airportMap {
		airportMap[i] = make([]int, cols)
		for j := range airportMap[i] {
			airportMap[i][j] = Walkable
		}
	}

	// Generate walls
	for i := 0; i < numWalls; i++ {
		x := rand.Intn(rows)
		y := rand.Intn(cols)
		airportMap[x][y] = Wall
	}

	// Generate bays and store their locations
	bays := make([]Point, numBays)
	for i := 0; i < numBays; i++ {
		x := rand.Intn(rows)
		y := rand.Intn(cols)
		airportMap[x][y] = Bay
		bays[i] = Point{X: x, Y: y}
	}

	return airportMap, bays
}

func (d Point) String() string {
	switch d {
	case Point{0, 1}:
		return "FORWARD"
	case Point{-1, 0}:
		return "TURN LEFT"
	case Point{1, 0}:
		return "TURN RIGHT"
	case Point{0, -1}:
		return "TURN AROUND"
	default:
		return "unknown"
	}
}
