package main

import 	"testing"

func TestFullCycle(t *testing.T) {
	scr := NewScreen(3, 4)

	for row := range(scr.Main) {
		for col := range(scr.Main[row]) {
			scr.Main[row][col] = '.'
		}
	}

	centre := scr.allocate(1, 2, 0, -1)
	centre.draw(occupied, 0, 0)

	want := "\x1b[32m" +
                "...." + "\n\r" +
		".[]." + "\n\r" +
		"...." + "\n\r" +
                "\x1b[0m"

	got := scr.Printable()

	if want != got {
		t.Errorf("\n\rGot:\n\r%s\n\rWant:\n\r%s", got, want)
	}
}
