package airport

import (
	"container/heap"
	"math"
	"math/rand"
	"time"
)

// Int mapping
const (
	Empty    int = -1
	Walkable int = 0
	Wall     int = 1
	Person   int = 2
	Bay      int = 3
)

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

// generateRandomMap creates a random map of given size with specified percentage of walls.
func GenerateRandomMap(rows, cols, wallPercentage int) [][]int {
	rand.Seed(time.Now().UnixNano())

	// Initialize the map
	randomMap := make([][]int, rows)
	for i := range randomMap {
		randomMap[i] = make([]int, cols)
	}

	// Calculate the number of walls to add based on the wallPercentage
	totalCells := rows * cols
	totalWalls := (totalCells * wallPercentage) / 100

	// Randomly add walls to the map
	for totalWalls > 0 {
		row := rand.Intn(rows)
		col := rand.Intn(cols)

		// Skip if already a wall
		if randomMap[row][col] == Wall {
			continue
		}

		randomMap[row][col] = Wall
		totalWalls--
	}

	return randomMap
}
