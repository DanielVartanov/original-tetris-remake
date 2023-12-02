package main

const PieceSize = 4

type Piece [PieceSize][PieceSize]rune

type Orientation int

const (
	North Orientation = iota
	East
	South
	West
)

var PieceNames = []rune{'O', 'S', 'Z', 'T', 'L', 'J', 'I'}

var Pieces = map[rune]Piece{
	'O': {
		{' ', ' ', ' ', ' '},
		{' ', '■', '■', ' '},
		{' ', '■', '■', ' '},
		{' ', ' ', ' ', ' '},
	},
	'S': {
		{' ', ' ', ' ', ' '},
		{' ', '■', '■', ' '},
		{'■', '■', ' ', ' '},
		{' ', ' ', ' ', ' '},
	},
	'Z': {
		{' ', ' ', ' ', ' '},
		{'■', '■', ' ', ' '},
		{' ', '■', '■', ' '},
		{' ', ' ', ' ', ' '},
	},
	'T': {
		{' ', ' ', ' ', ' '},
		{' ', '■', ' ', ' '},
		{'■', '■', '■', ' '},
		{' ', ' ', ' ', ' '},
	},
	'L': {
		{' ', ' ', ' ', ' '},
		{' ', ' ', '■', ' '},
		{'■', '■', '■', ' '},
		{' ', ' ', ' ', ' '},
	},
	'J': {
		{' ', ' ', ' ', ' '},
		{' ', '■', ' ', ' '},
		{' ', '■', '■', '■'},
		{' ', ' ', ' ', ' '},
	},
	'I': {
		{' ', ' ', ' ', ' '},
		{'■', '■', '■', '■'},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
	},
}

func (pc Piece) IterateSolidParts(ornt Orientation, fn func(row int, col int)) {
	for row := range(pc) {
		for col := range(pc[row]) {
			if (pc.SolidAt(row, col, ornt)) {
				fn(row, col)
			}
		}
	}
}

func (pc Piece) SolidAt(row int, col int, ornt Orientation) bool {
	switch ornt {
	case East:
		buf := col
		col = row
		row = buf
		row = PieceSize - 1 - row
	case South:
		row = PieceSize - 1 - row
	case West:
		buf := col
		col = row
		row = buf
	}

	return pc[row][col] == '■'
}

func (ornt Orientation) RotateCW() Orientation {
	return (ornt + 1) % 4
}

func (ornt Orientation) RotateCCW() Orientation {
	return (ornt - 1 + 4) % 4
}
