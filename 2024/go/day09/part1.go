package day09

import (
	"os"
	"reflect"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day9/part1"] = solve
}

func locateCursor(disk []int16, initial int) int {
	for i := initial; i < len(disk); i++ {
		if disk[i] == -1 {
			return i
		}
	}
	return len(disk) - 1
}

func solve(inputFile string) (int64, error) {
	// Read input
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}

	// Parse input into disk
	disk := []int16{}
	for i, ch := range string(input) {
		cnt := int(ch) - '0'
		if i%2 == 0 {
			for range cnt {
				disk = append(disk, int16(i/2))
			}
		} else {
			for range cnt {
				disk = append(disk, -1)
			}
		}
	}

	// Create cursor for disk
	cursor := locateCursor(disk, 0)

	swapDisk := reflect.Swapper(disk)
	for i := len(disk) - 1; i >= 0; i-- {
		if cursor >= i {
			break
		}
		if disk[i] != -1 {
			swapDisk(cursor, i)
			cursor = locateCursor(disk, cursor+1)
		}
	}

	checksum := int64(0)

	for i, n := range disk {
		if n == -1 {
			continue
		}
		checksum += int64(i) * int64(n)
	}

	return checksum, nil
}
