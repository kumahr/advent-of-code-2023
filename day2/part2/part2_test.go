package part2

import (
	"adventofcode2023/day2/gameparser"
	"adventofcode2023/util"
	"testing"
)

func TestSumPowers(t *testing.T) {
	lines := util.LoadInput("input_test.txt")
	games := gameparser.ParseGames(lines)

	got := SumPowers(games)
	want := 2286

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
