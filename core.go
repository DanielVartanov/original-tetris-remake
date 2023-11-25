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

func (ts *Tetris) AddPiece(p *Piece) {
	ts.piece = p
	ts.piecePos = Coords{0, (ts.Width - PieceSize) / 2}
}

func (ts *Tetris) IsOccupiedAt(pt Coords) bool {
	return ts.isFilledAt(pt) || ts.isPieceAt(pt)
}

func (ts *Tetris) MoveRight() {
	if ts.CanMoveRight() {
		ts.piecePos.Col += 1
	}
}

func (ts *Tetris) MoveLeft() {
	if ts.CanMoveLeft() {
		ts.piecePos.Col -= 1
	}
}

func (ts *Tetris) Fall() {
	if ts.CanFall() {
		ts.piecePos.Row += 1
	}
}

func (ts Tetris) isFilledAt(pt Coords) bool {
	return ts.field[pt.Row][pt.Col]
}

func (ts Tetris) isPieceAt(pt Coords) bool {
	if ts.piece == nil {
		return false
	}

	if pt.Row < ts.piecePos.Row || pt.Row - ts.piecePos.Row >= PieceSize ||
	   pt.Col < ts.piecePos.Col || pt.Col - ts.piecePos.Col >= PieceSize {

		   return false
	   }

	return ts.piece.SolidAt(pt.Row - ts.piecePos.Row, pt.Col - ts.piecePos.Col)
}

func (ts *Tetris) CanMoveLeft() bool {
	return !ts.WouldCollide(ts.piece, Coords{ts.piecePos.Row, ts.piecePos.Col - 1})
}

func (ts *Tetris) CanMoveRight() bool {
	return !ts.WouldCollide(ts.piece, Coords{ts.piecePos.Row, ts.piecePos.Col + 1})
}

func (ts *Tetris) CanFall() bool {
	return !ts.WouldCollide(ts.piece, Coords{ts.piecePos.Row + 1, ts.piecePos.Col})
}

func (ts *Tetris) WouldCollide(piece *Piece, pos Coords) bool {
	result := false

	piece.IterateSolidParts(
		func (row int, col int) {
			result = result || (pos.Row + row < 0 ||
             				    pos.Col + col < 0 ||
               				    pos.Row + row > ts.Height - 1 ||
	               			    pos.Col + col > ts.Width - 1)

			if result {
				return
			}

		        result = result || ts.isFilledAt(Coords{pos.Row + row, pos.Col + col})
		})

	return result
}
