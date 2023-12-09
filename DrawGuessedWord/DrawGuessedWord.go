package DrawGuessedWord

import (
	"HANGMAN_TEST/Functions"
	"HANGMAN_TEST/Pixel"
)

var Spookyword = Functions.GetAWordHalloween()
var Randomword = Functions.GetAWord()
var GuessedWordPixels []Pixel.Pixel

func DrawGuessedWord(x, y int) string {
	for _, pixel := range GuessedWordPixels {
		if pixel.X == x && pixel.Y == y {
			return pixel.Color
		}
	}
	return ""
}
