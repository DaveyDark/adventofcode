package utils

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func InGrid[T any](grid [][]T, pos [2]int) bool {
	return pos[1] >= 0 && pos[1] < len(grid) && pos[0] >= 0 && pos[0] < len(grid[0])
}

func ConstructGrid(inputFile string) ([][]rune, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	return constructGridFromReader(file)
}

func ContructGridFromStr(input string) ([][]rune, error) {
	reader := strings.NewReader(input)
	return constructGridFromReader(reader)
}

func constructGridFromReader(reader io.Reader) ([][]rune, error) {
	scanner := bufio.NewScanner(reader)

	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid, nil
}

func MapStrArrToInt(arr []string, start int, count int) ([]int, error) {
	intArr := make([]int, count)
	cnt := 0
	for i := start; i < start+count; i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			return intArr, err
		}
		intArr[cnt] = num
		cnt++
	}
	return intArr, nil
}

func GetDigits(inputFile string) ([]int, error) {
	digits := []int{}

	digitsRegex := regexp.MustCompile(`\d+`)
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return digits, err
	}
	inputStr := string(input)

	matches := digitsRegex.FindAllStringSubmatch(inputStr, -1)
	for _, match := range matches {
		matchInt, err := strconv.Atoi(match[0])
		if err != nil {
			return digits, err
		}
		digits = append(digits, matchInt)
	}

	return digits, nil
}
