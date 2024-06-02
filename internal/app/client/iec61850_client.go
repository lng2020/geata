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
)

type IEC61850Client struct {
	IP   string
	Port string
	Ctx  context.Context
}

func NewIEC61850Client(ip string, port string) *IEC61850Client {
	return &IEC61850Client{
		IP:   ip,
		Port: port,
	}
}

func (c *IEC61850Client) Start(s chan string) {
	dir, err := os.Getwd()
	os := runtime.GOOS
	if err != nil {
		slog.Error("Failed to get the current working directory", logger.ErrAttr(err))
	}
	dir = fmt.Sprintf("%s/internal/app/client/%s/client", dir, os)
	cmd := exec.CommandContext(c.Ctx, dir, c.IP, c.Port)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		slog.Error("Failed to pipe stdout", logger.ErrAttr(err))
	}
	if err := cmd.Start(); err != nil {
		slog.Error("Failed to start the command", logger.ErrAttr(err))
	}
	defer cmd.Wait()
	defer cmd.Process.Kill()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		s <- scanner.Text()
	}
}

func (c *IEC61850Client) Close() {
	c.Ctx.Done()
}
