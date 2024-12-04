package day04

import (
	"bufio"
	"os"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day4/part2"] = solve2
}

func findMas(grid [][]rune, i int, j int) bool {
	str1 := string(grid[i-1][j-1]) + string(grid[i][j]) + string(grid[i+1][j+1])
	str2 := string(grid[i+1][j-1]) + string(grid[i][j]) + string(grid[i-1][j+1])
	return (str1 == "MAS" || str1 == "SAM") && (str2 == "MAS" || str2 == "SAM")
}

func solve2(inputFile string) (int64, error) {
	// Read input file
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	// Parse input line by line
	grid := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	count := int64(0)

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			if grid[i][j] == 'A' && findMas(grid, i, j) {
				count++
			}
		}
	}

	return count, nil
}
