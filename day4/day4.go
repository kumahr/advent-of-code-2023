package day4

import (
	"adventofcode2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const base = 2

func getNumbers(str string) (l []int) {
	list := strings.Trim(str, " ")
	if nums := strings.Split(list, " "); len(nums) > 1 {
		for _, v := range nums {
			if n, err := strconv.Atoi(v); err == nil {
				l = append(l, n)
			}
		}
	}
	return l
}

func extractLists(line string) (winningNumbers []int, playerNumbers []int) {
	if groups := strings.Split(line, ":"); len(groups) > 1 {
		if lists := strings.Split(groups[1], "|"); len(lists) > 1 {
			winningNumbers = getNumbers(lists[0])
			playerNumbers = getNumbers(lists[1])
		}
	}
	return winningNumbers, playerNumbers
}

func solution1(lines []string) (sum float64) {
	for _, line := range lines {
		var occ float64
		w, p := extractLists(line)
		for _, wn := range w {
			for _, pn := range p {
				if wn == pn {
					occ++
				}
			}
		}
		if occ > 0 {
			sum += math.Pow(base, occ-1)
		}
	}
	return sum
}

func Run() {
	lines := util.LoadInput("day4/input.txt")
	fmt.Println("[DAY 4][PART 1]", solution1(lines))
}
