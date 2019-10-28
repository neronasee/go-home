package morse

import (
	"bytes"
	"fmt"
	"strings"
)

type Code string

var morseMap = map[string]string{
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
	" ":     " ",
}

func Translate(code Code) (string, error) {
	buffer := bytes.NewBufferString("")

	for _, word := range strings.Split(string(code), " / ") {
		for _, char := range strings.Split(word, " ") {
			v, ok := morseMap[char]

			if !ok {
				return "", fmt.Errorf("Uknown char: %s", char)
			}

			buffer.WriteString(v)
		}

		buffer.WriteString(" ")
	}
	return buffer.String(), nil
}
