package day10

import (
	"bufio"
	"os"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

var adjacent = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func init() {
	registry.Registry["day10/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	// Read file
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Parse file into grid
	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	// Locate heads
	heads := [][2]int{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == '0' {
				heads = append(heads, [2]int{i, j})
			}
		}
	}

	// Locate trails from each head
	trails := int64(0)
	for _, head := range heads {
		trails += int64(countTrails(grid, head))
	}

	return trails, nil
}

func countTrails(grid [][]rune, head [2]int) int {
	// BFS to count trails
	ends := hashset.New()
	queue := arrayqueue.New()
	queue.Enqueue(head)
	for !queue.Empty() {
		cnt := queue.Size()
		for range cnt {
			cell, _ := queue.Dequeue()
			val := cell.([2]int)
			i, j := val[0], val[1]
			for _, adj := range adjacent {
				cell := [2]int{i + adj[0], j + adj[1]}
				if !inGrid(grid, cell[0], cell[1]) {
					continue
				}
				if grid[cell[0]][cell[1]] == rune(grid[i][j]+1) {
					if grid[cell[0]][cell[1]] == '9' {
						ends.Add([2]int{cell[0], cell[1]})
					} else {
						queue.Enqueue(cell)
					}
				}
			}
		}
	}

	return ends.Size()
}

func inGrid(grid [][]rune, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}
