package eu

import "testing"

func TestPercent(t *testing.T) {
	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      23,
			v:        0,
			expected: "% 23",
		},
		{
			num:      23.45,
			v:        2,
			expected: "% 23,45",
		},
		{
			num:      1023.45,
			v:        2,
			expected: "% 1.023,45",
		},
		{
			num:      -1023.45,
			v:        2,
			expected: "% \u2212" + "1.023,45",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtPercent(tt.num, tt.v))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}
