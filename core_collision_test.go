package main

import "testing"

func TestTetris_MoveSideways_Boundary_Collision(t *testing.T) {
	piece := Pieces['O']

	ts := NewTetris(4, 4)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.MoveRight() },
			func() { ts.MoveRight() },
			func() { ts.MoveRight() },
		},
		film{
			{"|    |", "|    |", "|    |", "|    |"},
			{"| xx |", "|  xx|", "|  xx|", "|  xx|"},
			{"| xx |", "|  xx|", "|  xx|", "|  xx|"},
			{"|    |", "|    |", "|    |", "|    |"},
			{"|----|", "|----|", "|----|", "|----|"},
		},
	)

	ts = NewTetris(4, 4)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.MoveLeft() },
			func() { ts.MoveLeft() },
			func() { ts.MoveLeft() },
		},
		film{
			{"|    |", "|    |", "|    |", "|    |"},
			{"| xx |", "|xx  |", "|xx  |", "|xx  |"},
			{"| xx |", "|xx  |", "|xx  |", "|xx  |"},
			{"|    |", "|    |", "|    |", "|    |"},
			{"|----|", "|----|", "|----|", "|----|"},
		},
	)
}

func TestTetris_Fall_Boundary_Collision(t *testing.T) {
	piece := Pieces['J']

	ts := NewTetris(5, 5)
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
			{"|     |", "| xxx |", "| x   |", "| x   |"},
			{"|     |", "|     |", "| xxx |", "| xxx |"},
			{"|-----|", "|-----|", "|-----|", "|-----|"},
		},
	)
}

func TestTetris_Drop_Boundary_Collision(t *testing.T) {
	piece := Pieces['J']

	ts := NewTetris(5, 5)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.Drop() },
			func() { ts.Drop() },
		},
		film{
			{"|     |", "|     |", "|     |"},
			{"| x   |", "|     |", "|     |"},
			{"| xxx |", "|     |", "|     |"},
			{"|     |", "| x   |", "| x   |"},
			{"|     |", "| xxx |", "| xxx |"},
			{"|-----|", "|-----|", "|-----|"},
		},
	)
}

func TestTetris_Rotate_Boundary_Collision(t *testing.T) {
	piece := Pieces['I']

	ts := NewTetris(4, 5)
	ts.AddPiece(&piece)

	ts.RotateCW()
	ts.MoveRight()
	ts.MoveRight()

	assertFilm(t, &ts,
		actions{
			func() { ts.RotateCW() },
		},
		film{
			{"|    x|", "|    x|"},
			{"|    x|", "|    x|"},
			{"|    x|", "|    x|"},
			{"|    x|", "|    x|"},
			{"|-----|", "|-----|"},
		},
	)
}
