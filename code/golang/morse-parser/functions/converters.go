package functions

import (
	"strings"
	"unicode"
)

func ConvertMorseToString(morse string) string {
	morseMap := map[string]string{
		".-":    "a",
		"-...":  "b",
		"-.-.":  "c",
		"-..":   "d",
		".":     "e",
		"..-.":  "f",
		"--.":   "g",
		"....":  "h",
		"..":    "i",
		".---":  "j",
		"-.-":   "k",
		".-..":  "l",
		"--":    "m",
		"-.":    "n",
		"---":   "o",
		".--.":  "p",
		"--.-":  "q",
		".-.":   "r",
		"...":   "s",
		"-":     "t",
		"..-":   "u",
		"...-":  "v",
		".--":   "w",
		"-..-":  "x",
		"-.--":  "y",
		"--..":  "z",
		"-----": "0",
		".----": "1",
		"..---": "2",
		"...--": "3",
		"....-": "4",
		".....": "5",
		"-....": "6",
		"--...": "7",
		"---..": "8",
		"----.": "9",
	}

	var msg string
	for _, word := range strings.Split(morse, "/") {
		for _, code := range strings.Split(word, " ") {
			if c, ok := morseMap[code]; ok {
				msg += c
			}
		}
		msg += " "
	}

	return strings.TrimSpace(msg)
}

func ConvertStringToMorse(s string) string {
	morseMap := map[rune]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
		'0': "-----",
		'1': ".----",
		'2': "..---",
		'3': "...--",
		'4': "....-",
		'5': ".....",
		'6': "-....",
		'7': "--...",
		'8': "---..",
		'9': "----.",
		'.': ".-.-.-",
		',': "--..--",
		'?': "..--..",
		'!': "-.-.--",
		'-': "-....-",
		'"': ".-..-.",
		'\'': ".----.",
		'(': "-.--.",
		')': "-.--.-",
		'+': ".-.-.",
		'/': "-..-.",
		':': "---...",
		';': "-.-.-.",
		'=': "-...-",
		'_': "..--.-",
		'&': ".-...",
		'$': "...-..-",
		'@': ".--.-.",
		' ': "/",
	}

	var morse string
	for _, c := range s {
		morse += morseMap[unicode.ToLower(c)]
		morse += "/"
	}

	return strings.TrimSuffix(morse, "/")
}
