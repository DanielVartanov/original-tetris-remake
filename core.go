package main

type piece [4][4]rune

var thePiece = piece{
	{' ',' ',' ',' '},
	{' ','■','■',' '},
	{' ','■','■',' '},
	{' ',' ',' ',' '},
}

var field [20][10]bool

type coords struct {
	row int
	col int
}

var pieceLoc coords

func isFilledAt(pt coords) bool {
	return field[pt.row][pt.col]
}

func isPieceAt(pt coords) bool {
	if (pt.row < pieceLoc.row || pt.row - pieceLoc.row >=4 ||
            pt.col < pieceLoc.col || pt.col - pieceLoc.col >=4) {

		    return false
	    }

	return thePiece[pt.row - pieceLoc.row][pt.col - pieceLoc.col] == '■'
}

func isOccupiedAt(pt coords) bool {
	return isFilledAt(pt) || isPieceAt(pt)
}

func initGame() {
	pieceLoc = coords{2, 3}
}
