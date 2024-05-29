package handler

import (
	"geata/internal/app/client"
)

type IEC61850Handler struct {
	client *client.IEC61850Client
}

type IEC61850HandlerConfig struct {
	IP   string
	Port string
}

func (c IEC61850HandlerConfig) Type() HandlerType {
	return IEC61850HandlerType
}

func NewIEC61850HandlerConfig(ip, port string) HandlerConfig {
	return IEC61850HandlerConfig{
		IP:   ip,
		Port: port,
	}
}

func NewIEC61850Handler(hc HandlerConfig) Handler {
	config := hc.(IEC61850HandlerConfig)
	client := client.NewIEC61850Client(config.IP, config.Port)
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
