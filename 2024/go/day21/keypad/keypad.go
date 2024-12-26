package keypad

import (
	"strings"
)

func MapString(input string, inputMap map[rune][2]int) string {
	res := strings.Builder{}
	// Create memoization caches
	memo := make(map[[6]int]string)
	seqMemo := make(map[string]string)

	// Split into sequences ending with A
	sequences := strings.Split(input+"A", "A")
	sequences = sequences[:len(sequences)-1] // Remove empty string from split

	for _, seq := range sequences {
		// Check if sequence is cached
		if cached, ok := seqMemo[seq+"A"]; ok {
			res.WriteString(cached)
			continue
		}

		// Process new sequence
		seqRes := strings.Builder{}
		seqPos := inputMap['A']
		for _, ch := range seq + "A" {
			x := MoveFrom(seqPos, inputMap[ch], inputMap['x'], memo)
			seqPos = inputMap[ch]
			seqRes.WriteString(x)
		}

		// Cache and append result
		seqStr := seqRes.String()
		seqMemo[seq+"A"] = seqStr
		res.WriteString(seqStr)
	}

	return res.String()
}

func MoveFrom(src [2]int, dest [2]int, badTile [2]int, memo map[[6]int]string) string {
	// Create cache key from src, dest and badTile coordinates
	key := [6]int{src[0], src[1], dest[0], dest[1], badTile[0], badTile[1]}

	// Check if result is in cache
	if cached, ok := memo[key]; ok {
		return cached
	}

	path := strings.Builder{}

	pos := src
	move(&pos, dest, badTile, &path)
	if pos != dest {
		move(&pos, dest, badTile, &path)
	}

	// Add button press
	path.WriteRune('A')

	result := path.String()
	// Store in cache
	memo[key] = result
	return result
}

func move(pos *[2]int, dest [2]int, badTile [2]int, path *strings.Builder) {
	// Preference of moves: <, v, ^/>, A
	// Check if we can move left
	if pos[1] > dest[1] {
		// Projected position if we move left
		leftPos := [2]int{pos[0], dest[1]}
		// If the projected position is not a bad tile
		if leftPos != badTile {
			for i := 0; i < pos[1]-dest[1]; i++ {
				path.WriteRune('<')
			}
			*pos = leftPos
		}
	}
	// Check if we can move down
	if pos[0] < dest[0] {
		// Projected position if we move down
		downPos := [2]int{dest[0], pos[1]}
		// If the projected position is not a bad tile
		if downPos != badTile {
			for i := 0; i < dest[0]-pos[0]; i++ {
				path.WriteRune('v')
			}
			*pos = downPos
		}
	}
	// Check if we can move up
	if pos[0] > dest[0] {
		// Projected position if we move up
		upPos := [2]int{dest[0], pos[1]}
		// If the projected position is not a bad tile
		if upPos != badTile {
			for i := 0; i < pos[0]-dest[0]; i++ {
				path.WriteRune('^')
			}
			*pos = upPos
		}
	}
	// Check if we can move right
	if pos[1] < dest[1] {
		// Projected position if we move right
		rightPos := [2]int{pos[0], dest[1]}
		// If the projected position is not a bad tile
		if rightPos != badTile {
			for i := 0; i < dest[1]-pos[1]; i++ {
				path.WriteRune('>')
			}
			*pos = rightPos
		}
	}
}
