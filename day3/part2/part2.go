package part2

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Cell struct {
	X int
	Y int
}

type Matrix struct {
	Grid []string
	Dim  Cell
}

func (m *Matrix) isInRange(c Cell) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < m.Dim.X && c.Y < m.Dim.Y
}

func isSymbol(ch rune) bool {
	return !unicode.IsDigit(ch) && !(ch == '.')
}

func isStar(ch rune) bool {
	return ch == '*'
}

func getMatrixSize(grid []string) Cell {
	x := len(grid)
	var y int
	if x > 0 {
		y = len(grid[0])
	}
	return Cell{X: x, Y: y}
}

func digitIsAdjacentToSymbol(matrix Matrix, x int, y int) (bool, []string) {
	top := Cell{X: x - 1, Y: y}
	bot := Cell{X: x + 1, Y: y}
	left := Cell{X: x, Y: y - 1}
	right := Cell{X: x, Y: y + 1}
	topLeft := Cell{X: x - 1, Y: y - 1}
	topRight := Cell{X: x - 1, Y: y + 1}
	botLeft := Cell{X: x + 1, Y: y - 1}
	botRight := Cell{X: x + 1, Y: y + 1}
	adjacentCells := []Cell{top, bot, left, right, topLeft, topRight, botLeft, botRight}
	gears := []string{}
	found := false
	for _, cell := range adjacentCells {
		if matrix.isInRange(cell) && isSymbol([]rune(matrix.Grid[cell.X])[cell.Y]) {
			found = true
			if isStar([]rune(matrix.Grid[cell.X])[cell.Y]) {
				gears = append(gears, fmt.Sprint(cell))
			}
		}
	}
	return found, gears
}

func resetSearch(fAdjCell *bool, gearsByPart *[]string, m map[string][]string, chain *string) {
	if *fAdjCell {
		m[*chain] = *gearsByPart
	}
	*fAdjCell = false
	*gearsByPart = []string{}
	*chain = ""

}

func pushGear(gears []string, gear string) []string {
	nGears := gears
	if len(nGears) == 0 {
		nGears = append(nGears, gear)
	} else {
		isPresent := false
		for _, g := range gears {
			if g == gear {
				isPresent = true
			}
		}
		if !isPresent {
			nGears = append(nGears, gear)
		}
	}
	return nGears
}

func pushStars(hostGears []string, newGears []string) []string {
	nGears := hostGears
	for _, ng := range newGears {
		nGears = pushGear(hostGears, ng)
	}
	return nGears
}

func SumGears(lines []string) (sum int) {
	sum = 0
	matrix := Matrix{Grid: lines, Dim: getMatrixSize(lines)}
	m := map[string][]string{}
	for x, line := range lines {
		chain := ""
		gearsByPart := []string{}
		fAdjCell := false
		for y, ch := range line {
			if unicode.IsDigit(ch) {
				chain += fmt.Sprintf("%d->%d;", x, y)
				if b, gears := digitIsAdjacentToSymbol(matrix, x, y); b {
					fAdjCell = true
					gearsByPart = pushStars(gearsByPart, gears)
				}
				if y == len(line)-1 {
					resetSearch(&fAdjCell, &gearsByPart, m, &chain)
				}
			} else {
				resetSearch(&fAdjCell, &gearsByPart, m, &chain)
			}
		}
	}

	gearRatios := map[string][]string{}

	for k, v := range m {
		for _, i := range v {
			grs := []string{}
			for q, j := range m {
				if q != k {
					for _, jj := range j {
						if jj == i {
							grs = pushGear(grs, k)
							grs = pushGear(grs, q)
						}
					}
				}
			}
			gearRatios[i] = grs
		}
	}

	for _, v := range gearRatios {
		if len(v) == 2 {
			mul := 1
			for _, s := range v {
				str := ""
				vec := strings.Split(s, ";")
				for _, e := range vec {
					if len(e) > 0 {
						f := strings.Split(e, "->")
						x, _ := strconv.Atoi(f[0])
						y, _ := strconv.Atoi(f[1])
						str += string(lines[x][y])
					}

				}
				n, _ := strconv.Atoi(str)
				mul *= n
			}
			sum += mul
		}
	}
	return sum
}
