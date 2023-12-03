package part1

import (
	"adventofcode2023/util"
	"testing"
)

func TestSumPartNumbers(t *testing.T) {
	lines := util.LoadInput("input_test.txt")
	got := SumPartNumbers(lines)
	want := 4361

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
