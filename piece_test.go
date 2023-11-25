package main

import (
	"testing"
	"reflect"
)

func TestPiece_IterateSolidParts(t *testing.T) {
	wantCallbacks := []Coords{{1, 1}, {2, 0}, {2, 1}, {2, 2}}
	piece := Pieces['T']

	gotCallbacks := []Coords{}

	piece.IterateSolidParts(func(row int, col int) {
		gotCallbacks = append(gotCallbacks, Coords{row, col})
	})

	if !reflect.DeepEqual(gotCallbacks, wantCallbacks) {
		t.Errorf("\n\rGot:  %v\n\rWant: %v", gotCallbacks, wantCallbacks)
	}
}

func TestPiece_SolidAt(t *testing.T) {
	piece := Pieces['S']

	if piece.SolidAt(1, 0) {
		t.Fail()
	}

	if !piece.SolidAt(1, 1) {
		t.Fail()
	}

	if !piece.SolidAt(2, 0) {
		t.Fail()
	}

	if !piece.SolidAt(2, 1) {
		t.Fail()
	}

	if piece.SolidAt(2, 2) {
		t.Fail()
	}
}
