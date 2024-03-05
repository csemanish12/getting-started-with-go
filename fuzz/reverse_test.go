package main

import (
	"testing"
	"unicode/utf8"
)


func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, testcase := range testcases {
		f.Add(testcase) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, original string) {
		reverse := Reverse(original)
		doubleRev := Reverse(reverse)

		if original!= doubleRev{
			t.Errorf("Reverse producedinvalid UTF-8 string %q", reverse)
		}

		if utf8.ValidString(original) && !utf8.ValidString(reverse){
			t.Errorf("Reverse produce invalid UTF-8 string %q", reverse)
		}
	})
}