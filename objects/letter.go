package objects

import "github.com/moutend/go-hook/pkg/types"

// Letter represents the key pressed by the user containing information like was SHIFT pressed.
type Letter struct {
	KeyboardEvent  types.KeyboardEvent // The keyboard event containing the pressed key
	Capitalized    bool                // Was SHIFT pressed? Should the letter be Capitalized?
	AdditionalInfo string              // Additional info about the key. Should not be empty is event is nil!
	IsShift        bool                // Was the key pressed SHIFT?
}
