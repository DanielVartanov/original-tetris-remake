package main

import "testing"

func TestFieldWithPiece(t *testing.T) {
	ts := NewTetris(4, 6)
	pc := Pieces['I']
	ts.AddPiece(&pc)

	scr := NewScreen(6, 16)

	fld := NewField(&ts, &scr)
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
