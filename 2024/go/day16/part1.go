package day16

import (
	"math"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day16/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	// Get grid from input
	grid, err := utils.ConstructGrid(inputFile)
	if err != nil {
		return 0, err
	}

	// Get source and destination
	source := [2]int{0, 0}
	dest := [2]int{0, 0}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				source = [2]int{i, j}
			} else if cell == 'E' {
				dest = [2]int{i, j}
			}
		}
	}
	// Set source and dest to '.' in grid
	grid[source[0]][source[1]] = '.'
	grid[dest[0]][dest[1]] = '.'

	// Get distance
	dist := calculateMinPath(grid, source, dest)

	return int64(dist), nil
}

func calculateMinPath(grid [][]rune, source [2]int, dest [2]int) int {
	// Calculate minimum distance using Dijkstra's Algorithm
	distances := map[[2][2]int]int{}
	// Start facing east from source (0, 1)
	distances[[2][2]int{source, {0, 1}}] = 0
	visited := hashset.New()
	heap := priorityqueue.NewWith(func(a, b interface{}) int { return distances[a.([2][2]int)] - distances[b.([2][2]int)] })
	heap.Enqueue([2][2]int{source, {0, 1}})

	for !heap.Empty() {
		// Get node with shortest distance
		_node, _ := heap.Dequeue()
		node := _node.([2][2]int)
		visited.Add(node)

		// Check neighbors
		for _, adj := range utils.Directions {
			di := node[0][0] + adj[0]
			dj := node[0][1] + adj[1]
			neighbor := [2]int{di, dj}
			if grid[di][dj] != '.' || visited.Contains([2][2]int{neighbor, adj}) {
				// Not a movable spot or already visited
				continue
			}
			// Check if distance is already calculated
			dist, ok := distances[[2][2]int{neighbor, adj}]
			if !ok {
				dist = -1
			}
			// Calculate distance
			newDist := 1
			// Check if a turn is needed
			if node[1] != adj {
				// Check how many turns are needed
				if adj[0]+node[1][0] == 0 && adj[1]+node[1][1] == 0 {
					newDist += 2000
				} else {
					newDist += 1000
				}
			}
			newDist += distances[node]
			// If distance is not calculated or new distance is shorter, update distance
			if dist == -1 || newDist < dist {
				distances[[2][2]int{neighbor, adj}] = newDist
			}
			// Enqueue neighbor
			heap.Enqueue([2][2]int{neighbor, adj})
		}
	}

	// Check distance to destination from all directions
	dNorth := distances[[2][2]int{dest, {0, 1}}]
	dEast := distances[[2][2]int{dest, {1, 0}}]
	dSouth := distances[[2][2]int{dest, {0, -1}}]
	dWest := distances[[2][2]int{dest, {-1, 0}}]
	dist := math.MaxInt
	// Get mininum distance to destination
	if dNorth != 0 && dNorth < dist {
		dist = dNorth
	}
	if dEast != 0 && dEast < dist {
		dist = dEast
	}
	if dSouth != 0 && dSouth < dist {
		dist = dSouth
	}
	if dWest != 0 && dWest < dist {
		dist = dWest
	}

	return dist
}
