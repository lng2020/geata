package handler

import (
	"geata/internal/app/client"
)

type iec61850_handler struct {
	client *client.IEC61850Client
}

func NewIEC61850Handler(client *client.IEC61850Client) *iec61850_handler {
	return &iec61850_handler{client: client}
}

func (h *iec61850_handler) Handle(s chan string) {
	h.client.Start(s)
}
