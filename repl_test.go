package main

import "testing"

func TestCleanInput(t *testing.T) {
	testcases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  a b c  d  ",
			expected: []string{"a", "b", "c", "d"},
		},
		{
			input:    "  Oi vei",
			expected: []string{"oi", "vei"},
		},
	}
	for _, tc := range testcases {
		actual := cleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("actual len %d different than expected %d", len(actual), len(tc.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := tc.expected[i]
			if word != expectedWord {
				t.Errorf("failed: %s != %s", word, expectedWord)
			}
		}
	}
}
