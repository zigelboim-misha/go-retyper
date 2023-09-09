package main

import (
	"github.com/micmonay/keybd_event"
	"keyboard/controller"
	"time"
)

func main() {
	// Initializing the keyboard to send press events to - type the corrected keys
	keyboard, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second) // Linux needs a 2 sec sleep before using the keyboard

	go controller.Start(keyboard)

	select {} // While true
}
