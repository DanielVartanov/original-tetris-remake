package main

type glyph string

const (
	empty         glyph = " ."
	occupied      glyph = "[]"
	leftBoundary  glyph = "<!"
	rightBoundary glyph = "!>"
	bottom        glyph = "=="
	foundation    glyph = "\\/"
	space         glyph = "  "
)

type Field struct {
	tetris *Tetris
	canvas viewport
}

func NewField(tetris *Tetris, screen *Screen) Field {
	// `+2` gives the padding for the field boundaries on the sides and at the bottom
	fldHeight := tetris.Height + 2
	fldWidth := (tetris.Width + 2) * len(space)

	canvas := screen.allocate(fldHeight, fldWidth, -fldHeight / 2, -fldWidth / 2)

	return Field{
		tetris: tetris,

		canvas: canvas,
	}
}

func (fld *Field) Render() {
	fld.renderBox()

	for trow := 0; trow < fld.tetris.Height; trow++ {
		for tcol := 0; tcol < fld.tetris.Width; tcol++ {
			var gl glyph
			if fld.tetris.IsOccupiedAt(Coords{trow, tcol}) {
				gl = occupied
			} else {
				gl = empty
			}

			fld.drawTile(gl, trow, tcol)
		}
	}
}

func (fld *Field) renderBox() {
	for trow := 0; trow < fld.tetris.Height; trow++ {
		fld.drawTile(leftBoundary, trow, -1)
		fld.drawTile(rightBoundary, trow, fld.tetris.Width)
	}

	fld.drawTile(leftBoundary, fld.tetris.Height, -1)
	fld.drawTile(space, fld.tetris.Height + 1, -1)
	for tcol := 0; tcol < fld.tetris.Width; tcol++ {
		fld.drawTile(bottom, fld.tetris.Height, tcol)
		fld.drawTile(foundation, fld.tetris.Height + 1, tcol)
	}
	fld.drawTile(rightBoundary, fld.tetris.Height, fld.tetris.Width)
	fld.drawTile(space, fld.tetris.Height + 1, fld.tetris.Width)
}

func (fld *Field) drawTile(gl glyph, trow int, tcol int) {
	fld.canvas.draw(gl, trow, (tcol + 1) * 2)
}
