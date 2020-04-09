package main

import "testing"

func TestPrintMessage(t *testing.T) {
	expected := "Hello World !!"
	got := printMessage()
	if got != expected {
		t.Errorf("Message received [%s] not same as expected [%s]", got, expected)
	}
}
