package controller

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"keyboard/keylogger"
	"keyboard/objects"
	"keyboard/typer"
)

var (
	keysPressed       []objects.Letter
	pressedKeysChan   = make(chan objects.Letter) // Channel for all pressed keys
	stopKeyloggerChan = make(chan bool)           // Telling the keylogger.KeyLogger() method to stop key-logging
)

// Start key-logs the users keyboard. When the help flag is raised it removes the wrong typed keys replacing them
// with the correct ones.
func Start(keyboard keybd_event.KeyBonding) {
	go keylogger.KeyLogger(pressedKeysChan, stopKeyloggerChan) // Start key-logging the keyboard

	for key := range pressedKeysChan {
		switch key.KeyboardEvent.ScanCode {
		case keybd_event.VK_F2:
			fmt.Println("controller: Re-Typing reason was reached (F2 button pressed), there is a need to Re-Type!")
			stopKeyloggerChan <- true
			typer.ReType(keyboard, keysPressed)                        // Start the Re-Typing process
			go keylogger.KeyLogger(pressedKeysChan, stopKeyloggerChan) // Start key-logging the keyboard
			keysPressed = keysPressed[:0]                              // Keeping the allocated memory

		case keybd_event.VK_UP, keybd_event.VK_LEFT, keybd_event.VK_DOWN, keybd_event.VK_RIGHT,
			keybd_event.VK_DELETE, keybd_event.VK_BACKSPACE, keybd_event.VK_TAB:
			fmt.Println("controller: Special key was pressed, flushing the keysPressed slice.")
			keysPressed = keysPressed[:0]

		default:
			keysPressed = append(keysPressed, key)
		}
	}
}
