package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type PosNumber struct {
	StartPosition int
	EndPosition   int
	LineIndex     int
	Number        int
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

func numberTouchesSymbol(number *PosNumber, lineIndex, starPos, numberoflines int) bool {
	for looplines := max(0, lineIndex-1); looplines <= min(numberoflines-1, lineIndex+1); looplines++ {
		if number.LineIndex != looplines {
			continue
		}
		if number.StartPosition == starPos || number.StartPosition == starPos+1 || number.EndPosition == starPos || number.EndPosition-1 == starPos {
			return true
		}
		if number.LineIndex != lineIndex && number.StartPosition <= starPos && number.EndPosition >= starPos {
			return true
		}
	}

	return false
}

func findCharPositions(str string, targetChar rune) []int {
	positions := []int{}
	for i, char := range str {
		if char == targetChar {
			positions = append(positions, i)
		}
	}
	return positions
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
		numbersForLine := findNumbers(line, lineIndex)
		allNumbers = append(allNumbers, numbersForLine...)
	}

	totalSum := 0
	for lineIndexofStar, l := range lines {
		starPositions := findCharPositions(l, '*')

		for _, starPos := range starPositions {
			var foundNumbers []*PosNumber
			for _, number := range allNumbers {
				if numberTouchesSymbol(number, lineIndexofStar, starPos, len(lines)) {
					foundNumbers = append(foundNumbers, number)
				}
			}
			if len(foundNumbers) == 2 {
				totalSum += foundNumbers[0].Number * foundNumbers[1].Number
			}
		}
	}

	fmt.Printf("totalSum: %d\n", totalSum)
}
