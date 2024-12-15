package day15

import (
	"os"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
)

func init() {
	registry.Registry["day15/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	// Parse input into string
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}
	input := string(inputBytes)

	// Split into two parts - grid and steps
	inputSplit := strings.Split(input, "\n\n")
	inputSplit[1] = strings.Replace(inputSplit[1], "\n", "", -1)
	steps := []rune(inputSplit[1])

	// Construct grid
	grid, err := utils.ContructGridFromStr(inputSplit[0])
	if err != nil {
		return 0, err
	}

	// Find robot
	robot := findRobot(grid)

	// Move robot the given steps
	for _, step := range steps {
		robot = moveRobot(grid, robot, step)
	}

	// Calculate score of each obstacle and sum it up
	score := int64(0)
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'O' {
				score += int64((i * 100) + (j))
			}
		}
	}

	// Return the score
	return score, nil
}

func findRobot(grid [][]rune) [2]int {
	// Find location of the robot in the grid and remove the @ from the grid
	for i, row := range grid {
		for j, cell := range row {
			if cell == '@' {
				grid[i][j] = '.'
				return [2]int{i, j}
			}
		}
	}
	panic("Robot not found in input grid")
}

func moveRobot(grid [][]rune, robot [2]int, step rune) [2]int {
	var dir [2]int = getDirection(step)
	dest := [2]int{robot[0] + dir[0], robot[1] + dir[1]}

	// See if destination is a box
	if grid[dest[0]][dest[1]] == 'O' {
		pos := dest
		// Move further along any obstacles
		for grid[pos[0]][pos[1]] == 'O' {
			pos[0] += dir[0]
			pos[1] += dir[1]
		}
		// Check if we found space to move the obstacle
		if grid[pos[0]][pos[1]] == '.' {
			// Move the obstacle to this space
			grid[pos[0]][pos[1]] = 'O'
			grid[dest[0]][dest[1]] = '.'
			return dest
		}
		// No space was found so we return the original position
		return robot
	} else if grid[dest[0]][dest[1]] == '.' {
		// Move the robot to available position
		return dest
	}
	// Move not possible; Return initial position
	return robot
}

func getDirection(step rune) [2]int {
	// Get movement direction as [2]int from step
	switch step {
	case '^':
		return utils.Directions[0]
	case '>':
		return utils.Directions[1]
	case 'v':
		return utils.Directions[2]
	case '<':
		return utils.Directions[3]
	}
	panic("Invalid step in input")
}
