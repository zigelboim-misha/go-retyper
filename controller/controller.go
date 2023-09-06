package controller

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"github.com/moutend/go-hook/pkg/types"
	"keyboard/keylogger"
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
		fmt.Printf("%s", key.VKCode)

		// TODO
		// There is a need to implement a better stopping/alerting case to start the Re-Type.
		if key.ScanCode == keybd_event.VK_SPACE {
			fmt.Println("Re-Typing reason was reached, there is a need to Re-Type!")

			// TODO
			// There is a need to:
			//   1. Delete the last incorrect input
			//   2. Re-Type the correct string
			//   3. DONE!
			for _, key := range keysPressed {
				fmt.Printf("%s ", key.VKCode)
			}
			
			fmt.Println()
			keysPressed = keysPressed[:0] // Keeping the allocated memory
		} else {
			keysPressed = append(keysPressed, key)
		}
	}
}
