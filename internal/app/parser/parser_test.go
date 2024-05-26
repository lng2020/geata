package parser_test

import (
	"geata/internal/app/parser"
	"testing"
)

func TestParseIEC61850Model(t *testing.T) {
	filename := "test.icd"

	expectedHeader := struct {
		ID            string
		NameStructure string
	}{ID: "", NameStructure: "IEDName"}

	expectedIEDName := "TEMPLATE"
	expectedAccessPointName := "accessPoint1"

	scl, err := parser.ParseIEC61850Model(filename)
	if err != nil {
		t.Fatalf("Error parsing IEC 61850 model: %v", err)
	}

	if scl.Header.ID != expectedHeader.ID {
		t.Errorf("Unexpected header ID. Got: %s, Expected: %s", scl.Header.ID, expectedHeader.ID)
	}
	if scl.Header.NameStructure != expectedHeader.NameStructure {
		t.Errorf("Unexpected header nameStructure. Got: %s, Expected: %s", scl.Header.NameStructure, expectedHeader.NameStructure)
	}

	if len(scl.IED) == 0 {
		t.Fatal("No IED found in the model")
	}
	if scl.IED[0].Name != expectedIEDName {
		t.Errorf("Unexpected IED name. Got: %s, Expected: %s", scl.IED[0].Name, expectedIEDName)
	}

	if len(scl.IED[0].AccessPoint) == 0 {
		t.Fatal("No access point found in the IED")
	}
	if scl.IED[0].AccessPoint[0].Name != expectedAccessPointName {
		t.Errorf("Unexpected access point name. Got: %s, Expected: %s", scl.IED[0].AccessPoint[0].Name, expectedAccessPointName)
	}
}
