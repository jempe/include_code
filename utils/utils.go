package utils

import (
	"errors"
	"os"
	"strings"
)

// Exists checks if file exists
//
func Exists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// IsDir checks if path is a directory
//
func IsDirectory(file string) bool {
	if stat, err := os.Stat(file); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// Insert Beetween Matches
//
func InsertBeetweenMatches(original string, startSeparator string, endSeparator string, textToInsert string) (output string, err error) {
	if strings.Contains(original, startSeparator) && strings.Contains(original, endSeparator) {
		startSeparatorPosition := strings.Index(original, startSeparator)
		beforeStartSeparator := Substring(original, 0, startSeparatorPosition)
		afterStartSeparator := Substring(original, startSeparatorPosition, -1)

		if strings.Contains(afterStartSeparator, endSeparator) {
			endSeparatorPosition := startSeparatorPosition + strings.Index(afterStartSeparator, endSeparator) + len(endSeparator)

			output = beforeStartSeparator + startSeparator + textToInsert + endSeparator + Substring(original, endSeparatorPosition, -1)
		} else {
			err = errors.New("could not insert the content beetween " + startSeparator + " and " + endSeparator)
		}
	} else {
		err = errors.New("couldn't insert the content beetween " + startSeparator + " and " + endSeparator)
		output = original
	}

	return output, err
}

//Return substring of text
//
func Substring(text string, from int, until int) string {
	if until == -1 {
		return text[from:]
	}

	return text[from:until]
}
