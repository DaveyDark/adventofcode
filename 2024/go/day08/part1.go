package day08

import (
	"bufio"
	"os"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day8/part1"] = solve
}

func inGrid(grid [][]rune, pos [2]int) bool {
	return pos[1] >= 0 && pos[1] < len(grid) && pos[0] >= 0 && pos[0] < len(grid[0])
}

func findAntinodes(a [2]int, b [2]int, grid [][]rune) [][2]int {
	x_dist, y_dist := a[1]-b[1], a[0]-b[0]
	a1 := [2]int{a[1] + x_dist, a[0] + y_dist}
	a2 := [2]int{b[1] - x_dist, b[0] - y_dist}
	nodes := [][2]int{}
	if inGrid(grid, a1) {
		nodes = append(nodes, a1)
	}
	if inGrid(grid, a2) {
		nodes = append(nodes, a2)
	}
	return nodes
}

func calculateAntinodes(ant [][2]int, grid [][]rune, findFunc func([2]int, [2]int, [][]rune) [][2]int) {
	for i := range ant {
		for j := i + 1; j < len(ant); j++ {
			nodes := findFunc(ant[i], ant[j], grid)
			for _, node := range nodes {
				grid[node[1]][node[0]] = 'X'
			}
		}
	}
}

func solve(inputFile string) (int64, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	antennas := map[rune][][2]int{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == '.' {
				continue
			}
			antennas[cell] = append(antennas[cell], [2]int{i, j})
		}
	}

	antinodes := int64(0)
	for _, ant := range antennas {
		calculateAntinodes(ant, grid, findAntinodes)
	}

	for _, row := range grid {
		for _, cell := range row {
			if cell == 'X' {
				antinodes++
			}
		}
	}
	return antinodes, nil
}
