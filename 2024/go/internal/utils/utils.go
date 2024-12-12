package utils

import (
	"bufio"
	"os"
)

func InGrid(grid [][]rune, pos [2]int) bool {
	return pos[1] >= 0 && pos[1] < len(grid) && pos[0] >= 0 && pos[0] < len(grid[0])
}

func ConstructGrid(inputFile string) ([][]rune, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid, nil
}