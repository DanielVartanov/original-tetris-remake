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
