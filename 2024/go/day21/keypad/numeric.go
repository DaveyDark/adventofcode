package keypad

var NUMERIC = map[rune][2]int{
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'x': {3, 0},
	'0': {3, 1},
	'A': {3, 2},
}