package day13

import (
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day13/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
	// Parse input into a string
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}
	input := string(inputBytes)

	// RegEx to parse machines from input input
	machineRegex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

	tokens := int64(0)
	// Split string by newlines and iterate over them
	inputSplit := strings.Split(input, "\n\n")
	for _, machine := range inputSplit {
		matches := machineRegex.FindStringSubmatch(machine)
		button1, button2, prize, err := parseMachine(matches)
		if err != nil {
			return 0, err
		}
		prize[0] += 10000000000000
		prize[1] += 10000000000000
		machineTokens := solveTokensEquation(button1, button2, prize)
		tokens += machineTokens
	}
	return tokens, nil
}

func solveTokensEquation(button1 [2]int, button2 [2]int, prize [2]int) int64 {
	// Use a linear equation to solve for the number of tokens
	p1 := float64(prize[1]*button1[0]-prize[0]*button1[1]) / float64(button1[0]*button2[1]-button1[1]*button2[0])
	p0 := float64(prize[0]-button2[0]*int(p1)) / float64(button1[0])
	if p1 != math.Trunc(p1) || p0 != math.Trunc(p0) {
		return 0
	}
	return int64(3*p0 + p1)
}
