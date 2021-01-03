package server

import "testing"

func TestPrintMessage(t *testing.T) {
	expected := "Hello World !!!"
	got := PrintMessage()
	if got != expected {
		t.Errorf("Message received [%s] not same as expected [%s]", got, expected)
	}
}
