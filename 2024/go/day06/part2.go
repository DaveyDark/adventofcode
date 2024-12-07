package day06

import (
	"bufio"
	"os"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day6/part2"] = solve2
}

// checks if the (y, x) position is within the bounds of the grid.
func inGrid(grid [][]rune, pos [2]int) bool {
	y, x := pos[0], pos[1]
	return y >= 0 && x >= 0 && y < len(grid) && x < len(grid[0])
}

// moves the guard one step in the current direction.
// If an obstacle is encountered, it turns right and tries the next direction.
func advance(grid [][]rune, pos [2]int, dirs [4][2]int, dir int) ([2]int, int) {
	y, x := pos[0], pos[1]
	dy, dx := dirs[dir][0], dirs[dir][1]
	next := [2]int{y + dy, x + dx}

	for {
		// Check if next position is out of bounds
		if !inGrid(grid, next) {
			return next, dir
		}
		// If next position is not an obstacle, proceed
		if grid[next[0]][next[1]] != '#' {
			break
		}
		// Else, turn right and try the next direction
		dir = (dir + 1) % 4
		dy, dx = dirs[dir][0], dirs[dir][1]
		next = [2]int{y + dy, x + dx}
	}

	return next, dir
}

// checks if placing an obstacle at ob creates a cycle in the guard's movement.
func createsCycle(grid [][]rune, dirs [4][2]int, origin [2]int, ob [2]int) bool {
	if grid[ob[0]][ob[1]] == '^' {
		return false
	}

	// Place the obstacle on the grid
	grid[ob[0]][ob[1]] = '#'

	// Initialize two guards for cycle detection (slow and fast)
	guard1, guard2 := origin, origin
	dir1, dir2 := 0, 0

	for inGrid(grid, guard1) && inGrid(grid, guard2) {
		// Move slow guard by one step
		guard1, dir1 = advance(grid, guard1, dirs, dir1)

		// Move fast guard by two steps
		guard2, dir2 = advance(grid, guard2, dirs, dir2)
		guard2, dir2 = advance(grid, guard2, dirs, dir2)

		// Check if both guards meet
		if guard1 == guard2 && dir1 == dir2 {
			grid[ob[0]][ob[1]] = '.'
			return true
		}
	}

	grid[ob[0]][ob[1]] = '.'
	return false
}

func solve2(inputFile string) (int64, error) {
	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Parse the input grid and locate the guard's starting position
	grid := [][]rune{}
	var guard [2]int
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		if idx := strings.Index(line, "^"); idx != -1 {
			guard = [2]int{y, idx}
		}
		grid = append(grid, []rune(line))
	}

	// Directions (y, x)
	dirs := [4][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}
	dir := 0 // Initially facing Up
	path := hashset.New()
	origin := guard

	for inGrid(grid, guard) {
		// Calculate the next position and direction
		path.Add(guard)
		guard, dir = advance(grid, guard, dirs, dir)
	}

	ans := 0
	for _, pos := range path.Values() {
		position := pos.([2]int) // Type assertion from interface{} to [2]int
		if createsCycle(grid, dirs, origin, position) {
			ans++
		}
	}

	return int64(ans), nil
}
