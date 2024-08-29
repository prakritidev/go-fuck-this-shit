package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName (t *testing.T)  {
	name :="Glasys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Gladys")

	if !want.MatchString(msg) || err != nil {
		msg, err := Hello("")

		if msg !="" || err == nil {
			t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
		}
	}
}