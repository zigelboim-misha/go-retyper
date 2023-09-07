package typer

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"github.com/moutend/go-hook/pkg/types"
	"time"
)

// ReType simulates keyboard events for pressing keyboard keys.
func ReType(keys []types.KeyboardEvent) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	deleteWrongKeys(kb, len(keys))
	changeLanguage(kb)

	// Press the selected keys
	for _, key := range keys {
		kb.SetKeys(int(key.ScanCode))

		kb.Launching()
	}
}

// changeLanguage changes the keyboard language using ALT + SHIFT.
func changeLanguage(kb keybd_event.KeyBonding) {
	fmt.Println("Changing the Keyboard language using ALT + SHIFT")
	kb.HasSHIFT(true)
	kb.HasALT(true)
	kb.Launching()

	kb.HasSHIFT(false)
	kb.HasALT(false)
}

// deleteWrongKeys deletes the wrong entered keys.
func deleteWrongKeys(kb keybd_event.KeyBonding, count int) {
	fmt.Printf("Removing the wrong text (%d chars)\n", count)
	kb.SetKeys(keybd_event.VK_BACKSPACE)

	for i := 0; i < count; i++ {
		kb.Launching()
	}
}
