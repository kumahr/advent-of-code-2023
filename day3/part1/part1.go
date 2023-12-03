package part1

import (
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

func getMatrixSize(grid []string) Cell {
	x := len(grid)
	var y int
	if x > 0 {
		y = len(grid[0])
	}
	return Cell{X: x, Y: y}
}

func digitIsAdjacentToSymbol(matrix Matrix, x int, y int) bool {
	top := Cell{X: x - 1, Y: y}
	bot := Cell{X: x + 1, Y: y}
	left := Cell{X: x, Y: y - 1}
	right := Cell{X: x, Y: y + 1}
	topLeft := Cell{X: x - 1, Y: y - 1}
	topRight := Cell{X: x - 1, Y: y + 1}
	botLeft := Cell{X: x + 1, Y: y - 1}
	botRight := Cell{X: x + 1, Y: y + 1}
	adjacentCells := []Cell{top, bot, left, right, topLeft, topRight, botLeft, botRight}
	for _, cell := range adjacentCells {
		if matrix.isInRange(cell) && isSymbol([]rune(matrix.Grid[cell.X])[cell.Y]) {
			return true
		}
	}
	return false
}

func resetSearch(fAdjCell *bool, pNumbers *string, pNumber *string) {
	if *fAdjCell {
		*pNumbers += " " + *pNumber
	}
	*pNumber = ""
	*fAdjCell = false
}

func SumPartNumbers(lines []string) (sum int) {
	sum = 0
	strPartNumbers := ""
	matrix := Matrix{Grid: lines, Dim: getMatrixSize(lines)}
	for x, line := range lines {
		pNumbers := ""
		pNumber := ""
		fAdjCell := false
		for y, ch := range line {
			if unicode.IsDigit(ch) {
				pNumber += string(ch)
				if digitIsAdjacentToSymbol(matrix, x, y) {
					fAdjCell = true
				}
				if y == len(line)-1 {
					resetSearch(&fAdjCell, &pNumbers, &pNumber)
				}
			} else {
				resetSearch(&fAdjCell, &pNumbers, &pNumber)
			}
		}
		strPartNumbers += " " + pNumbers
	}
	for _, str := range strings.Split(strPartNumbers, " ") {
		if str != " " && str != "" {
			n, _ := strconv.Atoi(str)
			sum += n
		}
	}
	return sum
}
