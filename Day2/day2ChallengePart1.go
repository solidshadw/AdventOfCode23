package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameData := strings.Split(line, ": ")
		gameID, _ := strconv.Atoi(strings.TrimPrefix(gameData[0], "Game "))
		rounds := strings.Split(gameData[1], "; ")

		isPossible := true
		for _, round := range rounds {
			colors := strings.Split(round, ", ")
			for _, color := range colors {
				colorData := strings.Split(color, " ")
				count, _ := strconv.Atoi(colorData[0])
				if (colorData[1] == "red" && count > 12) || (colorData[1] == "green" && count > 13) || (colorData[1] == "blue" && count > 14) {
					isPossible = false
					break
				}
			}
			if !isPossible {
				break
			}
		}

		if isPossible {
			sum += gameID
		}
	}

	fmt.Println("Sum of possible game IDs:", sum)
}
