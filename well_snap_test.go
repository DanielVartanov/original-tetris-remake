package main

import "testing"

func TestWell_Snap_None(t *testing.T) {
	w := NewWell(3, 6)
	fillWell(w, snapshot{
		"| xxxxx|",
		"| xxxxx|",
		"| xxxxx|",
		"|------|",
	})

	if w.Snap() {
		t.Error()
	}

	assertSnapshot(t, w, snapshot{
		"| xxxxx|",
		"| xxxxx|",
		"| xxxxx|",
		"|------|",
	})
}

func TestWell_Snap_One(t *testing.T) {
	w := NewWell(3, 4)
	fillWell(w, snapshot{
		"| xxx|",
		"|xxxx|",
		"|xxx |",
		"|----|",
	})

	assertFilm(t, &w,
		actions{
			func() { if !w.Snap() { t.Error() } },
			func() { if w.Snap() { t.Error() } },
			func() { if w.Snap() { t.Error() } },
		},
		film{
			{"| xxx|", "|    |", "|    |", "|    |"},
			{"|xxxx|", "| xxx|", "| xxx|", "| xxx|"},
			{"|xxx |", "|xxx |", "|xxx |", "|xxx |"},
			{"|----|", "|----|", "|----|", "|----|"},
		},
	)
}

func TestWell_Snap_Multiple(t *testing.T) {
	w := NewWell(3, 4)
	fillWell(w, snapshot{
		"|xxxx|",
		"|x  x|",
		"|xxxx|",
		"|----|",
	})

	assertFilm(t, &w,
		actions{
			func() { if !w.Snap() { t.Error() } },
			func() { if !w.Snap() { t.Error() } },
			func() { if w.Snap() { t.Error() } },
		},
		film{
			{"|xxxx|", "|    |", "|    |", "|    |"},
			{"|x  x|", "|x  x|", "|    |", "|    |"},
			{"|xxxx|", "|xxxx|", "|x  x|", "|x  x|"},
			{"|----|", "|----|", "|----|", "|----|"},
		},
	)
}
