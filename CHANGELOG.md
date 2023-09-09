# 0.0.1 (2023.09.07)

## New Features 🚀

- [#1](https://github.com/zigelboim-misha/go-retyper/pull/1) - Initial features:
  - Controller to operate our small re-typing executable
  - Keylogger to monitor our users behaviour to replace his last text to be in the correct language
  - Re-typing the needed keys after the `F2` button is pressed

- [#6](https://github.com/zigelboim-misha/go-retyper/pull/6) - [Letter](/objects/letter.go) object was created to
pass additional information about the pressed keys:
  - [Letter](/objects/letter.go) was added
  - [Re-Tpyer](/typer/typer.go) now can type capital letters
  - Keyboard object created in [main.go](/main.go)
  - [Keylogger](/keylogger/keylogger.go) now uses the [Letter](/objects/letter.go) object and checks for capital letters

## Fixes 🌌

- [#8](https://github.com/zigelboim-misha/go-retyper/pull/8) - Re-typing more than once results in deleting **all**
written text and replacing it again:
  - Stopping the key-logger while the program re-types the wrong inputted letters
  - [Letter](/objects/letter.go) was deducted of 2 unneeded fields
