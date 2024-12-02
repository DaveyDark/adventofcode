package day02

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day2/part1"] = solve
}

func verifyRecords(records []int) bool {
	last := records[0]
	delta := records[1] - records[0]
	for _, n := range records[1:] {
		diff := n - last

		if (diff > 0 && delta < 0) || (diff < 0 && delta > 0) || (diff == 0) {
			return false
		}
		if math.Abs(float64(diff)) > 3.0 {
			return false
		}
		last = n
	}
	return true
}

func solve(inputFile string) (int64, error) {
	// Read input inputFile
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Process input
	safeCount := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		records := []int{}
		for _, field := range fields {
			record, err := strconv.Atoi(field)
			if err != nil {
				return 0, err
			}
			records = append(records, record)
		}
		if verifyRecords(records) {
			safeCount++
		}
	}

	return safeCount, nil
}
