package main

import "testing"

func TestPrintMessage(t *testing.T) {
	if printmessage() != "Hello Go World From Method !" {
		t.Fatal("Test fail")
	}
}
