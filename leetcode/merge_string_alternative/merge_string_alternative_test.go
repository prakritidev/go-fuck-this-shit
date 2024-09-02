package main

import (
	"testing"
)

func TestMergeStringAlternative(t *testing.T) {
	s1 := "abc"
	s2 := "def"

	result := Merge_string_alternative(s1, s2)
	expected := "adbecf"

	if result != expected {
		t.Errorf("MergeStringAlternative(%q, %q) = %q; want %q", s1, s2, result, expected)
	}
}
