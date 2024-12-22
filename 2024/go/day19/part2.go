package day19

import (
	"bufio"
	"os"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day19/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
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
		count += int64(towelCombinations(towel, colors, maxLen))
	}

	return count, nil
}

func towelCombinations(towel string, colors *hashset.Set, maxLen int) int {
	dp := make([]int, len(towel))

	for i := 0; i < len(towel); i++ {
		if i != 0 && dp[i-1] == 0 {
			continue
		}
		word := strings.Builder{}
		for j := i; j < i+maxLen && j < len(towel); j++ {
			word.WriteByte(towel[j])
			if colors.Contains(word.String()) {
				if i > 0 {
					dp[j] += dp[i-1]
				} else {
					dp[j] = 1
				}
			}
		}
	}

	return dp[len(towel)-1]
}
