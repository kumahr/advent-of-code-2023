package part2

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func detectWord(text string, reverse bool) string {
	m := strDigits
	if reverse {
		m = revStrDigits
	}

	for k, v := range m {
		if strings.Contains(text, k) {
			return fmt.Sprint(v)
		}
	}
	return ""
}

func getCalibrationValue(line string) int {
	x := 0
	y := len(line) - 1

	leftWord := string(line[x])
	rightWord := string(line[y])

	leftDigit := ""
	rightDigit := ""

	// Scan line from both sides at the same time so we can stop early when both digits have been detected
	// Either we detect a digit from a string number or from a word
	for (x < len(line) || y > 0) && (len(leftDigit) == 0 || len(rightDigit) == 0) {
		leftChar := []rune(line)[x]
		rightChar := []rune(line)[y]

		if len(leftDigit) == 0 && unicode.IsDigit(leftChar) {
			leftDigit = string(leftChar)
		}
		if len(rightDigit) == 0 && unicode.IsDigit(rightChar) {
			rightDigit = string(rightChar)
		}
		detectedLeftWord := detectWord(leftWord, false)
		detectedRightWord := detectWord(rightWord, true)
		if len(detectedLeftWord) > 0 {
			if len(leftDigit) == 0 {
				leftDigit = detectedLeftWord
			}
			leftWord = leftWord[len(leftWord)-1:]
		}
		if len(detectedRightWord) > 0 {
			if len(rightDigit) == 0 {
				rightDigit = detectedRightWord
			}
			rightWord = rightWord[len(rightWord)-1:]
		}

		if x >= len(line)-1 {
			leftWord += string(line[len(line)-1])
		} else {
			leftWord += string(line[x+1])
		}

		if y <= 0 {
			rightWord += string(line[0])
		} else {
			rightWord += string(line[y-1])
		}

		x++
		y--
	}
	strCalibrationValue := leftDigit + rightDigit
	if len(strCalibrationValue) > 0 {
		n, _ := strconv.Atoi(strCalibrationValue)
		return n
	}
	return 0
}

func SumCalibrationValues(lines []string) int {
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
