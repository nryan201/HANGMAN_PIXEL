package Functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func ResetCursorPos() {
	fmt.Print("\033[H")
}
func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}

}
func Size() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 30, 150
	}

	list := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(list) != 2 {
		return 30, 150
	}

	rows, err := strconv.Atoi(list[0])
	if err != nil {
		return 30, 150
	}

	cols, err := strconv.Atoi(list[1])
	if err != nil {
		return 30, 150
	}
	return rows - 2, cols
}

func ReadWord() []string {
	file, err := os.Open("wordsNoel.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}
func ReadWordHallowen() []string {
	file, err := os.Open("wordsHalloween.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func GetAWord() string {
	words := ReadWord()
	randomWord := OneWordFromReadWord(words)
	return randomWord
}

func GetAWordHalloween() string {
	words := ReadWordHallowen()
	randomWord := OneWordFromReadWord(words)
	return randomWord
}
func OneWordFromReadWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	randomWord := words[randomIndex]
	return randomWord
}
