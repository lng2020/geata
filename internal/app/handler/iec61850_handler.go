package handler

import (
	"geata/internal/app/client"
)

type IEC61850Handler struct {
	client *client.IEC61850Client
}

func NewIEC61850Handler(client *client.IEC61850Client) *IEC61850Handler {
	return &IEC61850Handler{client: client}
}

func (h *IEC61850Handler) Handle(s chan string) {
	h.client.Start(s)
}

func (h *IEC61850Handler) Close() {
	if h.client != nil {
		h.client.Close()
	}
}

func (h *IEC61850Handler) IsOnline() bool {
	return h.client != nil
}

func (h *IEC61850Handler) Type() HandlerType {
	return IEC61850HandlerType
}
