package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOpts{
		ListenAddr: ":5000",
		HandshakeFunc: func(p Peer) error {
			return nil
		},
	}
	listenAddr := ":5000"
	tr := NewTCPTransport(tcpOpts)

	assert.Equal(t, tr.TCPTransportOpts.ListenAddr, listenAddr)
	assert.Nil(t, tr.ListenAndAccept())
}
