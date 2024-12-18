package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
)

func init() {
	registry.Registry["day17/part1"] = solve
}

func solve(inputFile string) (int64, error) {
	input, err := utils.GetDigits(inputFile)
	if err != nil {
		return 0, err
	}
	computer := NewComputer(input[0], input[1], input[2])
	program := parseProgram(input[3:])

	instPtr := 0
	output := []int{}
	for instPtr < len(program) {
		res := computer.execute(program[instPtr][0], program[instPtr][1], &output)
		if res == -1 {
			instPtr++
		} else {
			instPtr = res
		}
	}

	resStr := ""
	for _, o := range output {
		resStr += strconv.Itoa(o) + ","
	}
	resStr = strings.TrimRight(resStr, ",")

	// Exception: Since this problem outputs a string as a result, we cannot return it since the solve() type
	// defines int64 as the successful return type, so we just print resStr and return 0
	fmt.Println(resStr)

	return 1, nil
}

func parseProgram(digits []int) [][2]int {
	program := [][2]int{}
	for i := 0; i < len(digits)/2; i++ {
		program = append(program, [2]int{digits[i*2], digits[i*2+1]})
	}
	return program
}
