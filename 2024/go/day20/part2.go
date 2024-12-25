package day20

import (
	"fmt"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day20/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
	// Get grid
	grid, err := utils.ConstructGrid(inputFile)
	if err != nil {
		return 0, err
	}

	// Get start and end
	src := [2]int{0, 0}
	dest := [2]int{0, 0}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				src = [2]int{i, j}
				grid[i][j] = '.'
			} else if cell == 'E' {
				dest = [2]int{i, j}
				grid[i][j] = '.'
			}
		}
	}

	// Get shortest distance
	path := calculatePath(grid, src, dest)

	// Cheat threshold
	threshold := len(path) - 50
	// Any cheat with a distance < threshold is valid
	cheats := map[int]int{} // distance -> count

	// Calculate possible cheats for each step in the path
	for _, node := range path {
		getValidCheats(grid, node, threshold, path, cheats)
	}

	for k, v := range cheats {
		fmt.Println(v, " cheats that save ", len(path)-k, " steps")
	}

	return 0, nil
}

func getValidCheats(grid [][]rune, node [2]int, threshold int, path [][2]int, cheats map[int]int) int {
	// A cheat allows walker to enter a wall and then take upto 19 steps before returning to the path
	// The destination tile after a cheat must be in the path and not a wall

	// Make a set of path tiles
	pathSet := hashset.New()
	for _, node := range path {
		pathSet.Add(node)
	}

	// Find adjacent walls
	for _, dir := range utils.Directions {
		di := node[0] + dir[0]
		dj := node[1] + dir[1]

		if grid[di][dj] != '#' {
			continue
		}

		// Start a cheat from this wall
		dests := runCheat(grid, [2]int{di, dj}, pathSet)

		// Check if any of the destinations are within the threshold
		for dest, dist := range dests {
			// Check if dest is the same as node
			if dest == node {
				continue
			}

			// Count tiles between node and dest in path
			count := -1
			for _, pathNode := range path {
				if pathNode == node {
					// Start counting
					count = 0
				} else if pathNode == dest {
					// Stop counting
					break
				}
				if count >= 0 {
					count++
				}
			}

			// Calculate distance after cheat
			newDist := len(path) - count + dist
			if newDist <= threshold {
				cheats[newDist]++
			}
		}
	}

	return 0
}

func runCheat(grid [][]rune, wall [2]int, path *hashset.Set) map[[2]int]int {
	// Return type -> map[dest]distance
	// BFS from wall to find all reachable path tiles and their distances
	queue := arrayqueue.New() // queue[di, dj, distance]
	visited := hashset.New()
	queue.Enqueue([3]int{wall[0], wall[1], 0})

	res := map[[2]int]int{} // dest -> distance
	for !queue.Empty() {
		// Dequeue
		_node, _ := queue.Dequeue()
		node := _node.([3]int)

		// Add to visited
		visited.Add([2]int{node[0], node[1]})

		// If distance is > 19, return
		if node[2] > 19 {
			continue
		}

		// Explore adjacent tiles
		for _, dir := range utils.Directions {
			di := node[0] + dir[0]
			dj := node[1] + dir[1]

			// Check if tile is valid
			if !utils.InGrid(grid, [2]int{di, dj}) || visited.Contains([2]int{di, dj}) {
				continue
			}

			// Check if tile is a path tile
			if grid[di][dj] == '.' && path.Contains([2]int{di, dj}) {
				// Add to results
				oldDist, ok := res[[2]int{di, dj}]
				if !ok || node[2]+1 < oldDist {
					res[[2]int{di, dj}] = node[2] + 1
				}
			}

			// Check if tile is a wall
			if grid[di][dj] == '#' {
				// Add to queue
				queue.Enqueue([3]int{di, dj, node[2] + 1})
			}
		}
	}

	return res
}
