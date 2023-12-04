package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the "input.txt" file
	file, _ := os.Open("input.txt")
	defer file.Close() // Close the file when we're done with it

	scanner := bufio.NewScanner(file) // Create a scanner to read the file line by line
	totalPoints := 0                  // Initialize a variable to keep track of the total points

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()            // Get the current line as a string
		parts := strings.Split(line, "|") // Split the line into two parts using the "|" separator
		winningNumbers := strings.Fields(parts[0])
		yourNumbers := strings.Fields(parts[1]) // Get your numbers as a slice of strings

		points := 0 // Initialize a variable to keep track of the points for the current line

		// Loop through each of your numbers
		for _, yourNumber := range yourNumbers {
			// Loop through each winning number
			for _, winningNumber := range winningNumbers {
				// Check if your number matches the winning number
				if yourNumber == winningNumber {
					// If it does, update the points based on the current points value
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
					break // Exit the inner loop since we found a match
				}
			}
		}

		totalPoints += points // Add the points for the current line to the total points
	}

	fmt.Println(totalPoints) // Print the total points
}
