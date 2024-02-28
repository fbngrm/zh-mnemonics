package tone

import (
	"fmt"
	"testing"
)

func TestCheckTone(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"bā", firstTone},
		{"báo", secondTone},
		{"bǎn", thirdTone},
		{"bà", fourthTone},
		{"abc", noTone},
		{"yí", secondTone},
		{"yùe", fourthTone},
		{"yue1", firstTone},
		{"yue2", secondTone},
		{"yue3", thirdTone},
		{"yue4", fourthTone},
		{"yue5", noTone},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := getTone(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, result)
			}
		})
	}
}

func TestReplaceToneCharacters(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"āēīōū", "aeiou"},
		{"áéíóú", "aeiou"},
		{"ǎěǐǒǔ", "aeiou"},
		{"àèìòù", "aeiou"},
		{"abc", "abc"},
		{"íóǒù", "ioou"},
		{"ùǚē~", "ue"},
		{"漢字", ""}, // Test with non-ASCII characters, they should remain unchanged
		{"dao1", "dao"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := ReplaceToneCharacters(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}
}
