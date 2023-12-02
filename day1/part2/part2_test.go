package part2

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestSumCalibrationValues(t *testing.T) {
	file, err := os.Open("input_test.txt")
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

	got := SumCalibrationValues(lines)
	want := 281

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
