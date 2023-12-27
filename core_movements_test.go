package main

import "testing"

func TestWell_MoveSideways(t *testing.T) {
	piece := Pieces['J']

	w := NewWell(4, 5)
	w.AddPiece(&piece)

	assertFilm(t, &w,
		actions{
			func() { w.MoveRight() },
		},
		film{
			{"|     |", "|     |"},
			{"| x   |", "|  x  |"},
			{"| xxx |", "|  xxx|"},
			{"|     |", "|     |"},
			{"|-----|", "|-----|"},
		},
	)

	assertFilm(t, &w,
		actions{
			func() { w.MoveLeft() },
			func() { w.MoveLeft() },
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

func TestWell_Fall(t *testing.T) {
	piece := Pieces['J']

	w := NewWell(6, 5)
	w.AddPiece(&piece)

	assertFilm(t, &w,
		actions{
			func() { w.Fall() },
			func() { w.Fall() },
			func() { w.Fall() },
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

func TestWell_Rotate(t *testing.T) {
	piece := Pieces['L']

	w := NewWell(4, 4)
	w.AddPiece(&piece)

	assertFilm(t, &w,
		actions{
			func() { w.RotateCW() },
			func() { w.RotateCW() },
			func() { w.RotateCW() },
			func() { w.RotateCW() },
		},
		film{
			{"|    |", "| x  |", "|    |", "|    |", "|    |"},
			{"|  x |", "| x  |", "| xxx|", "| xx |", "|  x |"},
			{"|xxx |", "| xx |", "| x  |", "|  x |", "|xxx |"},
			{"|    |", "|    |", "|    |", "|  x |", "|    |"},
			{"|----|", "|----|", "|----|", "|----|", "|----|"},
		},
	)
}

func TestWell_Drop(t *testing.T) {
	piece := Pieces['J']

	w := NewWell(5, 5)
	w.AddPiece(&piece)

	assertFilm(t, &w,
		actions{
			func() { w.Drop() },
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
