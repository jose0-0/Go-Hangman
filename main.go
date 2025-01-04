package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main () {
	word := "golang"
	attempts := 6
	currentWordState := initializeCurrentWordState(word)
	scanner := bufio.NewScanner(os.Stdin)
	guessedLetters := make(map[string]bool)

	fmt.Println("Welcome to Hangman in Go!")
	for attempts > 0 {
		displayCurrentWordState(currentWordState, attempts)
		userInput := getUserInput(scanner)

		if !isValidInput(userInput) {
			fmt.Println("Invalid input, Please enter a single letter")
			continue
		}

		if guessedLetters[userInput] {
			fmt.Println("You've already guessed that letter")
			continue
		}

		guessedLetters[userInput] = true

		correctGuess := updateGuessed(word, currentWordState, userInput)

		if !correctGuess {
			attempts--
		}

		displayHangman(6 - attempts)

	}
}

func displayHangman(incorrectGuesses int) {
	if incorrectGuesses >= 0 && incorrectGuesses < len(hangmanStates) {
		fmt.Println(hangmanStates[incorrectGuesses])
	}
}

func updateGuessed(word string, guessed []string, letter string) bool {
	correctGuess := false

	// if a match is found, the underscore will be replaced by the letter
	for i, char := range word {
		if string(char) == letter {
			guessed[i] = letter
			correctGuess = true
		}
	}

	return correctGuess
}

func isValidInput (input string) bool {
	// returns the amount of runes in a string
	return utf8.RuneCountInString(input) == 1
}

func getUserInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func initializeCurrentWordState(word string) []string {
	// the "make()" function allocates a zeroed array and returns a slice that refers to that array 
	currentWordState := make([]string, len(word))
	
	for i := range currentWordState {
		currentWordState[i] = "_"
	}

	return currentWordState
}

func displayCurrentWordState(currentWordState []string, attempts int) {
	// The Joins() function takes a given separator and concatenates of a slice of strings into a single string
	fmt.Println("Current Word State:", strings.Join(currentWordState, " "))
	fmt.Println("Attempts Left:", attempts)
}