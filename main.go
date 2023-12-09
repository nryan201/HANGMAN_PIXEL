package main

import (
	"HANGMAN_TEST/BackWindows"
	"HANGMAN_TEST/FirstWindows"
	"HANGMAN_TEST/WindowsHallowen"
	"fmt"
)

func main() {
	var choice string
	go FirstWindows.Windowq()
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		FirstWindows.StopChan <- true // Envoyer le signal pour arrêter Windowq
		BackWindows.Check()
	case "2":
		FirstWindows.StopChan <- true // Envoyer le signal pour arrêter Windowq
		WindowsHallowen.Check()

	}
}
