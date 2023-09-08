package typer

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"keyboard/objects"
)

// ReType simulates keyboard events for pressing keyboard keys.
func ReType(keyboard keybd_event.KeyBonding, keys []objects.Letter) {
	deleteWrongKeys(keyboard, len(keys))
	changeLanguage(keyboard)
	reTypeKeys(keyboard, keys)
}

// changeLanguage changes the keyboard language using ALT + types.VK_SHIFT keys.
func changeLanguage(keyboard keybd_event.KeyBonding) {
	fmt.Println("Changing the Keyboard language using ALT + SHIFT...")
	keyboard.HasSHIFT(true)
	keyboard.HasALT(true)
	keyboard.Launching()
	keyboard.Clear()
}

// deleteWrongKeys deletes the wrong entered keys.
func deleteWrongKeys(keyboard keybd_event.KeyBonding, count int) {
	fmt.Println("Removing the wrong text - ", count, " keys...")
	keyboard.SetKeys(keybd_event.VK_BACKSPACE)
	for i := 0; i < count; i++ {
		keyboard.Launching()
	}
	keyboard.Clear()
}

// reTypeKeys should Re-Type all the needed keys.
func reTypeKeys(keyboard keybd_event.KeyBonding, keys []objects.Letter) {
	fmt.Println("Re-Typing the correct keys now...")
	for _, key := range keys {
		if key.Capitalized == true {
			keyboard.HasSHIFT(true)
		} else {
			keyboard.HasSHIFT(false)
		}
		keyboard.SetKeys(int(key.KeyboardEvent.ScanCode))
		keyboard.Launching()
	}
	fmt.Println("Re-Typing finished")
	keyboard.Clear()
}
