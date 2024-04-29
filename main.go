package main

import (
	"geata/cmd"
	"log/slog"
)

func main() {
	err := cmd.Run()
	if err != nil {
		slog.Error("Failed to run the command: ", err)
	}
}
