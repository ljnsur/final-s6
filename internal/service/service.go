package service

import "s6-final/pkg/morse"

func MorseOrTextRecognition(data []byte) (string, error) {
	if string(data[0]) == "-" || string(data[0]) == "." {
		textFromMorse := morse.ToText(string(data))
		return textFromMorse, nil
	}

	morseFromText := morse.ToMorse(string(data))
	return morseFromText, nil
}
