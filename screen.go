package main

type Screen struct {
	Height int
	Width int

	main viewport
}

func NewScreen(height, width int, tetris *Tetris) Screen {
	main := make([][]rune, height)
	for row := range(main) {
		main[row] = make([]rune, width)
		for col := range(main[row]) {
			main[row][col] = ' '
		}
	}

	return Screen{
		Height: height,
		Width: width,

		main: main,
	}
}

func (scr *Screen) Printable() string {
	frame := "\x1b[32m"
	for _, line := range(scr.main) {
		frame += string(line) + "\n\r"
	}
	frame += "\x1b[0m"

	return frame
}

func (scr *Screen) allocate(height int, width int, rowOffset int, colOffset int) viewport {
	var vp viewport

	for row := scr.Height / 2 + rowOffset;
            row < scr.Height / 2 + height + rowOffset;
            row++ {
		    vp = append(vp, scr.main[row][scr.Width / 2 + colOffset : scr.Width / 2 + width + colOffset])
	    }

	return vp
}


type viewport [][]rune

func (vp viewport) draw(gl glyph, row int, col int) {
	vp[row][col] = rune(gl[0])
	vp[row][col+1] = rune(gl[1])
}
