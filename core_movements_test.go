package main

import "testing"

func TestTetris_MoveSideways(t *testing.T) {
	piece := Pieces['J']

	ts := NewTetris(4, 5)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.MoveRight() },
		},
		film{
			{"|     |", "|     |"},
			{"| x   |", "|  x  |"},
			{"| xxx |", "|  xxx|"},
			{"|     |", "|     |"},
			{"|-----|", "|-----|"},
		},
	)

	assertFilm(t, &ts,
		actions{
			func() { ts.MoveLeft() },
			func() { ts.MoveLeft() },
		},
		film{
			{"|     |", "|     |", "|     |"},
			{"|  x  |", "| x   |", "|x    |"},
			{"|  xxx|", "| xxx |", "|xxx  |"},
			{"|     |", "|     |", "|     |"},
			{"|-----|", "|-----|", "|-----|"},
		},
	)
}

func TestTetris_Fall(t *testing.T) {
	piece := Pieces['J']

	ts := NewTetris(6, 5)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.Fall() },
			func() { ts.Fall() },
			func() { ts.Fall() },
		},
		film{
			{"|     |", "|     |", "|     |", "|     |"},
			{"| x   |", "|     |", "|     |", "|     |"},
			{"| xxx |", "| x   |", "|     |", "|     |"},
			{"|     |", "| xxx |", "| x   |", "|     |"},
			{"|     |", "|     |", "| xxx |", "| x   |"},
			{"|     |", "|     |", "|     |", "| xxx |"},
			{"|-----|", "|-----|", "|-----|", "|-----|"},
		},
	)
}

func TestTetris_Rotate(t *testing.T) {
	piece := Pieces['T']

	ts := NewTetris(4, 4)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.RotateCW() },
			func() { ts.RotateCW() },
			func() { ts.RotateCW() },
			func() { ts.RotateCW() },
		},
		film{
			{"|    |", "| x  |", "|    |", "|  x |", "|    |"},
			{"| x  |", "| xx |", "|xxx |", "| xx |", "| x  |"},
			{"|xxx |", "| x  |", "| x  |", "|  x |", "|xxx |"},
			{"|    |", "|    |", "|    |", "|    |", "|    |"},
			{"|----|", "|----|", "|----|", "|----|", "|----|"},
		},
	)
}


func TestTetris_Drop(t *testing.T) {
	piece := Pieces['J']

	ts := NewTetris(5, 5)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.Drop() },
		},
		film{
			{"|     |", "|     |"},
			{"| x   |", "|     |"},
			{"| xxx |", "|     |"},
			{"|     |", "| x   |"},
			{"|     |", "| xxx |"},
			{"|-----|", "|-----|"},
		},
	)
}
