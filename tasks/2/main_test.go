package main

import "testing"

type test struct {
	input    string
	expected string
	err      error
}

func TestChangeString(t *testing.T) {
	var tests = []test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidStringDigit,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidStringDigit,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidTwoDigit,
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
		{
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
		{
			input:    "Hel2o!",
			expected: "Hello!",
		},
	}

	for _, tst := range tests {
		result, err := ChangeString(tst.input)
		if tst.err != err || tst.expected != result {
			t.Errorf("Unpack(%q) = %q, %v; want: %q, %v", tst.input, result, err, tst.expected, tst.err)
		}
	}
}

func TestUnpackWithSlash(t *testing.T) {
	var tests = []test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
		{
			input:    `qwe\\5a`,
			expected: `qwe\\\\\a`,
		},
		{
			input:    `\\`,
			expected: `\`,
		},
	}

	for _, tst := range tests {
		result, err := ChangeString(tst.input)
		if tst.err != err || tst.expected != result {
			t.Errorf("Unpack(%q) = %q, %v; want: %q", tst.input, result, err, tst.expected)
		}
	}
}
