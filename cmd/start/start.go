// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package start

import (
	"context"
	"geata/internal/app"
	"geata/internal/app/logger"
	"log/slog"

	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {
	var configFile string
	var serverPort int64
	var loggerLogLevel string

	var cmd = &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			myApp := app.NewApp(ctx)

			if err := myApp.NewConfig(configFile); err != nil {
				return err
			}

			if serverPort != 0 {
				myApp.Config.Server.Port = serverPort
			}
			if loggerLogLevel != "" {
				myApp.Config.Logger.LogLevel = loggerLogLevel
			}

			logger.InitLogger(myApp.Config.Logger.LogLevel)

			if err := myApp.Init(); err != nil {
				slog.Error("Failed to initialize the application: ", err)
				return nil
			}

			if err := myApp.Start(); err != nil {
				slog.Error("Failed to start the application: ", err)
				return nil
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "./internal/app/config.yaml", "Path to the configuration file")
	cmd.Flags().Int64Var(&serverPort, "server.port", 0, "Port on which the server should listen")
	cmd.Flags().StringVar(&loggerLogLevel, "logger.log_level", "", "Logging level (e.g., info, error, debug)")

	return cmd
}
