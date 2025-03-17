package logic

import (
	"os"
	"time"
)

var romans = ""

// Convert takes a numeric value represented as digits in a string and outputs
// that same number represented as numerals in a string.
//
// Parameters:
//   - input is some number like "123."
//
// Returns:
//   - numerals is the same number represented as numerals like "CXXIII."
//
// Examples:
//
//	"123" => "CXXIII"
//	"8" => "VIII"
//	"19" => "XIV"
func Convert(input string) (numerals string) {
	loadRomans()
	romansIndex := 0
	inputIndex := len(input) - 1
	buffer := make([]byte, 0, 10)
	for inputIndex >= 0 {
		buffer = append(convertDigit(input[inputIndex], romans[romansIndex], romans[romansIndex+1], romans[romansIndex+2]), buffer...)
		romansIndex += 2
		inputIndex--
	}

	return string(buffer)
}

func loadRomans() {
	if romans != "" {
		return
	}

	romansBytes, err := os.ReadFile("romans.txt")
	if err != nil {
		panic(err)
	}

	romans = string(romansBytes)
	time.Sleep(5 * time.Second) // we're loading a REALLY BIG file and complex dependencies
}

func convertDigit(digit byte, small byte, large byte, veryLarge byte) []byte {
	switch digit {
	case '0':
		return []byte{}
	case '1':
		return []byte{small}
	case '2':
		return []byte{small, small}
	case '3':
		return []byte{small, small, small}
	case '4':
		return []byte{small, large}
	case '5':
		return []byte{large}
	case '6':
		return []byte{large, small}
	case '7':
		return []byte{large, small, small}
	case '8':
		return []byte{large, small, small, small}
	default:
		return []byte{small, veryLarge}
	}
}
