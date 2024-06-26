package parser

import (
	"encoding/xml"
	"os"
)

func ParseIEC61850Model(filename string) (*SCL, error) {
	var scl SCL
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&scl); err != nil {
		return nil, err
	}

	return &scl, nil
}
