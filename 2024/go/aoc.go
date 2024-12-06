package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/daveydark/adventofcode/2024/day01"
	_ "github.com/daveydark/adventofcode/2024/day02"
	_ "github.com/daveydark/adventofcode/2024/day03"
	_ "github.com/daveydark/adventofcode/2024/day04"
	_ "github.com/daveydark/adventofcode/2024/day05"
	_ "github.com/daveydark/adventofcode/2024/day06"
	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func main() {
	// Parse command line arguments
	var day *int = flag.Int("d", 0, "Day number to run")
	var part *int = flag.Int("p", 0, "Part number to run")
	var input *string = flag.String("i", "", "Input file to use")

	flag.Parse()

	if *day == 0 || *part == 0 || *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Check if input file exists
	if _, err := os.Stat(*input); os.IsNotExist(err) {
		fmt.Printf("Input file %s does not exist\n", *input)
		os.Exit(1)
	}

	// Find the solver
	var registryKey = fmt.Sprintf("day%d/part%d", *day, *part)
	solver, found := registry.Registry[registryKey]
	if !found {
		fmt.Printf("Solution not found for Day %d Part %d\n", *day, *part)
		os.Exit(1)
	}

	answer, err := solver(*input)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Day: %d\n", *day)
	fmt.Printf("Part: %d\n", *part)
	fmt.Printf("Input: %s\n", *input)
	fmt.Printf("Answer: %d\n", answer)
}
