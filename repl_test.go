package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "this is BASED",
			expected: []string{"this", "is", "based"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Error("the length of arrays do not match")
		}
		for i := range actual {
			got := actual[i]
			want := c.expected[i]
			if got != want {
				t.Errorf("got %q and want %q", got, want)
			}
		}
	}

}
