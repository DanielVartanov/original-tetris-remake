package main

type tile string

const (
	empty tile = " ."
	leftBoundary tile = "<!"
	rightBoundary tile = "!>"
	bottom = "=="
	foundation tile = "\\/"
)

func main() {
	for row := 1; row <= 20; row++ {
		print(leftBoundary)
		for col := 1; col <= 10; col++ {
			print(empty)
		}
		print(rightBoundary)
		print("\n")
	}

	print(leftBoundary)
	for col := 1; col <= 10; col++ {
		print(bottom)
	}
	print(rightBoundary)
	print("\n")

	print(" ")
	for col := 1; col <= 10; col++ {
		print(foundation)
	}
	print(" ")
	print("\n")


	println("Hello, world")
}
