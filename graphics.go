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

var screen string

func clear() {
	screen = ""
}

func draw(gl glyph) {
	screen += string(gl)
}

func printLine(left glyph, central glyph, right glyph) {
	draw(left)
	for col := 1; col <= 10; col++ {
		draw(central)
	}
	draw(right)
	draw("\n\r")
}

func buildFrame(tetris Tetris) {
	for row := 0; row < tetris.Height; row++ {
		draw(leftBoundary)
		for col := 0; col < tetris.Width; col++ {
			if tetris.IsOccupiedAt(Coords{row, col}) {
				draw(occupied)
			} else {
				draw(empty)
			}

		}
		draw(rightBoundary)
		draw("\n\r")
	}
	printLine(leftBoundary, bottom, rightBoundary)
	printLine(space, foundation, space)
}

func drawTetris(tetris Tetris) {
	clear()
	buildFrame(tetris)
	print(screen)
}
