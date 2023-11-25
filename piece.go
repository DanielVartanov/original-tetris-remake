package main

const PieceSize = 4

type Piece [PieceSize][PieceSize]rune

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

func (pc Piece) IterateSolidParts(fn func(row int, col int)) {
	for row := range(pc) {
		for col := range(pc[row]) {
			if (pc.SolidAt(row, col)) {
				fn(row, col)
			}
		}
	}
}

func (pc Piece) SolidAt(row int, col int) bool {
	return pc[row][col] == '■'
}
