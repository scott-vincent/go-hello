package morestrings

import "testing"

func TestReverseString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseString(c.in)
		if got != c.want {
			t.Errorf("ReverseString(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
