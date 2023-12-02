package day1

import (
	"adventofcode2023/day1/part1"
	"adventofcode2023/day1/part2"
	"adventofcode2023/util"
	"fmt"
)

func Run(inputPath string) {
	lines := util.LoadInput(inputPath)
	fmt.Println("[DAY 1][PART 1]", part1.SumCalibrationValues(lines))
	fmt.Println("[DAY 1][PART 2]", part2.SumCalibrationValues(lines))
}
