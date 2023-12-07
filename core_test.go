package main

import "testing"

func TestEmptyField(t *testing.T) {
	ts := NewTetris(5, 6)

	assertSnapshot(t, ts, snapshot{
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

	ts := NewTetris(5, 5)
	ts.AddPiece(&piece)

	assertSnapshot(t, ts, snapshot{
		"|     |",
		"| xx  |",
		"| xx  |",
		"|     |",
		"|     |",
		"|-----|",
	})

	ts = NewTetris(5, 6)
	ts.AddPiece(&piece)

	assertSnapshot(t, ts, snapshot{
		"|      |",
		"|  xx  |",
		"|  xx  |",
		"|      |",
		"|      |",
		"|------|",
	})
}
