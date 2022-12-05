package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	countLines, countBytes := false, false
	res := count(b, countLines, countBytes)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\n word3 word4")

	exp := 2
	countLines, countBytes := true, false
	res := count(b, countLines, countBytes)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

// TestCountBytes tests the count function set to count lines
func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word")

	exp := 4

	countLines, countBytes := false, true
	res := count(b, countLines, countBytes)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}
