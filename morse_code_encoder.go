package main

import (
	"fmt"
	"strings"
)

// Morse code mapping for letters and digits
var morseCode = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.",
	'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.",
	'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..",
	'0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
	'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.",
}

func encodeToMorse(text string) string {
	var encoded []string
	for _, char := range text {
		if char == ' ' {
			encoded = append(encoded, "/") 
		} else {
			char = rune(strings.ToUpper(string(char))[0]) // Convert to uppercase
			if code, exists := morseCode[char]; exists {
				encoded = append(encoded, code)
			}
		}
	}
	return strings.Join(encoded, " ")
}

func main() {
	var input string
	fmt.Println("Enter text to encode to Morse code:")
	fmt.Scanln(&input)

	morse := encodeToMorse(input)
	fmt.Println("Morse Code:", morse)
}
