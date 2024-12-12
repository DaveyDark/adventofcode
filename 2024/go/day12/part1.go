package day12

import (
	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

var adjacent = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func init() {
	registry.Registry["day12/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	// Get grid
	grid, err := utils.ConstructGrid(inputFile)
	if err != nil {
		return 0, err
	}

	price := int64(0)
	// Traverse grid
	for i, row := range grid {
		for j, cell := range row {
			if cell != '.' {
				// Get area and perimeter
				regionArea, regionPerimeter := evaluateRegion(grid, i, j)
				price += regionArea * regionPerimeter
			}
		}
	}

	return price, nil
}

func evaluateRegion(grid [][]rune, i, j int) (int64, int64) {
	perimeter := int64(0)
	region := captureRegion(grid, i, j)
	area := int64(region.Size())

	for _, _node := range region.Values() {
		node := _node.([2]int)
		walls := 0
		for _, adj := range adjacent {
			di := node[0] + adj[0]
			dj := node[1] + adj[1]
			if !utils.InGrid(grid, [2]int{di, dj}) || !region.Contains([2]int{di, dj}) {
				walls++
			}
		}
		perimeter += int64(walls)
	}

	return area, perimeter
}

func captureRegion(grid [][]rune, i, j int) hashset.Set {
	region := hashset.New([2]int{i, j})
	queue := arrayqueue.New()
	queue.Enqueue([2]int{i, j})
	plant := grid[i][j]
	grid[i][j] = '.'

	for !queue.Empty() {
		size := queue.Size()
		for range size {
			_node, _ := queue.Dequeue()
			node := _node.([2]int)
			for _, adj := range adjacent {
				di := node[0] + adj[0]
				dj := node[1] + adj[1]
				if !utils.InGrid(grid, [2]int{di, dj}) {
					continue
				}
				if grid[di][dj] == plant && !region.Contains([2]int{di, dj}) {
					grid[di][dj] = '.'
					region.Add([2]int{di, dj})
					queue.Enqueue([2]int{di, dj})
				}
			}
		}
	}

	return *region
}
