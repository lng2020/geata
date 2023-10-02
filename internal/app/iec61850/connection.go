package iec61850

import "geata/internal/app/iec61850/mms"

type IedConnectionState int

const (
	IED_STATE_CLOSED IedConnectionState = iota
	IED_STATE_CONNECTING
	IED_STATE_CONNECTED
	IED_STATE_CLOSING
)

type IedConnection struct {
	connection    *mms.MmsConnection
	state         IedConnectionState
	enabledReport []*ClientReport
}

func CreateIedConnection() *IedConnection {
	return &IedConnection{
		state: IED_STATE_CLOSED,
	}
}
