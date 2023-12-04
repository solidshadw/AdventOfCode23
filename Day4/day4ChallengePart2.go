package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Card struct { // Define a struct to represent a card
	winningNumbers []string // Define a slice of strings to represent the winning numbers
	yourNumbers    []string // Define a slice of strings to represent your numbers
	instances      int      // Define an integer to represent the number of instances
	winnerCards    int      // Define an integer to represent the number of winning cards
}

func main() {
	file, _ := os.Open("input.txt")   // Open the "input.txt" file
	defer file.Close()                // Close the file when we're done with it
	scanner := bufio.NewScanner(file) // Create a scanner to read the file line by line

	cards := parseCards(scanner) // Parse the cards from the input file

	for i := range cards { // Loop through each card
		cards[i].calculateWinnerCards() // Calculate the number of winning cards for the current card
	}

	for i := range cards { // Loop through each card
		for times := 1; times <= cards[i].instances; times++ { // Loop through each instance of the current card
			for j := i + 1; j <= i+cards[i].winnerCards && j < len(cards); j++ { // Loop through each card that is a winner for the current card
				cards[j].instances++ // Increment the number of instances for the current card
			}
		}
	}

	totalInstances := sumInstances(cards) // Calculate the total number of instances for all cards

	fmt.Println(totalInstances) // Print the total number of instances
}

func parseCards(scanner *bufio.Scanner) []Card { // Parse the cards from the input file
	cards := make([]Card, 0) // Initialize a slice of cards
	for scanner.Scan() {     // Loop through each line in the file
		line := scanner.Text()            // Get the current line as a string
		parts := strings.Split(line, "|") // Split the line into two parts using the "|" separator
		card := Card{                     // Create a new card
			winningNumbers: strings.Fields(parts[0]), // Get the winning numbers as a slice of strings
			yourNumbers:    strings.Fields(parts[1]), // Get your numbers as a slice of strings
			instances:      1,                        // Initialize the number of instances to 1
		}
		cards = append(cards, card) // Add the card to the slice of cards
	}
	return cards // Return the slice of cards
}

func (c *Card) calculateWinnerCards() { // Calculate the number of winning cards for the current card
	for _, yourNumber := range c.yourNumbers { // Loop through each of your numbers
		for _, winningNumber := range c.winningNumbers { // Loop through each winning number
			if yourNumber == winningNumber { // Check if your number matches the winning number
				c.winnerCards++ // If it does, increment the number of winning cards
				break           // Exit the inner loop since we found a match
			}
		}
	}
}

func sumInstances(cards []Card) int { // Calculate the total number of instances for all cards
	totalInstances := 0          // Initialize a variable to keep track of the total number of instances
	for _, card := range cards { // Loop through each card
		totalInstances += card.instances // Add the number of instances for the current card to the total number of instances
	}
	return totalInstances // Return the total number of instances
}
