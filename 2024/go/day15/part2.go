package day15

import (
	"os"
	"slices"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day15/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
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

	// Upscale grid
	grid = upscaleGrid(grid)

	// Find robot
	robot := findRobot(grid)

	// Move robot the given steps
	for _, step := range steps {
		robot = moveRobotWide(grid, robot, step)
	}

	// Calculate score of each obstacle and sum it up
	score := int64(0)
	for i, row := range grid {
		for j, cell := range row {
			if cell == '[' {
				score += int64((i * 100) + (j))
			}
		}
	}

	// Return the score
	return score, nil
}

func moveRobotWide(grid [][]rune, robot [2]int, step rune) [2]int {
	dir := getDirection(step)
	is_vertical := dir[0] != 0
	dest := [2]int{robot[0] + dir[0], robot[1] + dir[1]}

	// See if destination is a box
	if grid[dest[0]][dest[1]] == '[' || grid[dest[0]][dest[1]] == ']' {
		pos := dest
		// Move further along any obstacles
		for grid[pos[0]][pos[1]] == '[' || grid[pos[0]][pos[1]] == ']' {
			pos[0] += dir[0]
			pos[1] += dir[1]
		}
		// Check if we found space to move the obstacle
		if grid[pos[0]][pos[1]] == '.' {
			if is_vertical {
				// Move obstacle vertically; This could move multiple columns of obstacles
				box := []int{}
				if grid[dest[0]][dest[1]] == '[' {
					box = append(box, dest[1], dest[1]+1)
				} else {
					box = append(box, dest[1]-1, dest[1])
				}
				if canMoveBoxVertically(grid, dest[0], dir[0], box) {
					moveBoxVertically(grid, dest[0], dir[0], box, []int{})
					// Boxes moved successfully
					return dest
				} else {
					// Cannot perform move
					return robot
				}
			} else {
				// Move obstacle horizontally
				row := grid[pos[0]]
				row = slices.Delete(row, pos[1], pos[1]+1)
				row = slices.Insert(row, dest[1], '.')
				grid[pos[0]] = row
			}
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

func canMoveBoxVertically(grid [][]rune, row int, dir int, boxes []int) bool {
	overlaps := hashset.New()
	for _, b := range boxes {
		if grid[row+dir][b] == '#' {
			// Wall in the way, cannot move
			return false
		} else if grid[row+dir][b] == '[' {
			// Box in the way, add to overlaps
			overlaps.Add(b, b+1)
		} else if grid[row+dir][b] == ']' {
			// Box in the way, add to overlaps
			overlaps.Add(b-1, b)
		}
	}
	if overlaps.Size() == 0 {
		// No overlaps, safe to move
		return true
	} else {
		// Convert overlapsSlice to []int
		overlapsSlice := []int{}
		for _, v := range overlaps.Values() {
			overlapsSlice = append(overlapsSlice, v.(int))
		}
		// Recursion to check overlapped boxes
		return canMoveBoxVertically(grid, row+dir, dir, overlapsSlice)
	}
}

func moveBoxVertically(grid [][]rune, row int, dir int, boxes []int, last []int) {
	overlaps := hashset.New()
	for _, col := range boxes {
		// Clear initial position
		if !slices.Contains(last, col) {
			grid[row][col] = '.'
		}
		// Check for overlapping boxes
		if grid[row+dir][col] == '[' {
			overlaps.Add(col, col+1)
		} else if grid[row+dir][col] == ']' {
			overlaps.Add(col-1, col)
		}
	}

	// Place boxes in new position
	for i, col := range boxes {
		if i%2 == 0 {
			grid[row+dir][col] = '['
		} else {
			grid[row+dir][col] = ']'
		}
	}

	// Handle any overlapping boxes recursively
	if overlaps.Size() > 0 {
		overlapsSlice := []int{}
		for _, v := range overlaps.Values() {
			overlapsSlice = append(overlapsSlice, v.(int))
		}
		slices.Sort(overlapsSlice)
		moveBoxVertically(grid, row+dir, dir, overlapsSlice, boxes)
	}
}

func upscaleGrid(grid [][]rune) [][]rune {
	// Upsacle grid to make all tiles except robot double wide
	newGrid := [][]rune{}
	for _, row := range grid {
		newRow := []rune{}
		for _, cell := range row {
			if cell == '@' {
				newRow = append(newRow, '@', '.')
			} else if cell == 'O' {
				// Larger box
				newRow = append(newRow, '[', ']')
			} else {
				newRow = append(newRow, cell, cell)
			}
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}
