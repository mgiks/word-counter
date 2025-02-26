package main

import "slices"

func getNonLexicalChars() []rune {
	var asciiPunctuationAndSymbols []rune

	// From control chars to '/'
	for char := 1; char < 48; char++ {
		asciiPunctuationAndSymbols = append(asciiPunctuationAndSymbols, int32(char))
	}

	// From ':' to '@'
	for char := 58; char < 65; char++ {
		asciiPunctuationAndSymbols = append(asciiPunctuationAndSymbols, int32(char))
	}

	// From '[' to '`'
	for char := 91; char < 97; char++ {
		asciiPunctuationAndSymbols = append(asciiPunctuationAndSymbols, int32(char))
	}

	// From '{' to '~'
	for char := 123; char < 127; char++ {
		asciiPunctuationAndSymbols = append(asciiPunctuationAndSymbols, int32(char))
	}

	var latinPunctuationAndSymbols []rune

	// From 'Non breaking space' to '¿'
	for char := 160; char < 192; char++ {
		latinPunctuationAndSymbols = append(latinPunctuationAndSymbols, int32(char))
	}

	var generalPunctuation []rune

	// From 'Whitespace char' to 'Bidirectional text char'
	for char := 8192; char < 8297; char++ {
		generalPunctuation = append(generalPunctuation, int32(char))
	}

	var cjkSymbolsAndPunctuation []rune

	for char := 3000; char < 12351; char++ {
		cjkSymbolsAndPunctuation = append(cjkSymbolsAndPunctuation, int32(char))
	}

	var nonLexicalChars []rune

	nonLexicalChars = append(nonLexicalChars, asciiPunctuationAndSymbols...)
	nonLexicalChars = append(nonLexicalChars, latinPunctuationAndSymbols...)
	nonLexicalChars = append(nonLexicalChars, generalPunctuation...)
	nonLexicalChars = append(nonLexicalChars, cjkSymbolsAndPunctuation...)

	return nonLexicalChars
}

func isLexicalWrapper() func(rune) bool {
	nonLexicalChars := getNonLexicalChars()

	return func(char rune) bool {
		if slices.Contains(nonLexicalChars, char) {
			return false
		}

		return true
	}
}

func countWords(text string) int {
	wordCount := 0

	// Needed for support outside of ASCII
	runifiedText := []rune(text)

	isOnWord := false
	textLength := len(runifiedText)
	charIndex := 0

	isLexical := isLexicalWrapper()

	for charIndex < textLength {
		char := runifiedText[charIndex]

		if isLexical(char) {
			isOnWord = true
		}

		if !isOnWord {
			charIndex++
			continue
		}

		if !isLexical(char) || charIndex == textLength-1 {
			wordCount += 1
			isOnWord = false
		}

		charIndex++
	}

	return wordCount
}
