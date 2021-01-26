package pocket

import "testing"

func TestInStringArray(t *testing.T) {
	testCases := []struct {
		str      string
		strArr   []string
		expected bool
	}{
		{"a", []string{}, false},
		{"", []string{"a", "b"}, false},
		{"abc", []string{"a", "abc"}, true},
		{"aa", []string{"aa"}, true},
	}

	for _, tc := range testCases {
		v := InStringArray(tc.str, tc.strArr)
		if v != tc.expected {
			t.Fatalf("unexpected value: %t != %t", tc.expected, v)
		}
	}
}
