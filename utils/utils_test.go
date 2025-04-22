package utils

import "testing"

func TestExists(t *testing.T) {
	// Test cases for Exists function
}

func TestIsDirectory(t *testing.T) {
	// Test cases for IsDirectory function
}

func TestInsertBeetweenMatches(t *testing.T) {
	cases := []struct {
		original       string
		startSeparator string
		endSeparator   string
		textToInsert   string
		expected       string
	}{
		// Case 1: Basic insertion
		{"Hello [World]", "[", "]", "Beautiful ", "Hello [Beautiful ]"},
		// Case 2: No match
		{"Hello World", "[", "]", "Beautiful ", "Hello World"},
		// Case 3: Multiple matches
		{"[Hello] [World]", "[", "]", "Beautiful ", "[Beautiful ] [Beautiful ]"},
		// Case 4: Multibyte characters
		{"こんにちは[世界]", "[", "]", "美しい", "こんにちは[美しい]"},
	}

	for _, c := range cases {
		output, err := InsertBeetweenMatches(c.original, c.startSeparator, c.endSeparator, c.textToInsert)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if output != c.expected {
			t.Errorf("expected %q but got %q", c.expected, output)
		}
	}
}

func TestSubstring(t *testing.T) {
	// Test cases for Substring function
}
