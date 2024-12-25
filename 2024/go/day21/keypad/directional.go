package keypad

var DIRECTION = map[rune][2]int{
	'x': {0, 0},
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}
