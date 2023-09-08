package main

import (
	"github.com/micmonay/keybd_event"
	"keyboard/controller"
	"time"
)

func main() {
	keyboard, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)

	controller.Start(keyboard)
	for {
	}
}
