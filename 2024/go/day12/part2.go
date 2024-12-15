package day12

import (
	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
)

func init() {
	registry.Registry["day12/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
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
				regionArea, regionPerimeter := evaluateRegionAdvanced(grid, i, j)
				price += regionArea * regionPerimeter
			}
		}
	}

	return price, nil
}

func evaluateRegionAdvanced(grid [][]rune, i, j int) (int64, int64) {
	perimeter := int64(0)
	region := captureRegion(grid, i, j)
	area := int64(region.Size())

	// Find limits of region
	rowLimits, colLimits := [2]int{0, 0}, [2]int{0, 0} // [min, max]
	for _, _node := range region.Values() {
		node := _node.([2]int)
		if node[0] < rowLimits[0] {
			rowLimits[0] = node[0]
		}
		if node[0] > rowLimits[1] {
			rowLimits[1] = node[0]
		}
		if node[1] < colLimits[0] {
			colLimits[0] = node[1]
		}
		if node[1] > colLimits[1] {
			colLimits[1] = node[1]
		}
	}

	// Find edge count, counting adjacent edges as one
	for i := rowLimits[0]; i <= rowLimits[1]; i++ {
		for j := colLimits[0]; j <= colLimits[1]; j++ {
			// Top edge
			if region.Contains([2]int{i, j}) && !region.Contains([2]int{i - 1, j}) {
				if region.Contains([2]int{i - 1, j - 1}) || !region.Contains([2]int{i, j - 1}) {
					perimeter++
				}
			}
			// Bottom edge
			if region.Contains([2]int{i, j}) && !region.Contains([2]int{i + 1, j}) {
				if region.Contains([2]int{i + 1, j - 1}) || !region.Contains([2]int{i, j - 1}) {
					perimeter++
				}
			}
			// Left edge
			if region.Contains([2]int{i, j}) && !region.Contains([2]int{i, j - 1}) {
				if region.Contains([2]int{i - 1, j - 1}) || !region.Contains([2]int{i - 1, j}) {
					perimeter++
				}
			}
			// Right edge
			if region.Contains([2]int{i, j}) && !region.Contains([2]int{i, j + 1}) {
				if region.Contains([2]int{i - 1, j + 1}) || !region.Contains([2]int{i - 1, j}) {
					perimeter++
				}
			}
		}
	}

	return area, perimeter
}
