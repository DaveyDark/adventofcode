package day21

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/day21/keypad"
	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day21/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	type result struct {
		value int64
		err   error
	}

	resultChan := make(chan result)
	var lineCount int

	for scanner.Scan() {
		lineCount++
		input := scanner.Text()

		go func(input string) {
			coefficientStr := strings.TrimRight(input, "A")
			coefficient, err := strconv.Atoi(coefficientStr)
			if err != nil {
				resultChan <- result{0, err}
				return
			}

			// Parse input numbers to directions using numeric keypad
			presses := keypad.MapString(input, keypad.NUMERIC)
			for i := range 25 {
				presses = keypad.MapString(presses, keypad.DIRECTION)
				fmt.Println(i, input, len(presses))
			}

			resultChan <- result{int64(len(presses) * coefficient), nil}
		}(input)
	}

	var totalResult int64
	for i := 0; i < lineCount; i++ {
		res := <-resultChan
		if res.err != nil {
			return 0, res.err
		}
		totalResult += res.value
	}

	return totalResult, nil
}
