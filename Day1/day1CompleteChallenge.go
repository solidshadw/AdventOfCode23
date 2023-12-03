package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func extractCalibrationValue(line string) (int, error) {
	// Map of spelled out digits to their numeric values
	digitMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	// Replace spelled out digits with their numeric values
	for word, digit := range digitMap {
		line = strings.ReplaceAll(line, word, digit)
	}

	// Extract the first and last digit from the line
	re := regexp.MustCompile(`\d`)
	digits := re.FindAllString(line, -1)

	// If there is only one digit, duplicate it to form a two-digit number
	if len(digits) == 1 {
		digits = append(digits, digits[0])
	}

	// Check if there are at least two digits
	if len(digits) < 2 {
		return 0, fmt.Errorf("not enough digits in the line '%s'", line)
	}
	fmt.Println("Digits in the line:", digits)

	// Combine the first and last digits to form a two-digit number
	calibrationValueStr := string(digits[0]) + string(digits[len(digits)-1])

	// Convert the combined string to an integer
	calibrationValue, err := strconv.Atoi(calibrationValueStr)
	if err != nil {
		return 0, err
	}

	return calibrationValue, nil
}

func challengeComplete() {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize a variable to store the sum of calibration values
	totalSum := 0

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Extract the calibration value from the line
		calibrationValue, err := extractCalibrationValue(line)
		if err != nil {
			fmt.Printf("Error extracting calibration value from line '%s': %v\n", line, err)
			continue
		}

		// Print the extracted calibration value for each line
		fmt.Printf("Line: %s, Extracted Calibration Value: %d\n", line, calibrationValue)

		// Add the calibration value to the total sum
		totalSum += calibrationValue
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Print the total sum of calibration values at the end
	fmt.Println("Total Sum of Calibration Values:", totalSum)
}
