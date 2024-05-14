// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package cmd

import (
	"geata/cmd/mock"
	"geata/cmd/start"

	"github.com/spf13/cobra"
)

func Run() error {
	var rootCmd = &cobra.Command{
		Use:   "geata",
		Short: "GEATA is a gateway application for SCADA systems",
		Long:  `GEATA is a gateway application for SCADA systems`,
	}

	rootCmd.AddCommand(start.NewStartCmd())
	rootCmd.AddCommand(mock.NewMockCmd())

	return rootCmd.Execute()
}
