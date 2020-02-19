package main

import (
	"testing"
)

func Test_tokenize_simple(t *testing.T) {
	tokens := tokenize("123 abc[123]{abc}")
	if len(tokens) != 8 {
		t.Errorf("expect 8 but got %d", len(tokens))
	}
}

func Test_read_str_simple(t *testing.T) {
	reader := read_str("123 abc[123]{abc}")
	if len(reader.tokens) != 8 {
		t.Errorf("expect 8 tokens but got %d", len(reader.tokens))
	}	
}