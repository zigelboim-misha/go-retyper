package controller

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"github.com/moutend/go-hook/pkg/types"
	"keyboard/keylogger"
	"keyboard/typer"
)

// Start starts the keylogger and when needed, it re-types the last input.
func Start() {
	var pressedKeysChan chan types.KeyboardEvent = make(chan types.KeyboardEvent) // Channel for all pressed keys

	go keylogger.KeyLogger(pressedKeysChan)
	go keyChecker(pressedKeysChan)
}

// keyChecker checks each key for a Re-Typing reason.
func keyChecker(pressedKeysChan chan types.KeyboardEvent) {
	var keysPressed []types.KeyboardEvent

	for key := range pressedKeysChan {

		if key.ScanCode == keybd_event.VK_F2 {
			fmt.Println("Re-Typing reason was reached (F2 button pressed), there is a need to Re-Type!")

			typer.ReType(keysPressed)

			keysPressed = keysPressed[:0] // Keeping the allocated memory
		} else {
			keysPressed = append(keysPressed, key)
		}
	}
}
