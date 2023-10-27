package main

type glyph string

const (
	empty glyph = " ."
	leftBoundary glyph = "<!"
	rightBoundary glyph = "!>"
	bottom = "=="
	foundation glyph = "\\/"
	space glyph = "  "
)

func printLine(left glyph, central glyph, right glyph) {
	print(left)
	for col := 1; col <= 10; col++ {
		print(central)
	}
	print(right)
	print("\n")
}

func main() {
	for row := 1; row <= 20; row++ {
		printLine(leftBoundary, empty, rightBoundary)
	}
	printLine(leftBoundary, bottom, rightBoundary)
	printLine(space, foundation, space)
}
