package main

import (
	"testing"
)

// TestRomantoInt calls hello.romanInt with a MMVI, checking
// for a valid return value.
func TestRomantoInt(t *testing.T) {
	expected := 2006
	ans := romanToInt("MMVI")
	if ans != expected {
		t.Fatalf("expected %v but got %v", expected, ans)
	}
}

// TestRomantoInt calls hello.romanInt with a MCMXLIV, checking
// for a valid return value.
func TestRomantoIntSubtract(t *testing.T) {
	expected := 1944
	ans := romanToInt("MCMXLIV")
	if ans != expected {
		t.Fatalf("expected %v but got %v", expected, ans)
	}
}

// TestRomantoInt calls hello.romanInt with a MCMXLIV, checking
// for a invalid return value.
func TestRomantoIntSubtractFailCase(t *testing.T) {
	expected := 0
	ans := romanToInt("")
	if ans != expected {
		t.Fatalf("expected %v but got %v", expected, ans)
	}
}
