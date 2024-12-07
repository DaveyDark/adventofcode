package day06

import (
	"bufio"
	"os"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day6/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	// Read inputFile
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	// Parse lines
	grid := [][]rune{}
	guard := [2]int{0, 0}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		idx := strings.Index(line, "^")
		if idx != -1 {
			guard[0] = idx
			guard[1] = i
		}
		grid = append(grid, []rune(line))
		i++
	}

	dirs := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	dir := 0
	spots := hashset.New()
	for guard[0]+dirs[dir][0] >= 0 && guard[0]+dirs[dir][0] < len(grid[0]) && guard[1]+dirs[dir][1] >= 0 && guard[1]+dirs[dir][1] < len(grid) {
		spots.Add(guard)
		if grid[guard[1]+dirs[dir][1]][guard[0]+dirs[dir][0]] == '#' {
			dir = (dir + 1) % 4
		}
		guard[0] += dirs[dir][0]
		guard[1] += dirs[dir][1]
	}
	spots.Add(guard)

	return int64(spots.Size()), nil
}
