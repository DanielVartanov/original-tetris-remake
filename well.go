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
}

func (w *Well) IsOccupiedAt(pt Coords) bool {
	return w.isFilledAt(pt) || w.isPieceAt(pt)
}

func (w *Well) MoveRight() {
	if w.CanMoveRight() {
		w.piecePos.Col += 1
	}
}

func (w *Well) MoveLeft() {
	if w.CanMoveLeft() {
		w.piecePos.Col -= 1
	}
}

func (w *Well) RotateCW() {
	if w.CanRotateCW() {
		w.pieceOrnt = w.pieceOrnt.RotateCW()
	}
}

func (w *Well) RotateCCW() {
	if w.CanRotateCCW() {
		w.pieceOrnt = w.pieceOrnt.RotateCCW()
	}
}

func (w *Well) Drop() {
	for w.CanFall() {
		w.Fall()
	}
}

func (w *Well) Fall() {
	if w.CanFall() {
		w.piecePos.Row += 1
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

func (w *Well) CanMoveLeft() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Row, w.piecePos.Col - 1}, w.pieceOrnt)
}

func (w *Well) CanMoveRight() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Row, w.piecePos.Col + 1}, w.pieceOrnt)
}

func (w *Well) CanFall() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Row + 1, w.piecePos.Col}, w.pieceOrnt)
}

func (w *Well) CanRotateCW() bool {
	return !w.WouldCollide(w.piece, Coords{w.piecePos.Col, w.piecePos.Row}, w.pieceOrnt.RotateCW())
}

func (w *Well) CanRotateCCW() bool {
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
		})

	return result
}
