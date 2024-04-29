package client

import (
	"bufio"
	"context"
	"log"
	"os/exec"
)

type IEC61850Client struct {
	IP   string
	Port string
	ctx  context.Context
}

func NewIEC61850Client(ip string, port string) *IEC61850Client {
	return &IEC61850Client{
		IP:   ip,
		Port: port,
	}
}

func (c *IEC61850Client) Start(s chan string) {
	cmd := exec.CommandContext(c.ctx, "iec61850_client", c.IP, c.Port)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	defer cmd.Wait()
	defer cmd.Process.Kill()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		s <- parse(scanner.Text())
	}
}

func (c *IEC61850Client) Close() {
	c.ctx.Done()
}

func parse(_ string) string {
	// TODO: Implement the parsing logic here
	return ""
}
