package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day3/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
	// Read input file
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	// Parse input line by line
	sum := int64(0)
	mulRegexStr := `mul\((\d{1,3}),(\d{1,3})\)`
	mulRegex := regexp.MustCompile(mulRegexStr)
	searchRegex := regexp.MustCompile(mulRegexStr + `|do\(\)|don't\(\)`)
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := searchRegex.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println(match)
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else if enabled {
				capture := mulRegex.FindStringSubmatch(match)
				x, err := strconv.Atoi(capture[1])
				if err != nil {
					return 0, err
				}
				y, err := strconv.Atoi(capture[2])
				if err != nil {
					return 0, err
				}
				sum += int64(x * y)
			}
		}
	}
	return sum, nil
}
