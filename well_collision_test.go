package main

import "testing"

func TestWell_MoveSideways_Boundary_Collision(t *testing.T) {
	piece := Pieces['O']

	w := NewWell(4, 4)
	w.AddPiece(&piece)

	assertFilm(t, &w,
		actions{
			func() { w.MoveRight() },
			func() { w.MoveRight() },
			func() { w.MoveRight() },
		},
		film{
			{"|    |", "|    |", "|    |", "|    |"},
			{"| xx |", "|  xx|", "|  xx|", "|  xx|"},
			{"| xx |", "|  xx|", "|  xx|", "|  xx|"},
			{"|    |", "|    |", "|    |", "|    |"},
			{"|----|", "|----|", "|----|", "|----|"},
		},
	)

	w = NewWell(4, 4)
	w.AddPiece(&piece)

	assertFilm(t, &w,
		actions{
			func() { w.MoveLeft() },
			func() { w.MoveLeft() },
			func() { w.MoveLeft() },
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

func TestWell_Fall_Boundary_Collision(t *testing.T) {
	piece := Pieces['J']

	w := NewWell(5, 5)
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
			{"|     |", "| xxx |", "| x   |", "| x   |"},
			{"|     |", "|     |", "| xxx |", "| xxx |"},
			{"|-----|", "|-----|", "|-----|", "|-----|"},
		},
	)
}

func TestWell_Drop_Boundary_Collision(t *testing.T) {
	piece := Pieces['J']

	w := NewWell(5, 5)
	w.AddPiece(&piece)

	assertFilm(t, &w,
		actions{
			func() { w.Drop() },
			func() { w.Drop() },
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

func TestWell_Rotate_Boundary_Collision(t *testing.T) {
	piece := Pieces['I']

	w := NewWell(4, 5)
	w.AddPiece(&piece)

	w.RotateCW()
	w.MoveRight()
	w.MoveRight()

	assertFilm(t, &w,
		actions{
			func() { w.RotateCW() },
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
