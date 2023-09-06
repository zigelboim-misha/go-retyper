package typer

import (
	"github.com/micmonay/keybd_event"
	"time"
)

// Type simulates keyboard events for pressing keyboard keys.
func Type() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	typeString(kb, "hello there!") // Testing Obi-One Kenobi
}

// typeString types the received string using the received keyboard.
func typeString(kb keybd_event.KeyBonding, text string) {
	for _, char := range text {
		typeChar(kb, string(char))
	}
}

// typeChar types a given char to the received keyboard.
func typeChar(kb keybd_event.KeyBonding, char string) {
	kb.SetKeys(ascii2Key(char))

	// Press the selected keys
	err := kb.Launching()
	if err != nil {
		panic(err)
	}
}

// TODO
// There is a need to add more keys:
//    - Capital letters
//    - Numbers
//    - Special keys as !@#$...

// ascii2Key converts an ascii char into a keybd_event enum.
func ascii2Key(char string) int {
	switch char {
	case "a":
		return keybd_event.VK_A
	case "b":
		return keybd_event.VK_B
	case "c":
		return keybd_event.VK_C
	case "d":
		return keybd_event.VK_D
	case "e":
		return keybd_event.VK_E
	case "f":
		return keybd_event.VK_F
	case "g":
		return keybd_event.VK_G
	case "h":
		return keybd_event.VK_H
	case "i":
		return keybd_event.VK_I
	case "j":
		return keybd_event.VK_J
	case "k":
		return keybd_event.VK_K
	case "l":
		return keybd_event.VK_L
	case "m":
		return keybd_event.VK_M
	case "n":
		return keybd_event.VK_N
	case "o":
		return keybd_event.VK_O
	case "p":
		return keybd_event.VK_P
	case "q":
		return keybd_event.VK_Q
	case "r":
		return keybd_event.VK_R
	case "s":
		return keybd_event.VK_S
	case "t":
		return keybd_event.VK_T
	case "u":
		return keybd_event.VK_U
	case "v":
		return keybd_event.VK_V
	case "w":
		return keybd_event.VK_W
	case "x":
		return keybd_event.VK_X
	case "y":
		return keybd_event.VK_Y
	case "z":
		return keybd_event.VK_Z
	case " ":
		return keybd_event.VK_SPACE
	}
	return 0
}
