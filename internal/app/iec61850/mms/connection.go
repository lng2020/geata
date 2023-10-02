package mms

import (
	"geata/internal/app/iec61850/iso"
	"sync"
)

type MmsConnectionState int

const (
	MMS_CONNECTION_STATE_CLOSED MmsConnectionState = iota
	MMS_CONNECTION_STATE_CONNECTING
	MMS_CONNECTION_STATE_CONNECTED
	MMS_CONNECTION_STATE_CLOSING
)

type MmsOutstandingCall struct {
	isUsed   bool
	invokeId int
	// TODO
}

type MmsConnectionParameters struct {
	maxServOutstandingCalling int
	maxServOutstandingCalled  int
	dataStructureNestingLevel int
	maxPduSzie                int
	servicesSupported         [11]byte
}

type MmsConnection struct {
	mu           sync.Mutex
	nextInvokeId int

	outstandingCalls *MmsOutstandingCall

	requestTimeout int
	connectTimeout int

	isoClient *iso.IsoClientConnection

	connectionState MmsConnectionState

	parameters    *MmsConnectionParameters
	isoParameters *iso.IsoConnectionParameters
}
