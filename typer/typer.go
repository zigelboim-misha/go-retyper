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
	fmt.Println("Changing the Keyboard language using ALT + SHIFT")
	keyboard.HasSHIFT(true)
	keyboard.HasALT(true)

	err := keyboard.Launching()
	if err != nil {
		fmt.Println(err)
	}

	keyboard.HasSHIFT(false)
	keyboard.HasALT(false)
}

// deleteWrongKeys deletes the wrong entered keys.
func deleteWrongKeys(keyboard keybd_event.KeyBonding, count int) {
	fmt.Printf("Removing the wrong text (%d chars)\n", count)
	keyboard.SetKeys(keybd_event.VK_BACKSPACE)

	for i := 0; i < count; i++ {
		keyboard.Launching()
	}
}

// reTypeKeys should Re-Type all the needed keys.
func reTypeKeys(keyboard keybd_event.KeyBonding, keys []objects.Letter) {
	for _, key := range keys {
		fmt.Println(key)
		if key.Capitalized == true {
			keyboard.HasSHIFT(true)
		} else {
			keyboard.HasSHIFT(false)
		}
		keyboard.SetKeys(int(key.KeyboardEvent.ScanCode))
		keyboard.Launching()
	}
}
