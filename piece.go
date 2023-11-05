package main

const PieceSize = 4

type piece [PieceSize][PieceSize]rune

var ThePiece = piece{
	{' ', ' ', ' ', ' '},
	{' ', '■', '■', ' '},
	{' ', '■', '■', ' '},
	{' ', ' ', ' ', ' '},
}
