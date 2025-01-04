# Hangman in Go: A Learning Project

The Hangman game is a simple and fun way to learn the basics of the Go programming language. This project involves implementing the logic, user interaction, and error handling necessary to play the classic word-guessing game. Below is a breakdown of how the game works and what you can learn from it:

## Game Description

- Objective: Guess the hidden word one letter at a time before running out of attempts.
- Setup: A word is randomly selected from a predefined list, and the player is shown underscores representing each letter of the word.

### Gameplay

1. The player guesses one letter at a time.
2. Correct guesses reveal the letter's position(s) in the word.
3. Incorrect guesses reduce the remaining attempts and update a visual "hangman" display.
4. Winning Condition: The player reveals the entire word before using all their attempts.
5. Losing Condition: The player runs out of attempts without guessing the word.

## Setup

1. Clone the repository

## Running the game

There is no additional steps required to run the game. You can simply do the following:

1. run command `go run ./` in the CLI

### Key Features in Go Implementation

1. Word Selection

- Use a slice to store a list of words.
- Randomly select a word using Go's `math/rand` package.

2. User Input and validation

- Read users input from the console.
- Validate the input to ensure it is a single letter and hasn't been guessed already.

3. Game Logic

- Keep track of guessed letters, remaining attempts and the current state of the word.
- Update the games state dynamically based on correct and incorrect guesses.

4. Error Handling

- Handle invalid inputs gracefully (i.e Numbers, special characters, repeated guesses)
- Provide clear feedback to the player

5. Visual Representation

- Create a simple visual representation of the hangman
- Update the display after incorrect guess.

6. Game Loop

- Use a loop to manage the games progression
- Break out of the loop when the player wins or losses.

## Learning Objective

- Functions: Encapsulate game logic into reusable functions
- Data Structures: Work with slices and maps for mapping the word list and tracking guesses.
- Error handling: Handle invalid inputs using Go's idiomatic error management.
- Control Flow: Practice using `for` loops, conditional (`if`/`else`), and `switch` statements.
