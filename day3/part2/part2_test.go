package part2

import (
	"adventofcode2023/util"
	"testing"
)

func TestSumPartNumbers(t *testing.T) {
	lines := util.LoadInput("input_test.txt")
	got := SumGears(lines)
	want := 467835

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
