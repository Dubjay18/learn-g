package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"14:07:33", true},
		{"14:07", false},
		{"14:07:60", false},
		{"14:60:33", false},
		{"24:07:33", false},
		{"14:07:33", true},
		{"-14:07:33", false},
		{"14:-07:33", false},
		{"14:07:-33", false},
		{"14:07:33", true},
		{"14:07:33", true},
		{"14:07:33", true},
		{"14:07:33", true},
		{"14:07:33", true},
	}
	for _, test := range table {
		_, err := ParseTime(test.time)
		if test.ok && err != nil {
			t.Errorf("%v: unexpected error: %v", test.time, err)
		}
	}
}
