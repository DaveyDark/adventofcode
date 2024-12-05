package day05

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day5/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	// Read file
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	// Parse rules
	rules := map[int]hashset.Set{}
	ruleRegex := regexp.MustCompile(`(\d+)\|(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		// Get key and value
		match := ruleRegex.FindStringSubmatch(line)
		key, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}
		value, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		// Add to rules map
		rule, ok := rules[key]
		if !ok {
			rules[key] = *hashset.New(value)
		} else {
			rule.Add(value)
		}
	}

	// Verify updates
	result := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ",")
		// Gather update
		update := []int{}
		for _, u := range lineSplit {
			upd, err := strconv.Atoi(u)
			if err != nil {
				return 0, nil
			}
			update = append(update, upd)
		}
		if !verifyUpdate(update, rules) {
			continue
		}
		// Get middle element
		result += int64(update[len(update)/2])
	}
	return result, nil
}
