package client

import (
	"testing"
)

var raw = `{"data": "test"}`
var expected = `{"data": "test"}`

func TestParse(t *testing.T) {
	s := make(chan string)
	go parse(raw, s)
	actual := <-s
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
