package main

import (
	"errors"
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gurusudhan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("%s") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}

}

// Unit testing function
func Hello(name string) (string, error) {
	var err error
	if len(name) == 2 {
		err = errors.New("short name")
	}
	return name, err
}
