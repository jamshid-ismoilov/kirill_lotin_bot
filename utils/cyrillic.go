package utils

import (
	"regexp"
	"strings"
)

var (
	IsCyrillic = regexp.MustCompile(`\p{Cyrillic}`).MatchString
)

// CyrillicToLatin converts Cyrillic text to Latin text
func CyrillicToLatin(text string) string {
	// Map Cyrillic characters to Latin characters
	mapping := map[string]string{
		"а": "a", "А": "A",
		"б": "b", "Б": "B",
		"в": "v", "В": "V",
		"г": "g", "Г": "G",
		"д": "d", "Д": "D",
		"е": "e", "Е": "E",
		"ё": "yo", "Ё": "Yo",
		"ж": "j", "Ж": "J",
		"з": "z", "З": "Z",
		"и": "i", "И": "I",
		"й": "y", "Й": "Y",
		"к": "k", "К": "K",
		"л": "l", "Л": "L",
		"м": "m", "М": "M",
		"н": "n", "Н": "N",
		"о": "o", "О": "O",
		"п": "p", "П": "P",
		"р": "r", "Р": "R",
		"с": "s", "С": "S",
		"т": "t", "Т": "T",
		"у": "u", "У": "U",
		"ф": "f", "Ф": "F",
		"х": "x", "Х": "X",
		"ц": "ts", "Ц": "Ts",
		"ч": "ch", "Ч": "Ch",
		"ш": "sh", "Ш": "Sh",
		"щ": "sh", "Щ": "Sh",
		"ъ": "'", "Ъ": "'",
		"ы": "y", "Ы": "Y",
		"ь": "'", "Ь": "'",
		"э": "e", "Э": "E",
		"ю": "yu", "Ю": "Yu",
		"я": "ya", "Я": "Ya",
		"ў": "oʼ", "Ў": "Oʼ",
		"ҳ": "h", "Ҳ": "H",
		"қ": "q", "Қ": "Q",
		// Add more mappings for the remaining Cyrillic characters
	}
	// extra letters
	eLetters := map[rune]string{
		'е': "y", 'Е': "Y",
	}
	// Convert the text word by word
	words := strings.Fields(text)
	for i, word := range words {
		// Check if the word starts with Cyrillic "e"
		if strings.HasPrefix(strings.ToLower(word), "е") {
			runes := []rune(word)
			word = eLetters[runes[0]] + word
		}
		// Convert the word character by character
		convertedWord := ""
		for _, char := range word {
			cyrillicChar := string(char)
			latinChar, found := mapping[cyrillicChar]
			if found {
				convertedWord += latinChar
			} else {
				convertedWord += cyrillicChar
			}
		}
		words[i] = convertedWord
	}

	// Reconstruct the converted text with the modified words
	convertedText := strings.Join(words, " ")

	return convertedText
}
