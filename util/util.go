package util

import (
	"bufio"
	"log"
	"os"
)

func LoadInput(path string) (lines []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Println("Could not open input file", err)
		return lines
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
