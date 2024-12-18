package day16

import (
	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

type Crawler struct {
	position  [2]int
	direction [2]int
	path      [][2]int
	distance  int
}

func init() {
	registry.Registry["day16/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
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

	// Calculate paths
	paths := calculatePaths(grid, source, dest, dist)

	return int64(paths), nil
}

func calculatePaths(grid [][]rune, source [2]int, dest [2]int, dist int) int {
	// BFS to calculate number of paths, stop when max distance is reached
	queue := arrayqueue.New()
	visited := map[[2][2]int]int{}
	queue.Enqueue(Crawler{source, [2]int{0, 1}, [][2]int{source}, 0})
	successfulCrawlers := []Crawler{}

	for !queue.Empty() {
		// Get node
		node, _ := queue.Dequeue()
		crawler := node.(Crawler)
		visited[[2][2]int{crawler.position, crawler.direction}] = crawler.distance

		// If we exceed max distance, stop
		if crawler.distance > dist {
			continue
		}

		// Check if we reached destination
		if crawler.position == dest {
			if crawler.distance == dist {
				successfulCrawlers = append(successfulCrawlers, crawler)
			}
			continue
		}

		// Check neighbors
		for _, adj := range utils.Directions {
			di := crawler.position[0] + adj[0]
			dj := crawler.position[1] + adj[1]

			visitDist, isVisited := visited[[2][2]int{{di, dj}, adj}]

			if grid[di][dj] != '.' || (isVisited && visitDist <= crawler.distance) {
				// if neighbor is not a valid cell or has been visited, continue
				continue
			}

			// Create new crawler
			newCrawler := Crawler{[2]int{di, dj}, adj, [][2]int{}, crawler.distance + 1}
			// Copy path into newCrawler
			newCrawler.path = append(newCrawler.path, crawler.path...)

			// Check if we changed direction
			if adj != crawler.direction {
				if adj[0]+crawler.direction[0] == 0 && adj[1]+crawler.direction[1] == 0 {
					newCrawler.distance += 2000
				} else {
					newCrawler.distance += 1000
				}
			}
			newCrawler.path = append(newCrawler.path, [2]int{di, dj})
			queue.Enqueue(newCrawler)
		}
	}

	// Count number of paths
	pathSet := hashset.New()
	for _, crawler := range successfulCrawlers {
		for _, path := range crawler.path {
			pathSet.Add(path)
		}
	}

	return pathSet.Size()
}
