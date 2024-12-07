package day07

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day7/part2"] = solve2
}

func isPossible2(result int64, variables []int64) bool {
	for i := 0; i < int(math.Pow(3, float64(len(variables)))); i++ {
		ans := int64(0)
		n := i
		for _, v := range variables {
			if n%3 == 0 {
				ans += v
			} else if n%3 == 1 {
				ans *= v
			} else {
				vStr := strconv.FormatInt(v, 10)
				ansStr := strconv.FormatInt(ans, 10)
				catStr := ansStr + vStr
				cat, err := strconv.ParseInt(catStr, 10, 64)
				if err != nil {
					continue
				}
				ans = cat
			}
			n /= 3
		}
		if ans == result {
			return true
		}
	}
	return false
}

func solve2(inputFile string) (int64, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	ans := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ":")
		res, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			return 0, err
		}
		variablesStr := strings.Fields(splits[1])
		variables := make([]int64, len(variablesStr))
		for i, v := range variablesStr {
			variables[i], err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				return 0, err
			}
		}

		if isPossible2(res, variables) {
			ans += res
		}
	}

	return ans, nil
}
