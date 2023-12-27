package main

import (
	"testing"
	"strings"
)

type snapshot []string

type actions []func()

type film [][]string

func fillWell(w Well, filling snapshot) {
	for row := 0; row <= len(filling) - 2; row++ {
		for col := 1; col <= len(filling[row]) - 2; col++ {
			if filling[row][col] == 'x' {
				w.field[row][col - 1] = true
			}
		}
	}
}

func assertSnapshot(t *testing.T, well Well, want snapshot) {
	t.Helper()

	got := takeSnapshot(well)

	if strings.Join(got, "") != strings.Join(want, "") {
		errMsg := "\n\r\x1b[31mGot:\tWant:\n\r"

		for row := 0; row < well.Height + 1; row++ {
			errMsg += got[row] + "\t" + want[row] + "\n\r"
		}

		errMsg += "\x1b[0m"

		t.Error(errMsg)
	} else {
		print("\n\r\x1b[32m" + strings.Join(got, "\n\r") + "\x1b[0m\n\r")
	}

}

func takeSnapshot(well Well) snapshot {
	img := make([]string, well.Height + 1)

	for row := 0; row < well.Height; row++ {
		line := "|"

		for col := 0; col < well.Width; col++ {
			if well.IsOccupiedAt(Coords{row, col}) {
				line += "x"
			} else {
				line += " "
			}
		}

		line += "|"

		img[row] = line
	}

	img[well.Height] = "|" + strings.Repeat("-", well.Width) + "|"

	return img
}

func logSnapshot(well Well) {
	print("\n\r\x1b[37m")
	for _, line := range(takeSnapshot(well)) {
		println(line)
	}
	print("\x1b[0m\n\r")
}

func assertFilm(t *testing.T, well *Well, actions actions, wantFlm film) {
	t.Helper()

	var want, got []snapshot

	want = filmToSnapshots(wantFlm)

	got = append(got, takeSnapshot(*well))
	for _, action := range(actions) {
		action()
		got = append(got, takeSnapshot(*well))
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
