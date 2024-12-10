package day10

import (
	"bufio"
	"os"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

func init() {
	registry.Registry["day10/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
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
		trails += int64(countTrailsDistinct(grid, head))
	}

	return trails, nil
}

func countTrailsDistinct(grid [][]rune, head [2]int) int {
	// BFS to count trails
	trails := 0
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
						trails++
					} else {
						queue.Enqueue(cell)
					}
				}
			}
		}
	}

	return trails
}
