package main

import "testing"

func TestReverse(t *testing.T){
	testcases := []struct{
		in, want string
	}{
		{"Hello, world", "dlrow, olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, testcase := range testcases {
		rev := Reverse(testcase.in)
		if rev != testcase.want{
			t.Errorf("Reverse: %q, want %q", rev, testcase.want)
		}
	}
}