package main

import (
	"testing"
	"reflect"
)

func TestPiece_IterateSolidParts(t *testing.T) {
	wantCallbacks := []Coords{{1, 1}, {2, 0}, {2, 1}, {2, 2}}
	piece := Pieces['T']

	gotCallbacks := []Coords{}

	piece.IterateSolidParts(North, func(row int, col int) {
		gotCallbacks = append(gotCallbacks, Coords{row, col})
	})

	if !reflect.DeepEqual(gotCallbacks, wantCallbacks) {
		t.Errorf("\n\rGot:  %v\n\rWant: %v", gotCallbacks, wantCallbacks)
	}
}

func TestPiece_IterateSolidParts_Rotated(t *testing.T) {
	wantCallbacks := []Coords{{0, 1}, {1, 1}, {1, 2}, {2, 1}}
	piece := Pieces['T']

	gotCallbacks := []Coords{}

	piece.IterateSolidParts(East, func(row int, col int) {
		gotCallbacks = append(gotCallbacks, Coords{row, col})
	})

	if !reflect.DeepEqual(gotCallbacks, wantCallbacks) {
		t.Errorf("\n\rGot:  %v\n\rWant: %v", gotCallbacks, wantCallbacks)
	}
}

func TestPiece_SolidAt(t *testing.T) {
	piece := Pieces['S']

	if piece.SolidAt(1, 0, North) {
		t.Fail()
	}

	if !piece.SolidAt(1, 1, North) {
		t.Fail()
	}

	if !piece.SolidAt(2, 0, North) {
		t.Fail()
	}

	if !piece.SolidAt(2, 1, North) {
		t.Fail()
	}

	if piece.SolidAt(2, 2, North) {
		t.Fail()
	}
}

func TestPiece_SolidAt_Rotated(t *testing.T) {
	piece := Pieces['T']

	if piece.SolidAt(2, 0, East) {
		t.Fail()
	}

	if !piece.SolidAt(0, 1, East) {
		t.Fail()
	}

	if !piece.SolidAt(1, 1, East) {
		t.Fail()
	}

	if !piece.SolidAt(1, 2, East) {
		t.Fail()
	}

	if piece.SolidAt(2, 2, East) {
		t.Fail()
	}
}

func TestOrientation_Rotate(t *testing.T) {
	var ornt Orientation = East

	ornt = ornt.RotateCW()
	if ornt != South {
		t.Errorf("Got: %v, Want: %v", ornt, South)
	}

	ornt = ornt.RotateCW()
	if ornt != West {
		t.Errorf("Got: %v, Want: %v", ornt, West)
	}

	ornt = ornt.RotateCW()
	if ornt != North {
		t.Errorf("Got: %v, Want: %v", ornt, North)
	}

	ornt = ornt.RotateCCW()
	if ornt != West {
		t.Errorf("Got: %v, Want: %v", ornt, West)
	}

	ornt = ornt.RotateCCW()
	if ornt != South {
		t.Errorf("Got: %v, Want: %v", ornt, South)
	}
}
