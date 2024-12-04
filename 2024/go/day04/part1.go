package day04

import (
	"bufio"
	"os"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day4/part1"] = solve
}

func checkRows(grid [][]rune) int64 {
	// Check rows
	count := int64(0)
	for _, row := range grid {
		rowStr := string(row)
		rowCount := strings.Count(rowStr, "XMAS")
		rowCount += strings.Count(rowStr, "SAMX")
		count += int64(rowCount)
	}
	return count
}

func checkColumns(grid [][]rune) int64 {
	// Check columns
	count := int64(0)
	for i := 0; i < len(grid[0]); i++ {
		colStr := ""
		for j := 0; j < len(grid); j++ {
			colStr += string(grid[j][i])
		}
		colCount := strings.Count(colStr, "XMAS")
		colCount += strings.Count(colStr, "SAMX")
		count += int64(colCount)
	}
	return count
}
func checkDiagonals(grid [][]rune) int64 {
	// Check diagonals
	count := int64(0)
	for i := 0; i < len(grid); i++ {
		// Top left to bottom right diagonals
		diagStr := ""
		for j := 0; j+i < len(grid); j++ {
			diagStr += string(grid[i+j][j])
		}
		diagCount := int64(strings.Count(diagStr, "XMAS") + strings.Count(diagStr, "SAMX"))
		if i != 0 {
			diagStr = ""
			for j := 0; j+i < len(grid); j++ {
				diagStr += string(grid[j][i+j])
			}
			diagCount += int64(strings.Count(diagStr, "XMAS") + strings.Count(diagStr, "SAMX"))
		}

		// Top right to bottom left diagonals
		diagStr = ""
		for j := 0; j <= i; j++ {
			diagStr += string(grid[j][i-j])
		}
		diagCount += int64(strings.Count(diagStr, "XMAS") + strings.Count(diagStr, "SAMX"))
		diagStr = ""
		for j := 0; j <= i+len(grid); j++ {
			if j >= len(grid) || i+len(grid)-j >= len(grid) {
				continue
			}
			diagStr += string(grid[j][i+len(grid)-j])
		}
		diagCount += int64(strings.Count(diagStr, "XMAS") + strings.Count(diagStr, "SAMX"))

		count += diagCount
	}
	return count
}
func solve(inputFile string) (int64, error) {
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

	count += checkRows(grid)
	count += checkColumns(grid)
	count += checkDiagonals(grid)

	return count, nil
}
