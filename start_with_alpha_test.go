package validatorx

import (
	`testing`
)

func TestStartWithAlpha(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "", expected: false},
		{input: "a", expected: true},
		{input: "a1", expected: true},
		{input: "&", expected: false},
		{input: "1abcdefghijklmnopqrstuvwxyz", expected: false},
		{input: " ", expected: false},
	}

	for _, tc := range tests {
		got := StartWithAlpha(tc.input)
		if got != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}
