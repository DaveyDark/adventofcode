package day03

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day3/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	// Read input file
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	// Parse input line by line
	sum := int64(0)
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for scanner.Scan() {
		line := scanner.Text()
		mulOps := mulRegex.FindAllStringSubmatch(line, -1)
		for _, op := range mulOps {
			// Access the two numbers
			a, err := strconv.Atoi(op[1])
			if err != nil {
				return 0, err
			}
			b, err := strconv.Atoi(op[2])
			if err != nil {
				return 0, err
			}
			sum += int64(a * b)
		}
	}
	return sum, nil
}
