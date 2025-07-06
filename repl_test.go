package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input      string
		expected   []string
	}{
		{
			input:     "  hello  world  ",
			expected:  []string{"hello", "world"},
		},
		{
			input:     "HEY world",
			expected:  []string{"hey", "world"},
		},
		{
			input:     "hi     world",
			expected:  []string{"hi", "world"},
		},
		{
			input:     "",
			expected:  []string{},
		},
		{
			input:     "world",
			expected:  []string{"world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("the actual output length %v does not match the length of the expected output %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("the word %v does not match the expected word %v", actual, c.expected)
			}
		}
	}
}
