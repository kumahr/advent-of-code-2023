package day2

import (
	"adventofcode2023/day2/gameparser"
	"adventofcode2023/day2/model"
	"adventofcode2023/day2/part1"
	"adventofcode2023/util"
	"fmt"
)

func Run() {
	lines := util.LoadInput("day2/input.txt")
	world := &model.World{MaxRed: 12, MaxGreen: 13, MaxBlue: 14}
	games := gameparser.ParseGames(lines)
	fmt.Println("[DAY 2][PART 1]", part1.SumPossibleGames(world, games))
}
