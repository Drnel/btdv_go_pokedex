package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    later gg hjkujkkkkvjj           ",
			expected: []string{"later", "gg", "hjkujkkkkvjj"},
		},
		{
			input:    "go gogo gogogogogo tt__$$123",
			expected: []string{"go", "gogo", "gogogogogo", "tt__$$123"},
		},
		{
			input:    "  23 123 12345",
			expected: []string{"23", "123", "12345"},
		},
		{
			input:    "45 678 999_999    ",
			expected: []string{"45", "678", "999_999"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test  failed %s != %s", word, expectedWord)
			}
		}
	}
}
