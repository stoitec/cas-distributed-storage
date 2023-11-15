package p2p

import "errors"

// ErrInvalidHandshake is returned if the connection between the
// local and the remote node couldn't be established.
var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakeFunc func(Peer) error
