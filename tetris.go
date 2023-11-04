package main

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)


// Game

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


// Graphics

type glyph string

const (
	empty         glyph = " ."
	occupied      glyph = "[]"
	leftBoundary  glyph = "<!"
	rightBoundary glyph = "!>"
	bottom        glyph = "=="
	foundation    glyph = "\\/"
	space         glyph = "  "
)

var screen string

func clear() {
	screen = ""
}

func draw(gl glyph) {
	screen += string(gl)
}

func printLine(left glyph, central glyph, right glyph) {
	draw(left)
	for col := 1; col <= 10; col++ {
		draw(central)
	}
	draw(right)
	draw("\n\r")
}

func buildFrame() {
	for row := 0; row < 20; row++ {
		draw(leftBoundary)
		for col := 0; col < 10; col++ {
			if isOccupiedAt(coords{row, col}) {
				draw(occupied)
			} else {
				draw(empty)
			}

		}
		draw(rightBoundary)
		draw("\n\r")
	}
	printLine(leftBoundary, bottom, rightBoundary)
	printLine(space, foundation, space)
}

func printPlayingField() {
	clear()
	buildFrame()
	print(screen)
}


// Initialisation

func hideCursor() {
	print("\x1b[?25l")
}

func showCursor() {
	print("\x1b[?25h")
}

func ttyfd() int {
	return int(os.Stdin.Fd())
}

func enableRawMode() *term.State {
	saved, err := term.MakeRaw(ttyfd())
	if err != nil {
		println("Error enabling terminal raw mode")
		os.Exit(1)
	}
	return saved
}

func disableRawMode(oldstate *term.State) {
	err := term.Restore(ttyfd(), oldstate)
	if err != nil {
		println("Error disabling terminal raw mode")
		os.Exit(1)
	}
}

func clearTerminal() {
	print("\x1b[2J")
}

func resetCursor() {
	print("\x1b[H")
}

func keystrokes() <-chan byte {
	keys := make(chan byte)

	go func() {
		buf := make([]byte, 1)
		for {
			n, err := syscall.Read(int(os.Stdin.Fd()), buf)
			if err != nil {
				print("Error reading from stdin:", err)
				os.Exit(1)
			}

			if n > 0 {
				keys <- buf[0]
			}
		}
	}()

	return keys
}

func main() {
	sig := make (chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	saved := enableRawMode()
	defer disableRawMode(saved)

	hideCursor()
	defer showCursor()

	keys := keystrokes()

	clearTerminal()

	initGame()

	mainloop:
		for {
			select {
			case <-sig:
				break mainloop
			case key := <-keys:
				switch key{
				case 0x03: // Ctrl+C
					break mainloop
				case 'q':
					break mainloop
				}
			default:

			}

			resetCursor()
			printPlayingField()
		}

	clearTerminal()
}
