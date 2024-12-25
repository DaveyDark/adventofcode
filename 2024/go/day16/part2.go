package day16

import (
	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
)

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

	// Calculate paths
	paths := calculatePaths(grid, source, dest)

	return int64(paths), nil
}
