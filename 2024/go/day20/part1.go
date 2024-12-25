package day20

import (
	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day20/part1"] = solve
}

func solve(inputFile string) (int64, error) {
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

	// Find isolated walls along path
	walls := findIsolatedWalls(grid, path)

	// Cheats map
	cheats := map[int]int{} // path length -> number of cheats

	// Remove each isolated wall and calculate path
	for _, wall := range walls.Values() {
		// Remove wall
		wallNode := wall.([2]int)
		grid[wallNode[0]][wallNode[1]] = '.'

		// Calculate new path
		newPath := calculatePath(grid, src, dest)

		// See if new path is shorter than threshold
		if len(newPath) < len(path) {
			cheats[len(newPath)]++
		}
		grid[wallNode[0]][wallNode[1]] = '#'
	}

	// Calculate valid cheat threshold
	threshold := len(path) - 100

	// Any cheat shorter than threshold is valid
	validCheats := 0
	for cheat, count := range cheats {
		if cheat <= threshold {
			validCheats += count
		}
	}

	return int64(validCheats), nil
}

func findIsolatedWalls(grid [][]rune, path [][2]int) *hashset.Set {
	// Use a hashset to store walls and ensure no duplicates
	walls := hashset.New()

	// Go through each node in path
	for _, node := range path {
		// A neighboring tile is an isolated wall if it is a wall
		// and it has a neighbor(excluding the initial tile) that is a floor
		for _, dir := range utils.Directions {
			di := node[0] + dir[0]
			dj := node[1] + dir[1]

			// Loop through neighbors of di, dj
			for _, dir2 := range utils.Directions {
				ddi := di + dir2[0]
				ddj := dj + dir2[1]

				if !utils.InGrid(grid, [2]int{ddi, ddj}) || ddi == node[0] && ddj == node[1] {
					continue
				}

				if grid[di][dj] == '#' && grid[ddi][ddj] == '.' {
					walls.Add([2]int{di, dj})
				}
			}
		}
	}

	return walls
}

func calculatePath(grid [][]rune, src, dest [2]int) [][2]int {
	// Calculate shortest distance in grid using Dijkstra with backrefs
	distances := map[[2]int][3]int{} // (x, y) -> (distance, backref_x, backref_y)
	pq := priorityqueue.NewWith(func(n1, n2 interface{}) int {
		a := n1.([2]int)
		b := n2.([2]int)
		return distances[a][0] - distances[b][0]
	})
	visited := hashset.New()

	// Init queue with source
	pq.Enqueue(src)
	distances[src] = [3]int{0, -1, -1}

	for !pq.Empty() {
		// Dequeue a node
		_node, _ := pq.Dequeue()
		node := _node.([2]int)

		// Add to visited set
		visited.Add([2]int{node[0], node[1]})

		// Explore neighbors
		for _, dir := range utils.Directions {
			di := node[0] + dir[0]
			dj := node[1] + dir[1]

			// Check if node is valid
			if grid[di][dj] != '.' || visited.Contains([2]int{di, dj}) {
				continue
			}

			// Update distance
			oldDist, exists := distances[[2]int{di, dj}]
			if !exists || distances[node][0]+1 < oldDist[0] {
				distances[[2]int{di, dj}] = [3]int{distances[node][0] + 1, node[0], node[1]}
			}

			// Enqueue node
			pq.Enqueue([2]int{di, dj})
		}
	}

	// Construct path
	path := [][2]int{}
	node := dest
	for node != src {
		path = append(path, node)
		node = [2]int{distances[node][1], distances[node][2]}
	}

	return path
}
