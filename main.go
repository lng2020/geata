package main

import (
	"geata/cmd"
	"log"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
