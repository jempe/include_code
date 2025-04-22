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
	var result strings.Builder
	isInReplaceMode := false

	for i := 0; i < len(original); {
		if isInReplaceMode {
			if strings.HasPrefix(original[i:], endSeparator) {
				isInReplaceMode = false
				result.WriteString(endSeparator)
				i += len(endSeparator)
				continue
			}
		} else {
			if strings.HasPrefix(original[i:], startSeparator) {
				isInReplaceMode = true
				result.WriteString(startSeparator)
				result.WriteString(textToInsert)
				i += len(startSeparator)
				continue
			}

			result.WriteByte(original[i])
		}
		i++
	}

	output = result.String()
	return output, nil
}

// Return substring of text
func Substring(text string, from int, until int) string {
	if until == -1 {
		return text[from:]
	}

	return text[from:until]
}
