package morse

import (
	"errors"
	"strings"
)

func Encode(text string) (string, error) {

	splitText := splitNormalText(strings.ToLower(text))

	var result string

	for i, j := range splitText {

		val, exist := morseMap[j]

		if !exist {
			return "", errors.New("Invalid value: " + string(j))
		}

		result += val

		if i != len(splitText)-1 {
			result += " "
		}
	}

	return result, nil

}

func Decode(text string) (string, error) {

	splitText := splitMorseCodeText(text)
	morseMap := morseMapReversed()

	var result string

	for _, j := range splitText {

		val, exist := morseMap[j]

		if !exist {
			return "", errors.New("Invalid value: " + j)
		}

		result += string(val)

	}

	return result, nil

}

func splitNormalText(text string) []rune {
	return []rune(text)
}

func splitMorseCodeText(text string) []string {
	return strings.Split(text, " ")
}
