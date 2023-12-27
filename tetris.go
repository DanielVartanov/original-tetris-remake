package main

import (
	"math/rand"
)

type Tetris struct {
	Well Well

	ticks int
}

func NewTetris() Tetris {
	well :=	NewWell(20, 10)

	tetris := Tetris{
		Well: well,
	}

	tetris.init()

	return tetris
}

func (ts *Tetris) Tick() {
	const speed int = 7

	ts.ticks += 1
	if ts.ticks == speed {
		ts.ticks = 0
		ts.step()
	}
}

func (ts *Tetris) MoveLeft() {
	ts.Well.MoveLeft()
}

func (ts *Tetris) MoveRight() {
	ts.Well.MoveRight()
}

func (ts *Tetris) RotateCW() {
	ts.Well.RotateCW()
}

func (ts *Tetris) Drop() {
	ts.Well.Drop()
}

func (ts *Tetris) init() {
	ts.addPiece()
}

func (ts *Tetris) addPiece() {
	randomPiece := Pieces[PieceNames[rand.Intn(len(PieceNames))]]
	ts.Well.AddPiece(&randomPiece)
}

func (ts *Tetris) step() {
	if ! ts.Well.Fall() {
		ts.endPiece()
	}
}

func (ts *Tetris) endPiece() {
	ts.Well.BakeIn()

	for ts.Well.Snap() {
		// Keep score here
	}

	ts.addPiece()
}
