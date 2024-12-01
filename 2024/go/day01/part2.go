package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day1/part2"] = solve2
}

func count(arr []int, n int) int {
	count := 0
	for _, x := range arr {
		if x == n {
			count++
		}
	}
	return count
}

func solve2(inputFile string) (int64, error) {
	// Read contents of input inputFile
	file, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}
	input := string(file)

	// Parse input line by line and store in arrays
	leftArr := []int{}
	rightArr := []int{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		nums := strings.Fields(line)
		if len(nums) != 2 {
			return 0, fmt.Errorf("invalid line format: %s", line)
		}
		left, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		leftArr = append(leftArr, int(left))
		rightArr = append(rightArr, int(right))
	}

	// Init map
	freq := make(map[int]int)

	// Record frequency of each element in leftArr
	for _, n := range leftArr {
		_, ok := freq[n]
		if ok {
			continue
		} else {
			freq[n] = count(rightArr, n)
		}
	}

	// Calculate similarity score
	score := int64(0)
	for _, n := range leftArr {
		f, ok := freq[n]
		if !ok {
			return 0, fmt.Errorf("Unexpected Error: element %d not found in leftArr", n)
		}
		score += int64(n * f)
	}

	return score, nil
}
