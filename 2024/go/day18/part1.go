package day18

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

// Sample input
// const GRID_SIZE = 7
// const BYTE_LIMIT = 12

const GRID_SIZE = 71
const BYTE_LIMIT = 1024

func init() {
	registry.Registry["day18/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	// Make grid
	grid := [GRID_SIZE][GRID_SIZE]bool{}

	// Read input file and create scanner
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)
	cntr := 0

	// Add obstacles to grid upto BYTE_LIMIT
	for scanner.Scan() {
		if cntr == BYTE_LIMIT {
			break
		}
		line := scanner.Text()
		numsStr := strings.Split(line, ",")
		x, _ := strconv.Atoi(numsStr[0])
		y, _ := strconv.Atoi(numsStr[1])
		grid[y][x] = true
		cntr++
	}

	// Calculate shortest path
	dist := calculateDistance(grid)

	return int64(dist), nil
}

func calculateDistance(grid [GRID_SIZE][GRID_SIZE]bool) int {
	// Dijkstra's algorithm
	distances := map[[2]int]int{}
	visited := hashset.New()
	queue := priorityqueue.NewWith(func(a, b interface{}) int { return distances[a.([2]int)] - distances[b.([2]int)] })
	queue.Enqueue([2]int{0, 0})
	distances[[2]int{0, 0}] = 0

	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.([2]int)
		visited.Add(node)

		for _, dir := range utils.Directions {
			di := node[0] + dir[0]
			dj := node[1] + dir[1]
			adj := [2]int{di, dj}
			if di < 0 || dj < 0 || di == GRID_SIZE || dj == GRID_SIZE {
				continue
			}
			if visited.Contains(adj) || grid[di][dj] {
				continue
			}
			oldDist, exists := distances[adj]
			newDist := distances[node] + 1
			if !exists || oldDist > newDist {
				distances[adj] = newDist
				queue.Enqueue([2]int{di, dj})
			}
		}
	}

	return distances[[2]int{GRID_SIZE - 1, GRID_SIZE - 1}]
}
