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

func extractCard(line string) (id int, winningNumbers []int, playerNumbers []int) {
	if groups := strings.Split(line, ":"); len(groups) > 1 {
		if cardProperties := strings.Split(groups[0], " "); len(cardProperties) > 1 {
			if i, err := strconv.Atoi(strings.Trim(cardProperties[len(cardProperties)-1], " ")); err == nil {
				id = i
			}
		}
		if lists := strings.Split(groups[1], "|"); len(lists) > 1 {
			winningNumbers = getNumbers(lists[0])
			playerNumbers = getNumbers(lists[1])
		}
	}
	return id, winningNumbers, playerNumbers
}

func computeWinningCards(w []int, p []int) (occ float64) {
	for _, wn := range w {
		for _, pn := range p {
			if wn == pn {
				occ++
			}
		}
	}
	return occ
}

func copyCards(from int, to float64, lim int) (copies []int) {
	k := from + 1
	for float64(k) <= float64(from)+to {
		if k <= lim {
			copies = append(copies, k)
		}
		k++
	}
	return copies
}

func solution1(lines []string) (sum float64) {
	for _, line := range lines {
		_, w, p := extractCard(line)
		occ := computeWinningCards(w, p)
		if occ > 0 {
			sum += math.Pow(base, occ-1)
		}
	}
	return sum
}

func solution2(lines []string) (stackSize int) {
	lim := len(lines)
	copiesByCardType := map[int][]int{}
	originals := []int{}
	copies := []int{}

	for _, line := range lines {
		i, w, p := extractCard(line)
		originals = append(originals, i)
		occ := computeWinningCards(w, p)
		c := copyCards(i, occ, lim)
		copies = append(copies, c...)
		copiesByCardType[i] = c
	}

	stackSize += len(originals)
	stackSize += len(copies)

	originals = copies
	copies = []int{}

	for len(originals) > 0 {
		for _, o := range originals {
			c := copiesByCardType[o]
			copies = append(copies, c...)
		}
		stackSize += len(copies)
		originals = copies
		copies = []int{}
	}
	return stackSize
}

func Run() {
	lines := util.LoadInput("day4/input.txt")
	fmt.Println("[DAY 4][PART 1]", solution1(lines))
	fmt.Println("[DAY 4][PART 2]", solution2(lines))
}
