package main

import (
	"testing"
)

func TestPuzzle1(t *testing.T) {
	expected := 9876
	actual := puzzle1(`1000
8000
876

1000
1000
2345`)

	if expected != actual {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}

func TestPuzzle2(t *testing.T) {
	actual := puzzle2(`1000
8000
876

1000
1000
2345

1
1
1

298
293
7654`)
	expected := 22466

	if expected != actual {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}
