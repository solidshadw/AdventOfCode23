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

	totalPower := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rounds := strings.Split(strings.Split(line, ": ")[1], "; ")

		minCubes := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, round := range rounds {
			colors := strings.Split(round, ", ")
			for _, color := range colors {
				colorData := strings.Split(color, " ")
				count, _ := strconv.Atoi(colorData[0])
				if count > minCubes[colorData[1]] {
					minCubes[colorData[1]] = count
				}
			}
		}

		power := minCubes["red"] * minCubes["green"] * minCubes["blue"]
		totalPower += power
	}

	fmt.Println("Total power:", totalPower)
}
