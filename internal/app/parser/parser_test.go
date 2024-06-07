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

	ied := scl.IED
	if ied.Name != expectedIEDName {
		t.Errorf("Unexpected IED name. Got: %s, Expected: %s", ied.Name, expectedIEDName)
	}

	if len(ied.AccessPoint) == 0 {
		t.Fatal("No access point found in the IED")
	}
	if ied.AccessPoint[0].Name != expectedAccessPointName {
		t.Errorf("Unexpected access point name. Got: %s, Expected: %s", ied.AccessPoint[0].Name, expectedAccessPointName)
	}
}
