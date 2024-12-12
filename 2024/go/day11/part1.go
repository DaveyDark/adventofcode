package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day11/part1"] = solve
}

func noOfDigits(n uint64) int {
	cnt := 0
	for n > 0 {
		cnt++
		n /= 10
	}
	return cnt
}

func processStone(stone uint64) []uint64 {
	res := []uint64{}
	if stone == 0 {
		res = append(res, 1)
	} else if noOfDigits(stone)%2 == 1 {
		res = append(res, stone*2024)
	} else {
		stoneStr := strconv.FormatUint(stone, 10)
		leftStr := stoneStr[:len(stoneStr)/2]
		rightStr := stoneStr[len(stoneStr)/2:]
		left, err := strconv.ParseUint(leftStr, 10, 64)
		if err != nil {
			return []uint64{}
		}
		right, err := strconv.ParseUint(rightStr, 10, 64)
		if err != nil {
			return []uint64{}
		}
		res = append(res, left, right)
	}
	return res
}
func blink(stones []uint64) []uint64 {
	next := []uint64{}

	for _, stone := range stones {
		next = append(next, processStone(stone)...)
	}

	return next
}

func parseInput(inputFile string) ([]uint64, error) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return []uint64{}, err
	}
	inputStr := string(input)

	stones := []uint64{}
	for _, field := range strings.Fields(inputStr) {
		stone, err := strconv.ParseUint(field, 10, 64)
		if err != nil {
			return []uint64{}, err
		}
		stones = append(stones, stone)
	}
	return stones, nil
}

func solve(inputFile string) (int64, error) {
	stones, err := parseInput(inputFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	for range 25 {
		stones = blink(stones)
	}

	return int64(len(stones)), nil
}
