// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"geata/cmd/start"
	"geata/internal/utils/logger"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "geata"}
	rootCmd.AddCommand(start.NewStartCmd())

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal.Printf("Error: %v", err)
	}
}
