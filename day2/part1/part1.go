package part1

import "adventofcode2023/day2/model"

func isGamePossible(world *model.World, draws []model.Draw) bool {
	for _, draw := range draws {
		if draw.Red > world.MaxRed || draw.Green > world.MaxGreen || draw.Blue > world.MaxBlue {
			return false
		}
	}
	return true
}

func getPossibleGameId(world *model.World, game *model.Game) int {
	if isGamePossible(world, game.Draws) {
		return game.Id
	}
	return 0
}

func SumPossibleGames(world *model.World, games []model.Game) (sum int) {
	sum = 0
	for _, game := range games {
		sum += getPossibleGameId(world, &game)
	}
	return sum
}
