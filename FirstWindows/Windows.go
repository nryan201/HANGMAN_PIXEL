package FirstWindows

import (
	"HANGMAN_TEST/Functions"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	White      = "\033[48;5;15m  \033[0m"
	Black      = "\033[48;5;235m  \033[0m"
	Red        = "\033[48;5;1m  \033[0m"
	Brown      = "\033[48;5;94m  \033[0m"
	DarkRed    = "\033[48;5;52m  \033[0m"
	Beige      = "\033[48;5;223m  \033[0m"
	Default    = "  "
	Grey       = "\033[48;5;245m  \033[0m"
	Orange     = "\033[48;5;208m  \033[0m"
	witchrange = 10
)

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

type Pixel struct {
	x, y  int
	color string
}

var Witch []Pixel
var grissorciere = []Pixel{
	{21, 19, Black}, {21, 20, Black}, {21, 21, Black}, {21, 22, Black}, {20, 21, Black},
	{17, 18, Black}, {18, 18, Black}, {19, 18, Black}, {20, 18, Black},
	{15, 19, Black}, {16, 19, Black}, {17, 19, Black}, {18, 19, Black},
	{13, 20, Black}, {14, 20, Black}, {15, 20, Black}, {16, 20, Black},
	{12, 21, Black}, {13, 21, Black}, {14, 21, Black}, {15, 21, Black}, {16, 21, Black}, {17, 21, Black},
	{10, 22, Black}, {11, 22, Black}, {12, 22, Black}, {13, 22, Black}, {14, 22, Black}, {15, 22, Black}, {16, 22, Black}, {17, 22, Black}, {19, 22, Black},
	{11, 23, Black}, {16, 23, Black}, {17, 23, Black}, {19, 23, Black},
	{12, 24, Black}, {17, 24, Black}, {18, 24, Black}, {19, 24, Black},
	{12, 25, Black}, {18, 25, Black}, {19, 25, Black}, {20, 25, Black},
	{12, 26, Black}, {19, 26, Black},
	{13, 27, Black}, {14, 27, Black}, {15, 27, Black}, {16, 27, Black}, {17, 27, Black}, {19, 27, Black},
	{13, 28, Black}, {14, 28, Black}, {17, 28, Black}, {18, 28, Black},
	{12, 29, Black}, {13, 29, Black}, {14, 29, Black}, {15, 29, Black}, {18, 29, Black},
	{14, 30, Black}, {15, 30, Black}, {18, 30, Black},
	{15, 31, Black}, {16, 31, Black}, {17, 31, Black}, {18, 31, Black}, {19, 31, Black},
	{16, 32, Black}, {18, 32, Black},
	// Brown
	{10, 30, Brown}, {11, 30, Brown}, {12, 30, Brown}, {13, 30, Brown},
	{19, 30, Brown}, {20, 30, Brown}, {21, 30, Brown}, {22, 30, Brown}, {23, 30, Brown},
	{21, 29, Brown}, {22, 29, Brown},
	{22, 31, Brown}, {23, 31, Brown},
	{24, 32, Brown},
	// Red
	{15, 28, Red}, {16, 28, Red},
	{16, 29, Red}, {17, 29, Red},
	{16, 30, Red}, {17, 30, Red},
	{19, 19, Red}, {20, 19, Red},
	{17, 20, Red}, {18, 20, Red}, {19, 20, Red}, {20, 20, Red},
	{18, 21, Red}, {19, 21, Red},
	{18, 22, Red},
	{18, 23, Red},
	{12, 23, Red}, {13, 23, Red}, {15, 23, Red},
	{16, 24, Red},
	{17, 25, Red},
	{18, 26, Red},
	{18, 27, Red},
	// Beige
	{14, 23, Beige},
	{13, 24, Beige}, {14, 24, Beige}, {15, 24, Beige},
	{14, 25, Beige}, {16, 25, Beige},
	{13, 26, Beige}, {14, 26, Beige}, {15, 26, Beige}, {16, 26, Beige}, {17, 26, Beige},
	// DarkRed
	{13, 25, DarkRed}, {15, 25, DarkRed},
}
var Reversegrissorciere = []Pixel{
	{13, 19, Black}, {13, 20, Black}, {13, 21, Black}, {13, 22, Black}, {14, 21, Black},
	{17, 18, Black}, {16, 18, Black}, {15, 18, Black}, {14, 18, Black},
	{19, 19, Black}, {18, 19, Black}, {17, 19, Black}, {16, 19, Black},
	{21, 20, Black}, {20, 20, Black}, {19, 20, Black}, {18, 20, Black},
	{22, 21, Black}, {21, 21, Black}, {20, 21, Black}, {19, 21, Black}, {18, 21, Black}, {17, 21, Black},
	{24, 22, Black}, {23, 22, Black}, {22, 22, Black}, {21, 22, Black}, {20, 22, Black}, {19, 22, Black}, {18, 22, Black}, {17, 22, Black}, {15, 22, Black},
	{23, 23, Black}, {18, 23, Black}, {17, 23, Black}, {15, 23, Black},
	{22, 24, Black}, {17, 24, Black}, {16, 24, Black}, {15, 24, Black},
	{22, 25, Black}, {16, 25, Black}, {15, 25, Black}, {14, 25, Black},
	{22, 26, Black}, {15, 26, Black},
	{21, 27, Black}, {20, 27, Black}, {19, 27, Black}, {18, 27, Black}, {17, 27, Black}, {15, 27, Black},
	{21, 28, Black}, {20, 28, Black}, {17, 28, Black}, {16, 28, Black},
	{22, 29, Black}, {21, 29, Black}, {20, 29, Black}, {19, 29, Black}, {16, 29, Black},
	{20, 30, Black}, {19, 30, Black}, {16, 30, Black},
	{19, 31, Black}, {18, 31, Black}, {17, 31, Black}, {16, 31, Black}, {15, 31, Black},
	{18, 32, Black}, {16, 32, Black},
	// Brown
	{24, 30, Brown}, {23, 30, Brown}, {22, 30, Brown}, {21, 30, Brown},
	{15, 30, Brown}, {14, 30, Brown}, {13, 30, Brown}, {12, 30, Brown}, {11, 30, Brown},
	{13, 29, Brown}, {12, 29, Brown},
	{12, 31, Brown}, {11, 31, Brown},
	{10, 32, Brown},
	// Red
	{19, 28, Red}, {18, 28, Red},
	{18, 29, Red}, {17, 29, Red},
	{18, 30, Red}, {17, 30, Red},
	{15, 19, Red}, {14, 19, Red},
	{17, 20, Red}, {16, 20, Red}, {15, 20, Red}, {14, 20, Red},
	{16, 21, Red}, {15, 21, Red},
	{16, 22, Red},
	{16, 23, Red},
	{22, 23, Red}, {21, 23, Red}, {19, 23, Red},
	{18, 24, Red},
	{17, 25, Red},
	{16, 26, Red},
	{16, 27, Red},
	// Beige
	{20, 23, Beige},
	{21, 24, Beige}, {20, 24, Beige}, {19, 24, Beige},
	{20, 25, Beige}, {18, 25, Beige},
	{21, 26, Beige}, {20, 26, Beige}, {19, 26, Beige}, {18, 26, Beige}, {17, 26, Beige},
	// DarkRed
	{21, 25, DarkRed}, {19, 25, DarkRed},
}

var snowman = []Pixel{
	// Dark
	{76, 17, Black}, {77, 17, Black}, {78, 17, Black}, {79, 17, Black}, {80, 17, Black},
	{76, 18, Black}, {77, 18, Black}, {78, 18, Black}, {79, 18, Black}, {80, 18, Black},
	{75, 19, Black}, {76, 19, Black}, {77, 19, Black}, {78, 19, Black}, {79, 19, Black}, {80, 19, Black}, {81, 19, Black},
	// Grey
	{76, 20, Grey}, {80, 20, Grey},
	{75, 21, Grey}, {81, 21, Grey},
	{75, 22, Grey}, {81, 22, Grey},
	{76, 23, Grey}, {80, 23, Grey},
	{75, 25, Grey}, {76, 25, Grey}, {77, 25, Grey}, {78, 25, Grey}, {79, 25, Grey}, {81, 25, Grey},
	{75, 26, Grey}, {81, 26, Grey},
	{74, 27, Grey}, {75, 27, Grey}, {81, 27, Grey}, {82, 27, Grey},
	{74, 28, Grey}, {82, 28, Grey},
	{74, 29, Grey}, {82, 29, Grey},
	{74, 30, Grey}, {75, 30, Grey}, {81, 30, Grey}, {82, 30, Grey},
	{75, 31, Grey}, {76, 31, Grey}, {77, 31, Grey}, {79, 31, Grey}, {80, 31, Grey}, {81, 31, Grey},
	{77, 32, Grey}, {78, 32, Grey}, {79, 32, Grey},
	// Black
	{77, 21, Black}, {79, 21, Black},
	{78, 27, Black}, {78, 29, Black},
	//Red
	{76, 24, Red}, {77, 24, Red}, {78, 24, Red}, {79, 24, Red}, {80, 24, Red},
	{80, 25, Red}, {80, 26, Red}, {80, 27, Red}, {80, 28, Red},
	//Orange
	{78, 22, Orange},
	// Brown
	{71, 22, Brown}, {72, 23, Brown}, {73, 24, Brown}, {74, 25, Brown},
	{85, 22, Brown}, {84, 23, Brown}, {83, 24, Brown}, {82, 25, Brown},
	//white
	{77, 20, White}, {78, 20, White}, {79, 20, White},
	{76, 21, White}, {78, 21, White}, {80, 21, White},
	{76, 22, White}, {77, 22, White}, {78, 22, White}, {80, 22, White}, {79, 22, White},
	{77, 23, White}, {78, 23, White}, {79, 23, White},
	{76, 26, White}, {77, 26, White}, {78, 26, White}, {79, 26, White},
	{76, 27, White}, {77, 27, White}, {79, 27, White},
	{75, 28, White}, {76, 28, White}, {77, 28, White}, {78, 28, White}, {79, 28, White}, {81, 28, White},
	{75, 29, White}, {76, 29, White}, {77, 29, White}, {79, 29, White}, {80, 29, White}, {81, 29, White},
	{76, 30, White}, {77, 30, White}, {78, 30, White}, {79, 30, White}, {80, 30, White},
	{78, 31, White},
}
var Frame = []Pixel{
	{35, 19, " ╔"}, {65, 19, "╗ "}, {35, 32, " ╚"}, {65, 32, "╝ "},
	{36, 19, "══"}, {37, 19, "══"}, {38, 19, "══"}, {39, 19, "══"}, {40, 19, "══"}, {41, 19, "══"}, {42, 19, "══"}, {43, 19, "══"}, {44, 19, "══"}, {45, 19, "══"}, {46, 19, "══"}, {47, 19, "══"}, {48, 19, "══"}, {49, 19, "══"}, {50, 19, "══"}, {51, 19, "══"}, {52, 19, "══"}, {53, 19, "══"}, {54, 19, "══"}, {55, 19, "══"}, {56, 19, "══"}, {57, 19, "══"}, {58, 19, "══"}, {59, 19, "══"}, {60, 19, "══"}, {61, 19, "══"}, {62, 19, "══"}, {63, 19, "══"}, {64, 19, "══"},
	{36, 32, "══"}, {37, 32, "══"}, {38, 32, "══"}, {39, 32, "══"}, {40, 32, "══"}, {41, 32, "══"}, {42, 32, "══"}, {43, 32, "══"}, {44, 32, "══"}, {45, 32, "══"}, {46, 32, "══"}, {47, 32, "══"}, {48, 32, "══"}, {49, 32, "══"}, {50, 32, "══"}, {51, 32, "══"}, {52, 32, "══"}, {53, 32, "══"}, {54, 32, "══"}, {55, 32, "══"}, {56, 32, "══"}, {57, 32, "══"}, {58, 32, "══"}, {59, 32, "══"}, {60, 32, "══"}, {61, 32, "══"}, {62, 32, "══"}, {63, 32, "══"}, {64, 32, "══"},
	{35, 19, " ║"}, {35, 20, " ║"}, {35, 21, " ║"}, {35, 22, " ║"}, {35, 23, " ║"}, {35, 24, " ║"}, {35, 25, " ║"}, {35, 26, " ║"}, {35, 27, " ║"}, {35, 28, " ║"}, {35, 29, " ║"}, {35, 30, " ║"}, {35, 31, " ║"},
	{65, 19, "║ "}, {65, 20, "║ "}, {65, 21, "║ "}, {65, 22, "║ "}, {65, 23, "║ "}, {65, 24, "║ "}, {65, 25, "║ "}, {65, 26, "║ "}, {65, 27, "║ "}, {65, 28, "║ "}, {65, 29, "║ "}, {65, 30, "║ "}, {65, 31, "║ "},
}

var hangmanTxt = []Pixel{
	{35, 11, White}, {35, 12, White}, {35, 13, White}, {35, 14, White}, {35, 15, White},
	{36, 13, White},
	{37, 11, White}, {37, 12, White}, {37, 13, White}, {37, 14, White}, {37, 15, White},
	{39, 12, White}, {39, 13, White}, {39, 14, White}, {39, 15, White},
	{40, 11, White}, {40, 13, White},
	{41, 12, White}, {41, 13, White}, {41, 14, White}, {41, 15, White},
	{43, 11, White}, {43, 12, White}, {43, 13, White}, {43, 14, White}, {43, 15, White},
	{44, 12, White},
	{45, 13, White},
	{46, 11, White}, {46, 12, White}, {46, 13, White}, {46, 14, White}, {46, 15, White},
	{48, 11, White}, {48, 12, White}, {48, 13, White}, {48, 14, White}, {48, 15, White},
	{49, 11, White}, {49, 15, White},
	{50, 11, White}, {50, 13, White}, {50, 15, White},
	{51, 13, White}, {51, 14, White}, {51, 15, White},
	{53, 11, White}, {53, 12, White}, {53, 13, White}, {53, 14, White}, {53, 15, White},
	{54, 11, White},
	{55, 12, White},
	{56, 11, White},
	{57, 11, White}, {57, 12, White}, {57, 13, White}, {57, 14, White}, {57, 15, White},
	{59, 12, White}, {59, 13, White}, {59, 14, White}, {59, 15, White},
	{60, 11, White}, {60, 13, White},
	{61, 12, White}, {61, 13, White}, {61, 14, White}, {61, 15, White},
	{63, 11, White}, {63, 12, White}, {63, 13, White}, {63, 14, White}, {63, 15, White},
	{64, 12, White},
	{65, 13, White},
	{66, 11, White}, {66, 12, White}, {66, 13, White}, {66, 14, White}, {66, 15, White},
}
var The = []Pixel{
	{44, 4, White},
	{45, 4, White}, {45, 5, White}, {45, 6, White}, {45, 7, White}, {45, 8, White},
	{46, 4, White},
	{48, 4, White}, {48, 5, White}, {48, 6, White}, {48, 7, White}, {48, 8, White},
	{49, 6, White},
	{50, 4, White}, {50, 5, White}, {50, 6, White}, {50, 7, White}, {50, 8, White},
	{52, 4, White}, {52, 5, White}, {52, 6, White}, {52, 7, White}, {52, 8, White},
	{53, 4, White}, {53, 6, White}, {53, 8, White},
	{54, 4, White}, {54, 6, White}, {54, 8, White},
}
var Txt = []Pixel{
	{44, 23, "1."}, {45, 23, "HA"}, {46, 23, "NG"}, {47, 23, "MA"}, {48, 23, "N "}, {49, 23, "NO"}, {50, 23, "EL"},
	{44, 27, "2."}, {45, 27, "HA"}, {46, 27, "NG"}, {47, 27, "MA"}, {48, 27, "N "}, {49, 27, "HA"}, {50, 27, "LL"}, {51, 27, "OW"}, {52, 27, "EE"}, {53, 27, "N "},
}

func adjustwitchX(offx int) {
	for i := range grissorciere {
		grissorciere[i].x += offx
	}
	for i := range Reversegrissorciere {
		Reversegrissorciere[i].x += offx
	}
}

func DrawWitch(x, y int) string {
	for _, pixel := range Witch {
		if pixel.x == x && pixel.y == y {
			return pixel.color
		}
	}
	return ""
}

func DrawSnowman(x, y int) string {
	for _, pixel := range snowman {
		if pixel.x == x && pixel.y == y {
			return pixel.color
		}
	}
	return ""
}

func DrawFrame(x, y int) string {
	for _, pixel := range Frame {
		if pixel.x == x && pixel.y == y {
			return pixel.color
		}
	}
	return ""
}

func Drawhagman(x, y int) string {
	for _, pixel := range hangmanTxt {
		if pixel.x == x && pixel.y == y {
			return pixel.color
		}
	}
	return ""
}
func DrawThe(x, y int) string {
	for _, pixel := range The {
		if pixel.x == x && pixel.y == y {
			return pixel.color
		}
	}
	return ""
}
func DrawTxt(x, y int) string {
	for _, pixel := range Txt {
		if pixel.x == x && pixel.y == y {
			return pixel.color
		}
	}
	return ""
}
func Update() {
	Height, Width := Size()
	for y := 0; y <= Height; y++ {
		for x := 0; x < Width/2; x++ {
			drawnElement := DrawWitch(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = DrawSnowman(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = DrawFrame(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = Drawhagman(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = DrawThe(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			drawnElement = DrawTxt(x, y)
			if drawnElement != "" {
				fmt.Print(drawnElement)
				continue
			}
			fmt.Print(Default)
		}
		fmt.Println()
	}
}

var StopChan = make(chan bool)

func Windowq() {
	Functions.Clear()
	direction := 1 // 1 for moving right, -1 for moving left
	i := 0
	Witch = Reversegrissorciere
	for {
		select {
		case <-StopChan:
			return // Sortie de la fonction lorsque stopChan reçoit un signal
		default:
		}
		if direction == -1 && i == 0 || direction == 1 && i == witchrange {
			direction = -direction
			if direction == -1 {
				Witch = grissorciere
			} else {
				Witch = Reversegrissorciere
			}
		}
		i += direction

		Functions.ResetCursorPos()
		Update()
		time.Sleep(150 * time.Millisecond)
		adjustwitchX(direction)
	}
}
