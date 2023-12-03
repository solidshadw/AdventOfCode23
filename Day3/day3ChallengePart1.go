package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PosNumber struct {
	StartPosition int
	EndPosition   int
	LineIndex     int
	Number        int
}

func symbolAround(number *PosNumber, lines []string) bool {
	from, to := number.StartPosition-1, number.EndPosition+1
	if from < 0 {
		from = 0
	}
	if to > len(lines[0]) {
		to = len(lines[0])
	}

	for looplines := number.LineIndex - 1; looplines <= number.LineIndex+1; looplines++ {
		if looplines < 0 || looplines >= len(lines) {
			continue
		}
		if strings.IndexAny(lines[looplines][from:to], "+#$*@/=%-&") > -1 {
			return true
		}
	}

	return false
}

func findNumbers(str string, lineIndex int) []*PosNumber {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringSubmatchIndex(str, -1)

	result := make([]*PosNumber, len(matches))
	for i, match := range matches {
		start, end := match[0], match[1]
		number, _ := strconv.Atoi(str[start:end])

		result[i] = &PosNumber{
			Number:        number,
			StartPosition: start,
			EndPosition:   end,
			LineIndex:     lineIndex,
		}
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var allNumbers []*PosNumber
	for lineIndex, line := range lines {
		allNumbers = append(allNumbers, findNumbers(line, lineIndex)...)
	}

	totalSum := 0
	for _, number := range allNumbers {
		if symbolAround(number, lines) {
			totalSum += number.Number
		}
	}

	fmt.Printf("totalSum: %d\n", totalSum)
}
