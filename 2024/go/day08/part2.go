package day08

import (
	"bufio"
	"os"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day8/part2"] = solve2
}

func findAntinodesAnyDist(a [2]int, b [2]int, grid [][]rune) [][2]int {
	x_dist, y_dist := a[1]-b[1], a[0]-b[0]
	node := [2]int{a[1] + x_dist, a[0] + y_dist}
	nodes := [][2]int{}
	for inGrid(grid, node) {
		nodes = append(nodes, node)
		node = [2]int{node[0] + x_dist, node[1] + y_dist}
	}
	node = [2]int{b[1] - x_dist, b[0] - y_dist}
	for inGrid(grid, node) {
		nodes = append(nodes, node)
		node = [2]int{node[0] - x_dist, node[1] - y_dist}
	}
	return nodes
}

func solve2(inputFile string) (int64, error) {
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
		calculateAntinodes(ant, grid, findAntinodesAnyDist)
	}

	for _, row := range grid {
		for _, cell := range row {
			if cell != '.' {
				antinodes++
			}
		}
	}

	return antinodes, nil
}
