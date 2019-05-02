package string

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Hello", "olleH"},
		{"Hello 你好", "好你 olleH"},
		{"", ""},
	}

	for _, c := range tests {
		rs := Reverse(c.s)
		if rs != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, rs, c.want)
		}
	}
}
