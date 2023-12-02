package part1

import (
	"adventofcode2023/day2/model"
	"testing"
)

func TestSumPossibleGames(t *testing.T) {
	world := &model.World{MaxRed: 12, MaxGreen: 13, MaxBlue: 14}
	games := []model.Game{
		model.Game{Id: 1, Draws: []model.Draw{
			model.Draw{
				Red:   1,
				Green: 4,
				Blue:  8,
			},
			model.Draw{
				Green: 5,
				Blue:  1,
			},
			model.Draw{
				Blue: 8,
			},
		},
		},
		model.Game{Id: 2, Draws: []model.Draw{
			model.Draw{
				Red:   1,
				Green: 43,
				Blue:  8,
			},
			model.Draw{
				Green: 5,
				Blue:  1,
			},
			model.Draw{
				Blue: 8,
			},
		},
		},
		model.Game{Id: 3, Draws: []model.Draw{
			model.Draw{
				Red: 1,
			},
			model.Draw{
				Blue: 1,
			},
			model.Draw{
				Green: 8,
			},
		},
		},
	}

	got := SumPossibleGames(world, games)
	want := 4

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
