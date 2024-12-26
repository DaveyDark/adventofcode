package day22

import (
	"bufio"
	"os"
	"strconv"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day22/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	input, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(input)

	sum := int64(0)
	for scanner.Scan() {
		numStr := scanner.Text()
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			return 0, err
		}
		for range 2000 {
			num = advanceSecret(num)
		}
		sum += num
	}
	return sum, nil
}

func advanceSecret(num int64) int64 {
	_num := num * 64
	num = (num ^ _num) % 16777216
	_num = num / 32
	num = (num ^ _num) % 16777216
	_num = num * 2048
	num = (num ^ _num) % 16777216

	return num
}
