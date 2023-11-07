package main

type Coords struct {
	Row int
	Col int
}

type Tetris struct {
	Height int
	Width  int

	field    [][]bool
	piece    *Piece
	piecePos Coords
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
	}
}

func (t *Tetris) AddPiece(p *Piece) {
	t.piece = p
	t.piecePos = Coords{2, 3}
}

func (t Tetris) IsOccupiedAt(pt Coords) bool {
	return t.isFilledAt(pt) || t.isPieceAt(pt)
}

func (t Tetris) isFilledAt(pt Coords) bool {
	return t.field[pt.Row][pt.Col]
}

func (t Tetris) isPieceAt(pt Coords) bool {
	if t.piece == nil {
		return false
	}

	if pt.Row < t.piecePos.Row || pt.Row - t.piecePos.Row >= PieceSize ||
	   pt.Col < t.piecePos.Col || pt.Col - t.piecePos.Col >= PieceSize {

		return false
	}

	return t.piece[pt.Row - t.piecePos.Row][pt.Col - t.piecePos.Col] == '■'
}
