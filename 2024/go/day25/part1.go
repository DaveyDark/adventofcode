package day25

import (
	"bufio"
	"os"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day25/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	keys := [][][]rune{}
	locks := [][][]rune{}

	// Read locks and keys
	for scanner.Scan() {
		block := [][]rune{}
		for range 6 {
			line := scanner.Text()
			block = append(block, []rune(line))
			scanner.Scan()
		}
		line := scanner.Text()
		block = append(block, []rune(line))
		scanner.Scan()
		if block[0][0] == '#' {
			locks = append(locks, block)
		} else {
			keys = append(keys, block)
		}
	}

	// For each lock
	fits := int64(0)
	for _, lock := range locks {
		// try each key
		slots := invertHeights(heights(lock, false))
		for _, key := range keys {
			pins := heights(key, true)
			// Check if key fits
			fit := true
			for i := range pins {
				if pins[i] > slots[i] {
					fit = false
					break
				}
			}
			if fit {
				fits++
			}
		}
	}

	return fits, nil
}

func heights(lock [][]rune, reverse bool) []int {
	heights := []int{}
	for j := range lock[0] {
		height := 0
		i := 0
		limit := len(lock) - 1
		increment := 1
		if reverse {
			i = len(lock) - 1
			limit = -1
			increment = -1
		}
		for i != limit {
			if lock[i][j] != '#' {
				break
			}
			height++
			i += increment
		}
		heights = append(heights, height)
	}
	return heights
}

func invertHeights(heights []int) []int {
	inverted := []int{}
	for _, h := range heights {
		inverted = append(inverted, 7-h)
	}
	return inverted
}
