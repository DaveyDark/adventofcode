package day21

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/day21/keypad"
	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day21/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	res := int64(0)
	for scanner.Scan() {
		input := scanner.Text()

		coefficientStr := strings.TrimRight(input, "A")
		coefficient, err := strconv.Atoi(coefficientStr)
		if err != nil {
			return 0, err
		}

		// Parse input numbers to directions using numeric keypad
		presses := keypad.MapString(input, keypad.NUMERIC)
		presses = keypad.MapString(presses, keypad.DIRECTION)
		presses = keypad.MapString(presses, keypad.DIRECTION)

		res += int64(len(presses) * coefficient)
	}

	return res, nil
}
