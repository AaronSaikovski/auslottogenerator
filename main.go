/*
SUMMARY:      Generates random numbers for use in NSW lotteries
DESCRIPTION:  Allows generation of numbers for Lotto, OzLotto and Powerball
			  Lotto - Allows for 1-45 numbers with 6 numbers per game and a maximum of 50 games per entry
			  Ozlotto - Allows for 1-47 numbers with 7 numbers per game and a maximum of 50 games per entry
			  Powerball - Allows for 1-35 numbers with 7 numbers per game and 1-20 single number for the Powerball number with a total maximum of 50 games per entry

AUTHOR/S:     asaikovski
VERSION:      1.3.0

VERSION HISTORY:
  1.0.0 - Initial version release
  1.1.0 - Fixed random number duplication results bug, optimisations and bug fixes. also added pre-commit hooks
  1.2.0 - Fixed/removed OS specific input readers - Verified working on Windows 11 and Mac OS Ventura 13.1
  1.3.0 - Optimised code - added input function
*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Generates the random numbers No duplicates - returns an array of random numbers per call
func generateRandomNumbers(maxVal int, numbersPerGame int) []int {

	// Create a slice to store the generated numbers
	var numbers []int

	//seed the randomiser
	rand.Seed(time.Now().UnixNano())

	// Generate random numbers until we have numbersPerGame of them
	for len(numbers) < numbersPerGame {

		// Generate a random number between 1 and maxVal
		n := rand.Intn(maxVal) + 1

		// Check if the number has already been generated
		duplicate := false
		for _, v := range numbers {
			if v == n {
				duplicate = true
				break
			}
		}

		// If the number is not a duplicate, add it to the slice
		if !duplicate {
			numbers = append(numbers, n)
		}
	}

	//return the results array
	return numbers
}

//get the console input
func getInput(prompt string) int {
	fmt.Print(prompt)
    var input string
    fmt.Scanln(&input)
    num, _ := strconv.ParseInt(strings.TrimSuffix(input, "\n"), 10, 64)
    return int(num)
}

// Main
func main() {

	fmt.Println("***************************************")
	fmt.Println("** Lottery number generator - v1.3.0 **")
	fmt.Println("***************************************")

	// get input for number of games
	maxNumberGames := getInput("How many games to play?: ")

	// Set the maximum random numbers per game
	maxRandomNumbersPerGame := getInput("Random number pool to use per game?: ")

	// get the maximum numbers per game
	maxNumbersPerGame := getInput("Maximum numbers per game?: ")

	fmt.Println("\n******************************")
	fmt.Println("** Results **")

	// Loop over number of games to play - generate random numbers each iteration
	for i := 0; i < maxNumberGames; i++ {
		fmt.Println("Game", i+1, generateRandomNumbers(maxRandomNumbersPerGame, maxNumbersPerGame))
	}
	fmt.Println("******************************")
}
