package day19

import (
	"bufio"
	"os"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day19/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	input, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	colors, maxLen := getColors(scanner.Text())
	scanner.Scan()

	count := int64(0)
	for scanner.Scan() {
		towel := scanner.Text()
		if isTowelPossible(towel, colors, maxLen) {
			count++
		}
	}

	return count, nil
}

func isTowelPossible(towel string, colors *hashset.Set, maxLen int) bool {
	dp := make([]bool, len(towel))

	for i := 0; i < len(towel); i++ {
		if i != 0 && !dp[i-1] {
			continue
		}
		word := strings.Builder{}
		for j := i; j < i+maxLen && j < len(towel); j++ {
			word.WriteByte(towel[j])
			if colors.Contains(word.String()) {
				dp[j] = true
			}
		}
	}

	return dp[len(towel)-1]
}

func getColors(str string) (*hashset.Set, int) {
	colors := hashset.New()
	maxLen := 0
	for _, s := range strings.Split(str, ",") {
		colors.Add(strings.Trim(s, " "))
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}
	return colors, maxLen
}
