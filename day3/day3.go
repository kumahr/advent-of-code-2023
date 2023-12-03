package day3

import (
	"adventofcode2023/day3/part1"
	"adventofcode2023/day3/part2"
	"adventofcode2023/util"
	"fmt"
)

func Run() {
	lines := util.LoadInput("day3/input.txt")
	fmt.Println("[DAY 3][PART 1]", part1.SumPartNumbers(lines))
	fmt.Println("[DAY 3][PART 2]", part2.SumGears(lines))
}
