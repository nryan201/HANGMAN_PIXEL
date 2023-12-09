package WindowsHallowen

import (
	"HANGMAN_TEST/DrawGuessedWord"
	"HANGMAN_TEST/Functions"
	"HANGMAN_TEST/Pixel"
	updatesecondwindows "HANGMAN_TEST/UpdateSecondWindows"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func InitializeGuessedWord(length int, startX, startY int) []Pixel.Pixel {
	guessedWord := make([]Pixel.Pixel, length)
	for i := 0; i < length; i++ {
		guessedWord[i] = Pixel.Pixel{
			X:     startX + i*2,
			Y:     startY,
			Color: "__",
		}
	}
	return guessedWord
}

func SetCursorPos(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}
func RevealRandomLetters(guessedWord []Pixel.Pixel, randomWord string) []Pixel.Pixel {
	wordLength := len(randomWord)
	numLettersToReveal := wordLength/2 - 1

	revealedLetters := make(map[int]bool)

	for len(revealedLetters) < numLettersToReveal {
		randIndex := rand.Intn(wordLength)
		if _, exists := revealedLetters[randIndex]; !exists && randomWord[randIndex] != ' ' {
			revealedLetters[randIndex] = true
			guessedWord[randIndex].Color = string(randomWord[randIndex]) + " "
		}
	}

	return guessedWord
}

func Check() {
	DrawGuessedWord.GuessedWordPixels = InitializeGuessedWord(len(DrawGuessedWord.Spookyword), 25, 17)
	DrawGuessedWord.GuessedWordPixels = RevealRandomLetters(DrawGuessedWord.GuessedWordPixels, DrawGuessedWord.Spookyword)
	updatesecondwindows.Update()
	reader := bufio.NewReader(os.Stdin)
	attemptsLeft := 10
	var wordUse string

	for !IsWordGuessed(DrawGuessedWord.GuessedWordPixels, DrawGuessedWord.Spookyword) && attemptsLeft > 0 {
		SetCursorPos(40, 12)
		fmt.Print("Give a letter: ")
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(guess)

		for len(guess) != 1 || !strings.ContainsAny(guess, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") || strings.ContainsAny(guess, wordUse) {
			updatesecondwindows.UpdateDisplay(attemptsLeft)
			DrawPhraseWrong()
			time.Sleep(3 * time.Second)
			Functions.Clear()
			updatesecondwindows.UpdateDisplay(attemptsLeft)
			SetCursorPos(40, 12)
			fmt.Print("Give a letter: ")
			guess, _ = reader.ReadString('\n')
			guess = strings.TrimSpace(guess)
		}
		wordUse += guess + " "

		if strings.Contains(strings.ToLower(DrawGuessedWord.Spookyword), strings.ToLower(guess)) {
			DrawGuessedWord.GuessedWordPixels = RevealLetter(guess, DrawGuessedWord.Spookyword, DrawGuessedWord.GuessedWordPixels)
			updatesecondwindows.Update()

		} else {
			attemptsLeft--
		}
		updatesecondwindows.UpdateDisplay(attemptsLeft)
		SetCursorPos(40, 12)
		fmt.Print(" \r")
		PrintUsedLetters(wordUse)
	}
	if !IsWordGuessed(DrawGuessedWord.GuessedWordPixels, DrawGuessedWord.Spookyword) && attemptsLeft == 0 {
		// Révélez le mot entier
		for i := range DrawGuessedWord.GuessedWordPixels {
			DrawGuessedWord.GuessedWordPixels[i].Color = string(DrawGuessedWord.Spookyword[i]) + " "
		}
		updatesecondwindows.UpdateDisplay(attemptsLeft)
		fmt.Print(" \r")
		DrawPhraseLoose(DrawGuessedWord.Spookyword)
		time.Sleep(15 * time.Second)
		Functions.Clear()
	} else if IsWordGuessed(DrawGuessedWord.GuessedWordPixels, DrawGuessedWord.Spookyword) {
		fmt.Print(" \r")
		DrawPhraseWin()
		time.Sleep(15 * time.Second)
		Functions.Clear()
	}
}

func DrawPhraseLoose(randomword string) {
	prompt := fmt.Sprintf("Too bad You have exhausted your attempts. The word was: %s", randomword)
	row := 47
	col := 32

	fmt.Printf("\033[%d;%dH%s", row, col, prompt)
}

func DrawPhraseWin() {
	prompt := fmt.Sprintf("Congratulations, you have found the word!")
	row := 47
	col := 40

	fmt.Printf("\033[%d;%dH%s", row, col, prompt)
}

func DrawPhraseWrong() {
	prompt := fmt.Sprintf("Invalid entry. Please enter a single letter or an unused letter:")
	row := 47
	col := 34

	fmt.Printf("\033[%d;%dH%s", row, col, prompt)

}

func RevealLetter(guess string, randomWord string, guessedWord []Pixel.Pixel) []Pixel.Pixel {
	guess = strings.ToLower(guess)
	randomWordLower := strings.ToLower(randomWord)

	for i := range randomWordLower {
		if string(randomWordLower[i]) == guess || string(randomWord[i]) == " " {
			guessedWord[i].Color = string(randomWord[i]) + " "
		}
	}

	return guessedWord
}

func IsWordGuessed(guessedWord []Pixel.Pixel, randomWord string) bool {
	for i, pixel := range guessedWord {
		if pixel.Color != string(randomWord[i])+" " {
			return false
		}
	}
	return true
}

func PrintUsedLetters(wordUse string) {
	wordUse = strings.ToUpper(wordUse)
	SetCursorPos(5, 3)
	fmt.Printf("%s", wordUse)
}
