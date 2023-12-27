package main

type Coords struct {
	Row int
	Col int
}

type Well struct {
	Height int
	Width  int

	field     [][]bool
	piece     *Piece
	piecePos  Coords
	pieceOrnt Orientation
}

func NewWell(height int, width int) Well {
	field := make([][]bool, height)
	for row := 0; row < height; row++ {
		field[row] = make([]bool, width)
	}

	return Well{
		Height: height,
		Width:  width,

		field: field,
	}
}

func (w *Well) AddPiece(p *Piece) {
	w.piece = p
	w.piecePos = Coords{0, (w.Width - PieceSize) / 2}
	w.pieceOrnt = North
}

func (w *Well) IsOccupiedAt(pt Coords) bool {
	return w.isFilledAt(pt) || w.isPieceAt(pt)
}

func (w *Well) MoveRight() {
	if w.canMoveRight() {
		w.piecePos.Col += 1
	}
}

func (w *Well) MoveLeft() {
	if w.canMoveLeft() {
		w.piecePos.Col -= 1
	}
}

func (w *Well) RotateCW() {
	if w.canRotateCW() {
		w.pieceOrnt = w.pieceOrnt.RotateCW()
	}
}

func (w *Well) RotateCCW() {
	if w.canRotateCCW() {
		w.pieceOrnt = w.pieceOrnt.RotateCCW()
	}
}

func (w *Well) Drop() {
	for w.canFall() {
		w.Fall()
	}
}

func (w *Well) Fall() bool {
	if w.canFall() {
		w.piecePos.Row += 1
		return true
	} else {
		return false
	}
}

func (w *Well) BakeIn() {
	w.piece.IterateSolidParts(
		w.pieceOrnt,
		func (row int, col int) {
			w.field[w.piecePos.Row + row][w.piecePos.Col + col] = true
		},
	)

	w.piece = nil
}

func (w *Well) Snap() bool {
	for row := 0; row < w.Height; row++ {
		rowFull := true
		for col := 0; col < w.Width; col++ {
			rowFull = rowFull && w.field[row][col]
		}

		if rowFull {
			w.snapRow(row)
			return true
		}
	}
	return false
}

func (w *Well) snapRow(r int) {
	for row := r; row > 0; row-- {
		copy(w.field[row], w.field[row-1])
	}

	for col := 0; col < w.Width; col++ {
		w.field[0][col] = false
	}
}

func (w Well) isFilledAt(pt Coords) bool {
	return w.field[pt.Row][pt.Col]
}

func (w Well) isPieceAt(pt Coords) bool {
	if w.piece == nil {
		return false
	}

	if pt.Row < w.piecePos.Row || pt.Row - w.piecePos.Row >= PieceSize ||
           pt.Col < w.piecePos.Col || pt.Col - w.piecePos.Col >= PieceSize {

		return false
	}

	return w.piece.SolidAt(pt.Row - w.piecePos.Row, pt.Col - w.piecePos.Col, w.pieceOrnt)
}

func (w *Well) canMoveLeft() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Row, w.piecePos.Col - 1}, w.pieceOrnt)
}

func (w *Well) canMoveRight() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Row, w.piecePos.Col + 1}, w.pieceOrnt)
}

func (w *Well) canFall() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Row + 1, w.piecePos.Col}, w.pieceOrnt)
}

func (w *Well) canRotateCW() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Col, w.piecePos.Row}, w.pieceOrnt.RotateCW())
}

func (w *Well) canRotateCCW() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Col, w.piecePos.Row}, w.pieceOrnt.RotateCCW())
}

func (w *Well) WouldCollide(piece *Piece, pos Coords, ornt Orientation) bool {
	result := false

	piece.IterateSolidParts(
		ornt,
		func (row int, col int) {
			result = result || (pos.Row + row < 0 ||
             				    pos.Col + col < 0 ||
               				    pos.Row + row > w.Height - 1 ||
	               			    pos.Col + col > w.Width - 1)

			if result {
				return
			}

		        result = result || w.isFilledAt(Coords{pos.Row + row, pos.Col + col})
		},
	)

	return result
}
