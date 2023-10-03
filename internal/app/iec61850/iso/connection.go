package iso

import (
	"net"
)

type IsoConnectionState int

const (
	ISO_STATE_IDLE IsoConnectionState = iota
	ISO_STATE_TCP_CONNECTING
	ISO_STATE_WAIT_FOR_COTP_CONNECTION_RESP
	ISO_STATE_WAIT_FOR_ACSE_RESP
	ISO_STATE_WAIT_FOR_DATA_MSG
	ISO_STATE_CLOSING_CONNECTION
	ISO_STATE_CLOSE_ON_ERROR
	ISO_STATE_ERROR
)

type IsoConnectionParameters struct {
	// TODO
}

type IsoClientConnection struct {
	parameters      IsoConnectionParameters
	state           IsoConnectionState
	readTimeoutInMs int
	nextReadTimeout int
	socket          *net.Conn

	cotpConnection *CotpConnection
	presentation   *Presentation
	session        *Session
	acseConnection *AcseConnection

	recevieBuffer []byte
	sendBuffer    []byte

	cotpReadBuffer  []byte
	cotpWriteBuffer []byte
}
