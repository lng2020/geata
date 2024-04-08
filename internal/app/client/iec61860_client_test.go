package client

import (
	"testing"
)

var raw = `{"data": "test"}`
var expected = `{"data": "test"}`

func TestParse(t *testing.T) {
	if expected != parse(raw) {
		t.Errorf("Failed to parse %s. Expected: %s, Actual: %s", raw, expected, parse(raw))
	}
}
