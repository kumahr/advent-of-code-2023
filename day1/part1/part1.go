package part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func getCalibrationValue(line string) int {
	var firstDigit string
	var secondDigit string

	x := 0
	y := len(line) - 1

	for (x < len(line) || y > 0) && (len(firstDigit) == 0 || len(secondDigit) == 0) {
		leftChar := []rune(line)[x]
		rightChar := []rune(line)[y]

		if unicode.IsDigit(leftChar) && len(firstDigit) == 0 {
			firstDigit = string(leftChar)
		}
		if unicode.IsDigit(rightChar) && len(secondDigit) == 0 {
			secondDigit = string(rightChar)
		}
		x++
		y--
	}
	calibrationValue := firstDigit + secondDigit
	if len(calibrationValue) > 0 {
		n, _ := strconv.Atoi(calibrationValue)
		return n
	}
	return 0
}

func sumCalibrationValues(lines []string) int {
	var calibrationValues []int
	for _, line := range lines {
		calibrationValues = append(calibrationValues, getCalibrationValue(line))
	}
	var sum int
	for _, v := range calibrationValues {
		sum += v
	}
	return sum
}

func RunPart1(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln("Could not open input file", err)
	}
	defer file.Close()

	var lines []string
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	log.Println(sumCalibrationValues(lines))
}
