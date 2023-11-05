package main

type Coords struct {
	Row int
	Col int
}

type Tetris struct {
	Height int
	Width  int

	field    [][]bool
	pieceLoc Coords
}

func NewTetris(height int, width int) Tetris {
	field := make([][]bool, height)
	for row := 0; row < height; row++ {
		field[row] = make([]bool, width)
	}

	return Tetris{
		Height: height,
		Width: width,

		field: field,
		pieceLoc: Coords{2, 3},
	}
}

func (t Tetris) IsOccupiedAt(pt Coords) bool {
	return t.isFilledAt(pt) || t.isPieceAt(pt)
}


func (t Tetris) isFilledAt(pt Coords) bool {
	return t.field[pt.Row][pt.Col]
}

func (t Tetris) isPieceAt(pt Coords) bool {
	if pt.Row < t.pieceLoc.Row || pt.Row - t.pieceLoc.Row >= PieceSize ||
	   pt.Col < t.pieceLoc.Col || pt.Col - t.pieceLoc.Col >= PieceSize {

		return false
	}

	return ThePiece[pt.Row - t.pieceLoc.Row][pt.Col - t.pieceLoc.Col] == '■'
}
