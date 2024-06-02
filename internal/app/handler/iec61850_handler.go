package handler

import (
	"context"
	"geata/internal/app/client"
	"geata/internal/app/logger"
	"log/slog"
	"strings"
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
	return &IEC61850HandlerConfig{
		IP:   ip,
		Port: port,
	}
}

func (hc *IEC61850HandlerConfig) NewHandler() Handler {
	client := client.NewIEC61850Client(hc.IP, hc.Port)
	return &IEC61850Handler{client: client}
}

func (h *IEC61850Handler) Handle(ctx context.Context, s chan Data) {
	stringChan := make(chan string)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer close(stringChan)
	h.client.Ctx = ctx
	go h.client.Start(stringChan)
	for output := range stringChan {
		data := Parse(output)
		slog.Info("Received data", logger.StringAttr("IEC61850Ref", data.IEC61850Ref), logger.StringAttr("Value", data.Value))
		if data.IEC61850Ref != "" {
			s <- data
		}
	}
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

func Parse(output string) Data {
	// Example: ZTPDFMONT/SBMO36.SupDevRun[ST]:{false,0000000000000,19700101000000.000Z}
	// ZTPDFMONT/SHUM21.DeHumOp[SG]:{false}
	// ZTPDFMONT/SNOI33.Noise[MX]:{{56.000000},0000000000000,19700101000000.000Z}
	// we split ':' to determine reference and value
	// refercence examle: ZTPDFMONT/SBMO36$SupDevRun$stVal
	// value example: false
	// reference example: ZTPDFMONT/SHUM21$DeHumOp$setVal
	// value example: false
	// reference example: ZTPDFMONT/SNOI33$Noise$mag
	// value example: 56.000000
	// we split '[]' to determine FC and accordingly parse the value
	refVal := strings.Split(output, ":")
	if len(refVal) != 2 {
		slog.Error("Failed to parse the output", logger.StringAttr("output", output))
		return Data{}
	}
	ref := refVal[0]
	val := refVal[1]
	fc := strings.Split(ref, "[")[1]
	fc = strings.Split(fc, "]")[0]
	ref = strings.Split(ref, "[")[0]
	ref = strings.ReplaceAll(ref, ".", "$")
	switch fc {
	case "ST":
		// {false,0000000000000,19700101000000.000Z}
		val = strings.Split(val, ",")[0]
		val = strings.Split(val, "{")[1]
		ref = ref + "$stVal"
		return Data{
			IEC61850Ref: ref,
			Value:       val,
		}
	case "SG":
		// {false}
		val = strings.Split(val, "{")[1]
		val = strings.Split(val, "}")[0]
		ref = ref + "$setVal"
		return Data{
			IEC61850Ref: ref,
			Value:       val,
		}
	case "MX":
		// {{56.000000},0000000000000,19700101000000.000Z}
		val = strings.Split(val, ",")[0]
		val = strings.Split(val, "{")[1]
		val = strings.Split(val, "}")[0]
		ref = ref + "$mag"
		return Data{
			IEC61850Ref: ref,
			Value:       val,
		}
	}
	return Data{}
}
