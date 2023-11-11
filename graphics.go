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

type viewport [][]rune

type Screen struct {
	Height int
	Width int

	tetris *Tetris

	main viewport
	field viewport
}

func NewScreen(height, width int, tetris *Tetris) Screen {
	main := make([][]rune, height)
	for row := range(main) {
		main[row] = make([]rune, width)
		for col := range(main[row]) {
			main[row][col] = ' '
		}
	}

	var field viewport
	fieldWidth := (tetris.Width + 2) * len(space)
	fieldHeight := tetris.Height + 2  // `+2` gives the padding for the field boundaries on the sides and at the bottom
	for row := height / 2 - fieldHeight / 2;
            row < height / 2 + fieldHeight / 2;
            row++ {
		    field = append(field, main[row][width / 2 - fieldWidth / 2 : width / 2 + fieldWidth / 2])
	    }

	return Screen{
		Height: height,
		Width: width,

		tetris: tetris,

		main: main,
		field: field,
	}
}

func (vp viewport) draw(gl glyph, row int, col int) {
	vp[row][col] = rune(gl[0])
	vp[row][col+1] = rune(gl[1])
}

func (scr *Screen) renderFieldTile(gl glyph, trow int, tcol int) {
	scr.field.draw(gl, trow, (tcol + 1) * 2)
}

func (scr *Screen) renderBox() {
	for trow := 0; trow < scr.tetris.Height; trow++ {
		scr.renderFieldTile(leftBoundary, trow, -1)
		scr.renderFieldTile(rightBoundary, trow, scr.tetris.Width)
	}

	scr.renderFieldTile(leftBoundary, scr.tetris.Height, -1)
	scr.renderFieldTile(space, scr.tetris.Height + 1, -1)
	for tcol := 0; tcol < scr.tetris.Width; tcol++ {
		scr.renderFieldTile(bottom, scr.tetris.Height, tcol)
		scr.renderFieldTile(foundation, scr.tetris.Height + 1, tcol)
	}
	scr.renderFieldTile(rightBoundary, scr.tetris.Height, scr.tetris.Width)
	scr.renderFieldTile(space, scr.tetris.Height + 1, scr.tetris.Width)
}

func (scr *Screen) renderField() {
	scr.renderBox()

	for trow := 0; trow < scr.tetris.Height; trow++ {
		for tcol := 0; tcol < scr.tetris.Width; tcol++ {
			var gl glyph
			if scr.tetris.IsOccupiedAt(Coords{trow, tcol}) {
				gl = occupied
			} else {
				gl = empty
			}

			scr.renderFieldTile(gl, trow, tcol)
		}
	}
}

func (scr *Screen) Render() {
	scr.renderField()
}

func (scr *Screen) Printable() string {
	frame := "\x1b[32m"
	for _, line := range(scr.main) {
		frame += string(line) + "\n\r"
	}
	frame += "\x1b[0m"

	return frame
}
