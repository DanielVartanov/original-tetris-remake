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
	well *Well
	canvas viewport
}

func NewField(well *Well, screen *Screen) Field {
	// `+2` gives the padding for the field boundaries on the sides and at the bottom
	fldHeight := well.Height + 2
	fldWidth := (well.Width + 2) * len(space)

	canvas := screen.allocate(fldHeight, fldWidth, -fldHeight / 2, -fldWidth / 2)

	return Field{
		well: well,

		canvas: canvas,
	}
}

func (fld *Field) Render() {
	fld.renderBox()

	for trow := 0; trow < fld.well.Height; trow++ {
		for tcol := 0; tcol < fld.well.Width; tcol++ {
			var gl glyph
			if fld.well.IsOccupiedAt(Coords{trow, tcol}) {
				gl = occupied
			} else {
				gl = empty
			}

			fld.drawTile(gl, trow, tcol)
		}
	}
}

func (fld *Field) renderBox() {
	for trow := 0; trow < fld.well.Height; trow++ {
		fld.drawTile(leftBoundary, trow, -1)
		fld.drawTile(rightBoundary, trow, fld.well.Width)
	}

	fld.drawTile(leftBoundary, fld.well.Height, -1)
	fld.drawTile(space, fld.well.Height + 1, -1)
	for tcol := 0; tcol < fld.well.Width; tcol++ {
		fld.drawTile(bottom, fld.well.Height, tcol)
		fld.drawTile(foundation, fld.well.Height + 1, tcol)
	}
	fld.drawTile(rightBoundary, fld.well.Height, fld.well.Width)
	fld.drawTile(space, fld.well.Height + 1, fld.well.Width)
}

func (fld *Field) drawTile(gl glyph, trow int, tcol int) {
	fld.canvas.draw(gl, trow, (tcol + 1) * 2)
}
