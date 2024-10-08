package main

import "testing"

func TestEmptyFieldWithPiece(t *testing.T) {
	well := NewWell(4, 6)
	pc := Pieces['I']
	well.AddPiece(&pc)

	scr := NewScreen(6, 16)

	fld := NewField(&well, &scr)
	fld.Render()

	want := "\x1b[32m" +
		"<! . . . . . .!>" + "\n\r" +
		"<! .[][][][] .!>" + "\n\r" +
		"<! . . . . . .!>" + "\n\r" +
		"<! . . . . . .!>" + "\n\r" +
		"<!============!>" + "\n\r" +
		`  \/\/\/\/\/\/  ` + "\n\r" +
		"\x1b[0m"

	got := scr.Printable()
	if got != want {
		t.Errorf("\n\rGot:\n\r%s\n\rWant:\n\r%s", got, want)
	}
}

func TestFilledFieldWithPiece(t* testing.T) {
	well := NewWell(5, 6)
	pc := Pieces['L']
	well.AddPiece(&pc)

	filled := []Coords{{4, 0}, {4, 1}, {3, 0}, {3, 1}, {4, 4}, {4, 5}}
	for _, coords := range(filled) {
		well.field[coords.Row][coords.Col] = true
	}

	scr := NewScreen(7, 16)

	fld := NewField(&well, &scr)
	fld.Render()

	want := "\x1b[32m" +
        	"<! . . . . . .!>" + "\n\r" +
		"<! . . .[] . .!>" + "\n\r" +
		"<! .[][][] . .!>" + "\n\r" +
		"<![][] . . . .!>" + "\n\r" +
	        "<![][] . .[][]!>" + "\n\r" +
		"<!============!>" + "\n\r" +
		`  \/\/\/\/\/\/  ` + "\n\r" +
		"\x1b[0m"

	got := scr.Printable()
	if got != want {
		t.Errorf("\n\rGot:\n\r%s\n\rWant:\n\r%s", got, want)
	}
}
