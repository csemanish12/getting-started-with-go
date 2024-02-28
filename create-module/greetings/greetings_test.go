package greetings

import (
	"testing"
	"regexp"
)
// est functions take a pointer to the testing package's testing.T type as a parameter. You use this parameter's methods for reporting and logging from your test.


func TestHaloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	message, err := Hello(name)
	
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	
	if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}