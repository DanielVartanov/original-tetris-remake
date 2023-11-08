
package main

import (
	"testing"
	"strings"
)

func TestEmptyField(t *testing.T) {
	tetris := NewTetris(5, 6)

	assertSnapshot(t, tetris, snapshot{
		"|      |",
		"|      |",
		"|      |",
		"|      |",
		"|      |",
		"|------|",
	})
}

func TestTetris_AddPiece(t *testing.T) {
	piece := Pieces['O']

	tetris := NewTetris(5, 5)
	tetris.AddPiece(&piece)

	assertSnapshot(t, tetris, snapshot{
		"|     |",
		"| xx  |",
		"| xx  |",
		"|     |",
		"|     |",
		"|-----|",
	})

	tetris = NewTetris(5, 6)
	tetris.AddPiece(&piece)

	assertSnapshot(t, tetris, snapshot{
		"|      |",
		"|  xx  |",
		"|  xx  |",
		"|      |",
		"|      |",
		"|------|",
	})
}


// Helpers

type snapshot []string

func assertSnapshot(t *testing.T, ts Tetris, want snapshot) {
	t.Helper()

	got := takeSnapshot(ts)

	if strings.Join(got, "") != strings.Join(want, "") {
		errMsg := "\n\r\x1b[31mGot:\tWant:\n\r"

		for row := 0; row < ts.Height + 1; row++ {
			errMsg += got[row] + "\t" + want[row] + "\n\r"
		}

		errMsg += "\x1b[0m"

		t.Error(errMsg)
	} else {
		print("\n\r\x1b[32m" + strings.Join(got, "\n\r") + "\x1b[0m\n\r")
	}

}

func takeSnapshot(ts Tetris) snapshot {
	img := make([]string, ts.Height + 1)

	for row := 0; row < ts.Height; row++ {
		line := "|"

		for col := 0; col < ts.Width; col++ {
			if ts.IsOccupiedAt(Coords{row, col}) {
				line += "x"
			} else {
				line += " "
			}
		}

		line += "|"

		img[row] = line
	}

	img[ts.Height] = "|" + strings.Repeat("-", ts.Width) + "|"

	return img
}
