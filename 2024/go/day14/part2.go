package day14

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"os"
	"regexp"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"golang.org/x/image/bmp"
)

func init() {
	registry.Registry["day14/part2"] = solve2
}

func moveRobots(robots []Robot) {
	for i, robot := range robots {
		p0 := (robot.position[0] + robot.velocity[0]) % GRID_ROWS
		p1 := (robot.position[1] + robot.velocity[1]) % GRID_COLS
		if p0 < 0 {
			p0 = GRID_ROWS + p0
		}
		if p1 < 0 {
			p1 = GRID_COLS + p1
		}
		robots[i] = Robot{[2]int{p0, p1}, robot.velocity}
	}
}

func buildGrid(robots []Robot) [GRID_ROWS][GRID_COLS]rune {
	grid := [GRID_ROWS][GRID_COLS]rune{}
	for _, robot := range robots {
		if grid[robot.position[0]][robot.position[1]] == 0 {
			grid[robot.position[0]][robot.position[1]] = 'X'
		}
	}
	return grid
}

func gridToString(grid [GRID_ROWS][GRID_COLS]rune) string {
	var result string
	for _, row := range grid {
		for _, cell := range row {
			if cell != 0 {
				result += string(cell)
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}

func solve2(inputFile string) (int64, error) {
	// Prase input and create scanner
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// Regex to parse input line
	inputRegex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	robots := []Robot{}
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
		robots = append(robots, Robot{[2]int{matches[1], matches[0]}, [2]int{matches[3], matches[2]}})
	}

	seconds := int64(1)
	largest := [2]int64{0, 0} // [seconds, size]
	largestGrid := ""
	for i := 0; i < 15000; i++ {
		fmt.Println("Seconds: ", i)
		moveRobots(robots)
		grid := buildGrid(robots)
		gridStr := gridToString(grid)
		size := int64(findLargestIsland(grid))
		if size > largest[1] {
			largest[0] = seconds
			largest[1] = size
			largestGrid = gridStr
		}

		seconds++
	}

	err = saveGrid(largestGrid)
	if err != nil {
		return 0, err
	}

	return largest[0], nil
}

func findLargestIsland(grid [GRID_ROWS][GRID_COLS]rune) int {
	visited := make([][]bool, GRID_ROWS)
	for i := range visited {
		visited[i] = make([]bool, GRID_COLS)
	}

	maxSize := 0

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		// Check bounds and if already visited
		if row < 0 || row >= GRID_ROWS || col < 0 || col >= GRID_COLS ||
			visited[row][col] || grid[row][col] != 'X' {
			return 0
		}

		// Mark as visited
		visited[row][col] = true
		size := 1

		// Check adjacent cells (up, down, left, right)
		size += dfs(row-1, col) // up
		size += dfs(row+1, col) // down
		size += dfs(row, col-1) // left
		size += dfs(row, col+1) // right

		return size
	}

	// Search each cell in grid
	for i := 0; i < GRID_ROWS; i++ {
		for j := 0; j < GRID_COLS; j++ {
			if !visited[i][j] && grid[i][j] == 'X' {
				islandSize := dfs(i, j)
				if islandSize > maxSize {
					maxSize = islandSize
				}
			}
		}
	}

	return maxSize
}

func saveGrid(gridStr string) error {
	img := image.NewRGBA(image.Rect(0, 0, GRID_COLS, GRID_ROWS))
	for y, row := range strings.Split(gridStr, "\n") {
		for x, char := range row {
			if char == 'X' {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	// Save bitmap file
	os.MkdirAll("tmp", 0755)
	f, err := os.Create(fmt.Sprintf("tmp/grid.bmp"))
	if err != nil {
		return err
	}
	bmp.Encode(f, img)
	f.Close()
	return nil
}
