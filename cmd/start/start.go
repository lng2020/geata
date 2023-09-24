// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package start

import (
	"geata/internal/app"

	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {
	var configFile string
	var serverPort int
	var loggerLogLevel string
	var modbusServerAddress string
	var modbusSlaveID int

	var cmd = &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		RunE: func(cmd *cobra.Command, args []string) error {
			myApp := app.NewApp()

			if err := myApp.NewConfig(configFile); err != nil {
				return err
			}

			if serverPort != 0 {
				myApp.Config.Server.Port = serverPort
			}
			if loggerLogLevel != "" {
				myApp.Config.Logger.LogLevel = loggerLogLevel
			}
			if modbusServerAddress != "" {
				myApp.Config.Plugins.Modbus.ServerAddress = modbusServerAddress
			}
			if modbusSlaveID != 0 {
				myApp.Config.Plugins.Modbus.SlaveID = modbusSlaveID
			}

			if err := myApp.NewHandlers(); err != nil {
				return err
			}

			if err := myApp.Start(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "./internal/app/config.yaml", "Path to the configuration file")
	cmd.Flags().IntVar(&serverPort, "server.port", 0, "Port on which the server should listen")
	cmd.Flags().StringVar(&loggerLogLevel, "logger.log_level", "", "Logging level (e.g., info, error, debug)")
	cmd.Flags().StringVar(&modbusServerAddress, "plugins.modbus.server_address", "", "Modbus server address")
	cmd.Flags().IntVar(&modbusSlaveID, "plugins.modbus.slave_id", 0, "Modbus slave ID")

	return cmd
}
