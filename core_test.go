package main

import (
	"testing"
	"strings"
)

func TestEmptyField(t *testing.T) {
	ts := NewTetris(5, 6)

	assertSnapshot(t, ts, snapshot{
		"|      |",
		"|      |",
		"|      |",
		"|      |",
		"|      |",
		"|------|",
	})
}

func TestTetris_AddPiece(t *testing.T) {
	piece := Pieces['O']

	ts := NewTetris(5, 5)
	ts.AddPiece(&piece)

	assertSnapshot(t, ts, snapshot{
		"|     |",
		"| xx  |",
		"| xx  |",
		"|     |",
		"|     |",
		"|-----|",
	})

	ts = NewTetris(5, 6)
	ts.AddPiece(&piece)

	assertSnapshot(t, ts, snapshot{
		"|      |",
		"|  xx  |",
		"|  xx  |",
		"|      |",
		"|      |",
		"|------|",
	})
}

func TestTetris_MoveSideways(t *testing.T) {
	piece := Pieces['J']

	ts := NewTetris(4, 5)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.MoveRight() },
		},
		film{
			{"|     |", "|     |"},
			{"| x   |", "|  x  |"},
			{"| xxx |", "|  xxx|"},
			{"|     |", "|     |"},
			{"|-----|", "|-----|"},
		},
	)

	assertFilm(t, &ts,
		actions{
			func() { ts.MoveLeft() },
			func() { ts.MoveLeft() },
		},
		film{
			{"|     |", "|     |", "|     |"},
			{"|  x  |", "| x   |", "|x    |"},
			{"|  xxx|", "| xxx |", "|xxx  |"},
			{"|     |", "|     |", "|     |"},
			{"|-----|", "|-----|", "|-----|"},
		},
	)
}

func TestTetris_Progress(t *testing.T) {
	piece := Pieces['J']

	ts := NewTetris(6, 5)
	ts.AddPiece(&piece)

	assertFilm(t, &ts,
		actions{
			func() { ts.Fall() },
			func() { ts.Fall() },
			func() { ts.Fall() },
		},
		film{
			{"|     |", "|     |", "|     |", "|     |"},
			{"| x   |", "|     |", "|     |", "|     |"},
			{"| xxx |", "| x   |", "|     |", "|     |"},
			{"|     |", "| xxx |", "| x   |", "|     |"},
			{"|     |", "|     |", "| xxx |", "| x   |"},
			{"|     |", "|     |", "|     |", "| xxx |"},
			{"|-----|", "|-----|", "|-----|", "|-----|"},
		},
	)
}


// Helpers

type snapshot []string

type actions []func()

type film [][]string

func assertSnapshot(t *testing.T, ts Tetris, want snapshot) {
	t.Helper()

	got := takeSnapshot(ts)

	if strings.Join(got, "") != strings.Join(want, "") {
		errMsg := "\n\r\x1b[31mGot:\tWant:\n\r"

		for row := 0; row < ts.Height + 1; row++ {
			errMsg += got[row] + "\t" + want[row] + "\n\r"
		}

		errMsg += "\x1b[0m"

		t.Error(errMsg)
	} else {
		print("\n\r\x1b[32m" + strings.Join(got, "\n\r") + "\x1b[0m\n\r")
	}

}

func takeSnapshot(ts Tetris) snapshot {
	img := make([]string, ts.Height + 1)

	for row := 0; row < ts.Height; row++ {
		line := "|"

		for col := 0; col < ts.Width; col++ {
			if ts.IsOccupiedAt(Coords{row, col}) {
				line += "x"
			} else {
				line += " "
			}
		}

		line += "|"

		img[row] = line
	}

	img[ts.Height] = "|" + strings.Repeat("-", ts.Width) + "|"

	return img
}

func assertFilm(t *testing.T, ts *Tetris, actions actions, wantFlm film) {
	t.Helper()

	var want, got []snapshot

	want = filmToSnapshots(wantFlm)

	got = append(got, takeSnapshot(*ts))
	for _, action := range(actions) {
		action()
		got = append(got, takeSnapshot(*ts))
	}

	if ! areSnapshotSeriesEqual(got, want) {
		t.Error(failedSnapshotSeriesErrMsg(got, want))
	} else {

		print(passSnapshotSeriesLogMsg(got))
	}
}

func failedSnapshotSeriesErrMsg(got []snapshot, want []snapshot) string {
	errMsg := "\n\r\x1b[31mGot:\n\r\t"

	numLines := len(got[0])
	for i := 0; i < numLines; i++ {
		for _, snapshot := range(got) {
			errMsg += snapshot[i] + "  "
		}
		errMsg += "\n\r\t"
	}

	errMsg += "\n\r\x1b[31mWant:\n\r\t"
	numLines = len(want[0])
	for i := 0; i < numLines; i++ {
		for _, snapshot := range(want) {
			errMsg += snapshot[i] + "  "
		}
		errMsg += "\n\r\t"
	}

        errMsg += "\x1b[0m\n\r"

	return errMsg
}

func passSnapshotSeriesLogMsg(got []snapshot) string {
	logMsg := "\n\r\x1b[32m"

	numLines := len(got[0])
	for i := 0; i < numLines; i++ {
		for _, snapshot := range(got) {
			logMsg += snapshot[i] + "  "
		}
		logMsg += "\n\r"
	}

	logMsg += "\x1b[0m\n\r"

	return logMsg
}

func areSnapshotSeriesEqual(got []snapshot, want []snapshot) bool {
	wantCombined := ""
	for _, img := range(want) {
		wantCombined += strings.Join(img, "")
	}

	gotCombined := ""
	for _, img := range(got) {
		gotCombined += strings.Join(img, "")
	}

	return wantCombined == gotCombined
}

func filmToSnapshots(flm film) []snapshot {
	numLines := len(flm)
	numSnapshots := len(flm[0])

	result := make([]snapshot, numSnapshots)

	for i := range(result) {
		result[i] = make(snapshot, numLines)
	}

	for li, lines := range(flm) {
		for si, line := range(lines) {
			result[si][li] = line
		}
	}

	return result
}
