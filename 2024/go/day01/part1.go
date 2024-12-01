package day01

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day1/part1"] = solve
}

func solve(inputFile string) (int64, error) {
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

	// Sort the arrays
	sort.Ints(leftArr)
	sort.Ints(rightArr)

	// Calculate difference
	diff := int64(0)
	for i := 0; i < len(leftArr); i++ {
		diff += int64(math.Abs(float64(leftArr[i]) - float64(rightArr[i])))
	}

	return diff, nil
}
