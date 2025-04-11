package utils

import (
	"os"
	"strings"
)

// Exists checks if file exists
func Exists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// IsDir checks if path is a directory
func IsDirectory(file string) bool {
	if stat, err := os.Stat(file); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// Insert Beetween Matches
func InsertBeetweenMatches(original string, startSeparator string, endSeparator string, textToInsert string) (output string, err error) {
	isInReplaceMode := false
	var result []rune
	i := 0
	for i < len(original) {
		if isInReplaceMode {
			if strings.HasSuffix(original[0:i], endSeparator) {
				isInReplaceMode = false
				//result = append(result, rune('\n'))
				result = append(result, []rune(endSeparator)...)
				result = append(result, rune(original[i]))
			}
		} else {
			if strings.HasSuffix(original[0:i], startSeparator) {
				isInReplaceMode = true
				//result = append(result, rune('\n'))
				result = append(result, []rune(textToInsert)...)
			} else {
				result = append(result, rune(original[i]))
			}
		}

		i++
	}
	output = string(result)
	return output, err
}

// Return substring of text
func Substring(text string, from int, until int) string {
	if until == -1 {
		return text[from:]
	}

	return text[from:until]
}
