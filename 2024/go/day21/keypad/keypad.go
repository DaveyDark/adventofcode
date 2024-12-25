package keypad

import (
	"strings"
)

func MapString(input string, inputMap map[rune][2]int) string {
	res := strings.Builder{}
	pos := inputMap['A']
	for _, ch := range input {
		x := MoveFrom(pos, inputMap[ch], inputMap['x'])
		pos = inputMap[ch]
		res.WriteString(x)
	}
	return res.String()
}

func MoveFrom(src [2]int, dest [2]int, badTile [2]int) string {
	path := strings.Builder{}

	pos := src
	// If we are in bad column, move horizontally first
	// If we are in bad row, move vertically first
	xDir := 1
	if pos[1] > dest[1] {
		xDir = -1
	}
	yDir := 1
	if pos[0] > dest[0] {
		yDir = -1
	}
	if pos[0] == badTile[0] {
		// Move vertically first
		moveVertically(&pos, dest, yDir, &path)
		moveHorizontally(&pos, dest, xDir, &path)
	} else {
		// Move horizontally first
		moveHorizontally(&pos, dest, xDir, &path)
		moveVertically(&pos, dest, yDir, &path)
	}

	// Add button press
	path.WriteRune('A')

	return path.String()
}

// Helper functions
func moveVertically(pos *[2]int, dest [2]int, yDir int, path *strings.Builder) {
	for pos[0] != dest[0] {
		pos[0] += yDir
		if yDir == 1 {
			path.WriteRune('v')
		} else {
			path.WriteRune('^')
		}
	}
}

func moveHorizontally(pos *[2]int, dest [2]int, xDir int, path *strings.Builder) {
	for pos[1] != dest[1] {
		pos[1] += xDir
		if xDir == 1 {
			path.WriteRune('>')
		} else {
			path.WriteRune('<')
		}
	}
}
