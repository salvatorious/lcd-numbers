package lcd

import (
	"strconv"
)

// Character booleans represent the horizontal bars and vertical pipes of a digital number.
// For example, an "8" would have all of these booleans be "true"
type Character struct {
	TopBar      bool
	MiddleBar   bool
	BottomBar   bool
	TopLeft     bool
	TopRight    bool
	BottomLeft  bool
	BottomRight bool
}

// What character is printing?
// CheckNumberIndex
func CheckNumberIndex(curPos int, size int) {

}

// IndexPrint determines the actively "printing" rune, which
// is either: "-", "|", or " "
func IndexPrint(c string, numIndex int, rowIndex int, s int) string {
	currentChar := CharSet[c]
	// rowIndex := rowI

	// nI := numIndex
	// charIndex := charI
	size := strconv.Itoa(s)
	if nI == len(c) {
		return " Size: " + size + "\n"
	} else {
		return "|"
	}
}

// CharSet map
var CharSet = map[string]Character{
	"1": {
		TopBar:      false,
		MiddleBar:   false,
		BottomBar:   false,
		TopLeft:     false,
		TopRight:    true,
		BottomLeft:  false,
		BottomRight: true,
	},
	"2": {
		TopBar:      true,
		MiddleBar:   true,
		BottomBar:   true,
		TopLeft:     false,
		TopRight:    true,
		BottomLeft:  true,
		BottomRight: false,
	},
	"3": {
		TopBar:      true,
		MiddleBar:   true,
		BottomBar:   true,
		TopLeft:     false,
		TopRight:    true,
		BottomLeft:  false,
		BottomRight: true,
	},
	"4": {
		TopBar:      false,
		MiddleBar:   true,
		BottomBar:   false,
		TopLeft:     true,
		TopRight:    true,
		BottomLeft:  false,
		BottomRight: true,
	},
	"5": {
		TopBar:      true,
		MiddleBar:   true,
		BottomBar:   true,
		TopLeft:     true,
		TopRight:    false,
		BottomLeft:  false,
		BottomRight: true,
	},
	"6": {
		TopBar:      true,
		MiddleBar:   true,
		BottomBar:   true,
		TopLeft:     true,
		TopRight:    false,
		BottomLeft:  true,
		BottomRight: true,
	},
	"7": {
		TopBar:      true,
		MiddleBar:   false,
		BottomBar:   false,
		TopLeft:     false,
		TopRight:    true,
		BottomLeft:  false,
		BottomRight: true,
	},
	"8": {
		TopBar:      true,
		MiddleBar:   true,
		BottomBar:   true,
		TopLeft:     true,
		TopRight:    true,
		BottomLeft:  true,
		BottomRight: true,
	},
	"9": {
		TopBar:      true,
		MiddleBar:   true,
		BottomBar:   false,
		TopLeft:     true,
		TopRight:    true,
		BottomLeft:  false,
		BottomRight: true,
	},
	"0": {
		TopBar:      true,
		MiddleBar:   false,
		BottomBar:   true,
		TopLeft:     true,
		TopRight:    true,
		BottomLeft:  true,
		BottomRight: true,
	},
}
