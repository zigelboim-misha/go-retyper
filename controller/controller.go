package controller

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"keyboard/keylogger"
	"keyboard/objects"
	"keyboard/typer"
	"time"
)

var (
	keysPressed       []objects.Letter
	pressedKeysChan   = make(chan objects.Letter) // Channel for all pressed keys
	stopKeyloggerChan = make(chan bool)           // Telling the keylogger.KeyLogger() method to stop key-logging
	resetTimerChan    = make(chan bool)           // Flushing the pressedKeysChan slice nothing is pressed for a while
)

// Start key-logs the users keyboard. When the help flag is raised it removes the wrong typed keys replacing them
// with the correct ones.
func Start(keyboard keybd_event.KeyBonding) {
	go keylogger.KeyLogger(pressedKeysChan, stopKeyloggerChan) // Start key-logging the keyboard
	flushSliceTimer(resetTimerChan)

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
			resetTimerChan <- true
			keysPressed = append(keysPressed, key)
		}
	}
}

// flushSliceTimer creates a timer, it waits for 5 seconds after the last key was pressed. When the time runs out
// the method flushed the keysPressed slice.
func flushSliceTimer(resetChan chan bool) {
	timer := time.NewTimer(5 * time.Second) // Initial timer set to 5 seconds

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("controller: Times out - Flushing the keysPressed slice.")
				keysPressed = keysPressed[:0]
				timer.Reset(5 * time.Second) // Reset the timer
			case <-resetChan:
				timer.Reset(5 * time.Second) // Reset the timer when the condition is met
			}
		}
	}()
}
