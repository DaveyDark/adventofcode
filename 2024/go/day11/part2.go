package day11

import (
	"fmt"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day11/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
	// Get stones array from input
	stonesArr, err := parseInput(inputFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	stones := map[uint64]int{}     // Stones count of stones with given number
	cache := map[uint64][]uint64{} // Stores result of processing a stone
	// Convert to map
	for _, st := range stonesArr {
		stones[st]++
	}

	// Do each blink
	for i := 1; i <= 75; i++ {
		fmt.Println("Blink: ", i)

		next := map[uint64]int{} // To store updated state
		for key, val := range stones {
			// get cached result
			result, cached := cache[key]
			if !cached {
				// if result is not cached, we calculate and store it
				res := processStone(key)
				cache[key] = res
				result = res
			}
			// Add frequency to each of the resulting stones in next map
			for _, r := range result {
				next[r] += val
			}
		}
		// Change stones to be the updated map
		stones = next
	}

	// Sum up all the values in the stones map
	count := int64(0)
	for _, cnt := range stones {
		count += int64(cnt)
	}

	return count, nil
}
