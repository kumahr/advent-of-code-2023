package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const ZERO = '0'
const NINE = '9'

func getCalibrationValue(line string) int {
	var digits []int
	for _, ch := range line {
		if ch >= ZERO && ch <= NINE {
			digits = append(digits, int(ch-ZERO))
		}
	}
	if len(digits) > 0 {
		var twoDigits []int
		firstDigit := digits[0]
		var secondDigit int
		if len(digits) > 1 {
			secondDigit = digits[len(digits)-1]
		} else {
			secondDigit = firstDigit
		}
		twoDigits = append(twoDigits, firstDigit, secondDigit)
		strNumber := ""
		for _, e := range twoDigits {
			strNumber += strconv.Itoa(e)
		}
		calibrationValue, err := strconv.Atoi(strNumber)
		if err != nil {
			log.Fatalln(err)
		}
		return calibrationValue
	}
	return 0
}

func sumCalibrationValues(file *os.File) int {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var calibrationValues []int
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		calibrationValues = append(calibrationValues, getCalibrationValue(txt))
	}
	var sum int
	for _, v := range calibrationValues {
		sum += v
	}
	return sum
}

func RunPart1() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file", err)
	}
	defer file.Close()
	log.Println(sumCalibrationValues(file))
}
