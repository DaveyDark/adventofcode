package day09

import (
	"os"
	"reflect"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

func init() {
	registry.Registry["day9/part2"] = solve2
}

func moveBlock(disk []int16, block [2]int, freeSpace [][2]int) {
	// Find best matching free space
	match := -1
	for i, space := range freeSpace {
		if space[0] > block[0] {
			continue
		}
		if space[1] >= block[1] && match == -1 {
			match = i
		}
	}

	if match == -1 {
		return
	}

	// Swap block with free space
	space := freeSpace[match]
	swapDisk := reflect.Swapper(disk)
	for i := 0; i < block[1]; i++ {
		swapDisk(block[0]+i, space[0]+i)
	}

	// Update free space
	if space[1] > block[1] {
		freeSpace[match] = [2]int{space[0] + block[1], space[1] - block[1]}
	} else {
		freeSpace[match] = freeSpace[len(freeSpace)-1]
		freeSpace = freeSpace[:len(freeSpace)-1]
	}
}

func solve2(inputFile string) (int64, error) {
	// Read input
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}

	// Parse input into disk
	disk := []int16{}
	blocks := [][2]int{} // Block ID: [start, length]
	freeSpace := [][2]int{}
	for i, ch := range string(input) {
		cnt := int(ch) - '0'
		if cnt == 0 {
			continue
		}
		if i%2 == 0 {
			blocks = append(blocks, [2]int{len(disk), cnt})
			for range cnt {
				disk = append(disk, int16(i/2))
			}
		} else {
			freeSpace = append(freeSpace, [2]int{len(disk), cnt})
			for range cnt {
				disk = append(disk, -1)
			}
		}
	}

	// Defragment disk
	for i := len(blocks) - 1; i >= 0; i-- {
		moveBlock(disk, blocks[i], freeSpace)
	}

	// Calculate checksum
	checksum := int64(0)
	for i, n := range disk {
		if n == -1 {
			continue
		}
		checksum += int64(i) * int64(n)
	}

	return checksum, nil
}
