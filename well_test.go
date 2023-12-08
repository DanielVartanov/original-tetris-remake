package main

import "testing"

func TestWell_Empty(t *testing.T) {
	w := NewWell(5, 6)

	assertSnapshot(t, w, snapshot{
		"|      |",
		"|      |",
		"|      |",
		"|      |",
		"|      |",
		"|------|",
	})
}

func TestWell_AddPiece(t *testing.T) {
	piece := Pieces['O']

	w := NewWell(5, 5)
	w.AddPiece(&piece)

	assertSnapshot(t, w, snapshot{
		"|     |",
		"| xx  |",
		"| xx  |",
		"|     |",
		"|     |",
		"|-----|",
	})

	w = NewWell(5, 6)
	w.AddPiece(&piece)

	assertSnapshot(t, w, snapshot{
		"|      |",
		"|  xx  |",
		"|  xx  |",
		"|      |",
		"|      |",
		"|------|",
	})
}
