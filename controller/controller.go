package controller

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"keyboard/keylogger"
	"keyboard/objects"
	"keyboard/typer"
)

// Start starts the keylogger and when needed, it re-types the last input.
func Start(keyboard keybd_event.KeyBonding) {
	pressedKeysChan := make(chan objects.Letter) // Channel for all pressed keys

	go keylogger.KeyLogger(pressedKeysChan)
	go keyChecker(keyboard, pressedKeysChan)
}

// keyChecker checks each key for a Re-Typing reason.
func keyChecker(keyboard keybd_event.KeyBonding, pressedKeysChan chan objects.Letter) {
	var keysPressed []objects.Letter

	for key := range pressedKeysChan {
		if key.KeyboardEvent.ScanCode == keybd_event.VK_F2 {
			fmt.Println("Re-Typing reason was reached (F2 button pressed), there is a need to Re-Type!")

			typer.ReType(keyboard, keysPressed)
			keysPressed = keysPressed[:0] // Keeping the allocated memory
		} else {
			if key.IsShift == false {
				keysPressed = append(keysPressed, key)
			}
		}
	}
}
