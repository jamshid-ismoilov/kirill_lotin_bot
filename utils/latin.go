package utils

import (
	"regexp"
	"strings"
)

var (
	IsLatin = regexp.MustCompile(`\p{Latin}`).MatchString
)

// CyrillicToLatin converts Cyrillic text to Latin text
func LatinToCyrillic(text string) string {
	// Map Cyrillic characters to Latin characters
	LatinMapping := map[string]string{
		"a": "а", "A": "А",
		"b": "б", "B": "Б",
		"v": "в", "V": "В",
		"g": "г", "G": "Г",
		"d": "д", "D": "Д",
		"j": "ж", "J": "Ж",
		"z": "з", "Z": "З",
		"i": "и", "I": "И",
		"y": "й", "Y": "Й",
		"k": "к", "K": "К",
		"l": "л", "L": "Л",
		"m": "м", "M": "М",
		"n": "н", "N": "Н",
		"o": "о", "O": "О",
		"p": "п", "P": "П",
		"r": "р", "R": "Р",
		"s": "с", "S": "С",
		"t": "т", "T": "Т",
		"u": "у", "U": "У",
		"f": "ф", "F": "Ф",
		"x": "х", "X": "Х",
		"ts": "ц", "Ts": "Ц",
		"sch": "щ", "Sch": "Щ",
		"'": "ъ",
		"e": "е", "E": "Е",
		"h": "ҳ", "H": "Ҳ",
		"q": "қ", "Q": "Қ",
		// Add more mappings for the remaining Latin to Cyrillic characters
	}
	// Map two Latin characters
	LatinMapping2 := map[string]string{
		"yo": "ё", "Yo": "Ё",
		"ch": "ч", "Ch": "Ч",
		"sh": "ш", "Sh": "Ш",
		"yu": "ю", "Yu": "Ю",
		"ya": "я", "Ya": "Я",
		"oʼ": "ў", "Oʼ": "Ў",
		"oʻ": "ў", "Oʻ": "Ў",
		"o'": "ў", "O'": "Ў",
		"o`": "ў", "O`": "Ў",
		// Add more mappings for the remaining Latin to Cyrillic characters
	}
	// extra letters
	eLetters := map[rune]string{
		'e': "э", 'E': "Э",
		'y': "е", 'Y': "Е",
	}
	// Convert the text word by word
	words := strings.Fields(text)
	for i, word := range words {
		// Check if the word starts with Cyrillic "e"
		if strings.HasPrefix(strings.ToLower(word), "e") {
			runes := []rune(word)
			word = eLetters[runes[0]] + word[1:]
		}

		// Check if the word starts with Cyrillic "ye"
		if strings.HasPrefix(strings.ToLower(word), "ye") {
			runes := []rune(word)
			word = eLetters[runes[0]] + word[2:]
		}
		// First replace two character Latin characters
		for key, value := range LatinMapping2 {
			word = strings.ReplaceAll(word, key, value)
		}

		// Convert the word character by character
		convertedWord := ""
		for _, char := range word {
			latinChar := string(char)
			cyrillicChar, found := LatinMapping[latinChar]
			if found {
				convertedWord += cyrillicChar
			} else {
				convertedWord += latinChar
			}
		}
		words[i] = convertedWord
	}

	// Reconstruct the converted text with the modified words
	convertedText := strings.Join(words, " ")

	return convertedText
}
