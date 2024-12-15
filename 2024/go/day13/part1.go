package day13

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day13/part1"] = solve
}

func solve(inputFile string) (int64, error) {
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
		machineTokens := calculateTokens(button1, button2, prize)
		tokens += machineTokens
	}
	return tokens, nil
}

func calculateTokens(button1 [2]int, button2 [2]int, prize [2]int) int64 {
	presses := [2]int{0, 0}
	limits := calculateTokenLimits(button1, button2, prize)
	// Start a two pointer approach to calculate the number of tokens
	// Start with the maximum number of presses for button 1 and 0 presses for button 2
	presses[0] = limits[0]
	minTokens := 0
	sum := [2]int{}
	for presses[0] >= 0 && presses[1] <= limits[1] {
		sum[0] = presses[0]*button1[0] + presses[1]*button2[0]
		sum[1] = presses[0]*button1[1] + presses[1]*button2[1]
		if sum[0] == prize[0] && sum[1] == prize[1] {
			// If the sum is equal to the prize, calculate the number of tokens and update the minimum
			if minTokens == 0 || presses[0]*3+presses[1] < minTokens {
				minTokens = presses[0]*3 + presses[1]
			}
			presses[0]--
		} else if sum[0] < prize[0] || sum[1] < prize[1] {
			// If the sum is less than the prize, increment the presses for button 1
			presses[1]++
		} else {
			// If the sum is greater than the prize, decrement the presses for button 1
			presses[0]--
		}
	}
	return int64(minTokens)
}

func calculateTokenLimits(button1 [2]int, button2 [2]int, prize [2]int) [2]int {
	// Calculate the maximum number of times we can press button 1 while staying under prize value, capped at 100
	presses1 := prize[0] / button1[0]
	presses2 := prize[1] / button1[1]
	b1presses := presses1
	if presses2 < presses1 {
		b1presses = presses2
	}
	if b1presses > 100 {
		b1presses = 100
	}
	// Calculate the maximum number of times we can press button 2 while staying under prize value, capped at 100
	presses1 = prize[0] / button2[0]
	presses2 = prize[1] / button2[1]
	b2presses := presses1
	if presses2 < presses1 {
		b2presses = presses2
	}
	if b2presses > 100 {
		b2presses = 100
	}
	return [2]int{b1presses, b2presses}
}

func parseMachine(matches []string) ([2]int, [2]int, [2]int, error) {
	nums := make([]int, 6)

	for i := 0; i < 6; i++ {
		num, err := strconv.Atoi(matches[i+1])
		if err != nil {
			return [2]int{}, [2]int{}, [2]int{}, err
		}
		nums[i] = num
	}

	return [2]int{nums[0], nums[1]}, [2]int{nums[2], nums[3]}, [2]int{nums[4], nums[5]}, nil
}
