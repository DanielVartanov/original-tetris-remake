package main

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

type glyph string

const (
	empty         glyph = " ."
	leftBoundary  glyph = "<!"
	rightBoundary glyph = "!>"
	bottom              = "=="
	foundation    glyph = "\\/"
	space         glyph = "  "
)

func printLine(left glyph, central glyph, right glyph) {
	print(left)
	for col := 1; col <= 10; col++ {
		print(central)
	}
	print(right)
	print("\n\r")
}

func printPlayingField() {
	for row := 1; row <= 20; row++ {
		printLine(leftBoundary, empty, rightBoundary)
	}
	printLine(leftBoundary, bottom, rightBoundary)
	printLine(space, foundation, space)
}

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

func clearScreen() {
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

	clearScreen()

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

	clearScreen()
}
