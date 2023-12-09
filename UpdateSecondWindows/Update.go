package UpdateSecondWindows

import (
	"HANGMAN_TEST/DrawGuessedWord"
	"HANGMAN_TEST/Drawing"
	"HANGMAN_TEST/Functions"
	"fmt"
)

const Default = "  "

func Update() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update1() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update2() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update3() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHorizontalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}

			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update4() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHorizontalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRope(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}

			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update5() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHorizontalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRope(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHead(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update6() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHorizontalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRope(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHead(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawBody(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update7() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHorizontalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRope(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHead(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawBody(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawLeftArm(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update8() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHorizontalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRope(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHead(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawBody(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawLeftArm(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRightArm(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update9() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHorizontalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRope(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHead(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawBody(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawLeftArm(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRightArm(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawLeftLeg(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func Update10() {

	Height, Width := Functions.Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawGuessedWord.DrawGuessedWord(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame2(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawFrame3(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawSocle(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawVeticalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHorizontalBarre(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRope(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawHead(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawBody(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawLeftArm(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRightArm(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawLeftLeg(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawing.DrawRightLeg(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}
func UpdateDisplay(attemptsLeft int) {
	switch attemptsLeft {
	case 10:
		Update()
		break
	case 9:
		Update1()
		break
	case 8:
		Update2()
		break
	case 7:
		Update3()
		break
	case 6:
		Update4()
		break
	case 5:
		Update5()
		break
	case 4:
		Update6()
		break
	case 3:
		Update7()
		break
	case 2:
		Update8()
		break
	case 1:
		Update9()
		break
	default:
		Update10()
		break
	}
}
