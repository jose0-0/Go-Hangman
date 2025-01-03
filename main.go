package main

import "fmt"

func main () {
	word := "golang"
	attempts := 6
	currentWordState := initializeCurrentWordState(word)

	fmt.Println("Welcome to Hangman in Go!")
}

func initializeCurrentWordState(word string) []string {
	// the "make()" function allocates a zeroed array and returns a slice that refers to that array 
	currentWordState := make([]string, len(word))
	
	for i := range currentWordState {
		currentWordState[i] = "_"
	}

	return currentWordState
}