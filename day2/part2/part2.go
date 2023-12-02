package part2

import "adventofcode2023/day2/model"

func findMinimumPower(sets []model.Draw) int {
	minimumSet := model.Draw{Red: 0, Green: 0, Blue: 0}
	for _, set := range sets {
		if set.Red > minimumSet.Red {
			minimumSet.Red = set.Red
		}
		if set.Green > minimumSet.Green {
			minimumSet.Green = set.Green
		}
		if set.Blue > minimumSet.Blue {
			minimumSet.Blue = set.Blue
		}
	}
	return minimumSet.Power()
}

func SumPowers(games []model.Game) (sum int) {
	sum = 0
	for _, game := range games {
		sum += findMinimumPower(game.Draws)
	}
	return sum
}
