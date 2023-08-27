package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, c := range cases {
		cleanedInput := cleanInput(c.input)
		if len(cleanedInput) != len(c.expected) {
			t.Errorf(
				"The lengths are not equal: %v vs %v",
				len(cleanedInput),
				len(c.expected),
			)
			continue
		}
		for i := range cleanedInput {
			cleanedInputWord := cleanedInput[i]
			expectedWord := c.expected[i]
			if cleanedInputWord != expectedWord {
				t.Errorf(
					"%v does not equal %v",
					cleanedInputWord,
					expectedWord,
				)
			}
			continue
		}
	}
}
