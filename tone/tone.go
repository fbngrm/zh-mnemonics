package tone

import "strings"

const (
	firstTone = iota
	secondTone
	thirdTone
	fourthTone
	noTone
)

var firstTones = []string{
	"1",
	"ā",
	"ē",
	"ī",
	"ō",
	"ū",
	"ǖ",
}
var secondTones = []string{
	"2",
	"á",
	"é",
	"í",
	"ó",
	"ú",
	"ǘ",
}
var thirdTones = []string{
	"3",
	"ǎ",
	"ě",
	"ǐ",
	"ǒ",
	"ǔ",
	"ǚ",
}
var fourthTones = []string{
	"4",
	"à",
	"è",
	"ì",
	"ò",
	"ù",
	"ǜ",
}

func getTone(input string) int {
	if containsAny(input, firstTones) {
		return firstTone
	} else if containsAny(input, secondTones) {
		return secondTone
	} else if containsAny(input, thirdTones) {
		return thirdTone
	} else if containsAny(input, fourthTones) {
		return fourthTone
	}
	return noTone
}

func containsAny(input string, toneSlice []string) bool {
	for _, char := range toneSlice {
		if strings.Contains(input, char) {
			return true
		}
	}
	return false
}

func ReplaceToneCharacters(input string) string {
	// map tone characters to their ASCII equivalents
	toneMapping := map[string]string{
		"ā": "a",
		"ē": "e",
		"ī": "i",
		"ō": "o",
		"ū": "u",
		"á": "a",
		"é": "e",
		"í": "i",
		"ó": "o",
		"ú": "u",
		"ǎ": "a",
		"ě": "e",
		"ǐ": "i",
		"ǒ": "o",
		"ǔ": "u",
		"à": "a",
		"è": "e",
		"ì": "i",
		"ò": "o",
		"ù": "u",
		"1": "",
		"2": "",
		"3": "",
		"4": "",
		"5": "",
	}

	// replace tone characters with ASCII equivalents
	for toneChar, asciiEquivalent := range toneMapping {
		input = strings.ReplaceAll(input, toneChar, asciiEquivalent)
	}

	// remove any non-ASCII characters
	var result strings.Builder
	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			result.WriteRune(char)
		}
	}

	return result.String()
}
