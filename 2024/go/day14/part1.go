package day14

import (
	"bufio"
	"os"
	"regexp"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
)

const GRID_COLS = 101
const GRID_ROWS = 103

// Constants for sample input
// const GRID_COLS = 11
// const GRID_ROWS = 7
const SECONDS = 100

func init() {
	registry.Registry["day14/part1"] = solve
}

type Robot struct {
	position [2]int
	velocity [2]int
}

func solve(inputFile string) (int64, error) {
	// Prase input and create scanner
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	// Regex to parse input line
	inputRegex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	sums := [4]int64{0, 0, 0, 0}
	for scanner.Scan() {
		// Parse line via regex
		line := scanner.Text()
		captures := inputRegex.FindStringSubmatch(line)

		// Convert matches to int
		matches, err := utils.MapStrArrToInt(captures, 1, 4)
		if err != nil {
			return 0, err
		}

		// Construct robot from matches
		robot := Robot{[2]int{matches[1], matches[0]}, [2]int{matches[3], matches[2]}}

		// Move robot for 100 seconds and rap final position around the grid
		p0 := (int64(robot.position[0]) + int64(robot.velocity[0]*SECONDS)) % GRID_ROWS
		p1 := (int64(robot.position[1]) + int64(robot.velocity[1]*SECONDS)) % GRID_COLS
		if p0 < 0 {
			p0 = GRID_ROWS + p0
		}
		if p1 < 0 {
			p1 = GRID_COLS + p1
		}

		// Set final position
		robot.position = [2]int{int(p0), int(p1)}

		// Add to the respective quadrant's sum
		quadrant := -1
		if robot.position[0] < GRID_ROWS/2 && robot.position[1] < GRID_COLS/2 {
			quadrant = 0
		} else if robot.position[0] > GRID_ROWS/2 && robot.position[1] > GRID_COLS/2 {
			quadrant = 3
		} else if robot.position[0] > GRID_ROWS/2 && robot.position[1] < GRID_COLS/2 {
			quadrant = 2
		} else if robot.position[0] < GRID_ROWS/2 && robot.position[1] > GRID_COLS/2 {
			quadrant = 1
		}
		if quadrant == -1 {
			continue
		}
		sums[quadrant]++
	}

	return sums[0] * sums[1] * sums[2] * sums[3], nil
}
