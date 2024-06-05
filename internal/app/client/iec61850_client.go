package client

import (
	"bufio"
	"context"
	"fmt"
	"geata/internal/app/logger"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type IEC61850Client struct {
	IP   string
	Port string
}

func NewIEC61850Client(ip string, port string) *IEC61850Client {
	return &IEC61850Client{
		IP:   ip,
		Port: port,
	}
}

func (c *IEC61850Client) Start(ctx context.Context, s chan string, interval time.Duration) {
	dir, err := os.Getwd()
	os := runtime.GOOS
	if err != nil {
		slog.Error("Failed to get the current working directory", logger.ErrAttr(err))
		return
	}
	dir = fmt.Sprintf("%s/internal/app/client/%s/client", dir, os)

	for {
		cmd := exec.CommandContext(ctx, dir, c.IP, c.Port)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			slog.Error("Failed to pipe stdout", logger.ErrAttr(err))
			return
		}
		if err := cmd.Start(); err != nil {
			slog.Error("Failed to start the command", logger.ErrAttr(err))
			return
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			s <- scanner.Text()
		}

		if err := cmd.Wait(); err != nil {
			slog.Error("Command execution failed", logger.ErrAttr(err))
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
		}
	}
}
