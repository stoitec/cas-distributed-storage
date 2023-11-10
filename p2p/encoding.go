package p2p

import (
	"encoding/gob"
	"errors"
	"io"
)

// ErrInvalidHandshake is returned if the connection between the
// local and the remote node couldn't be established.
var ErrInvalidHandshake = errors.New("invalid handshake")

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct {
}

func (dec GOBDecoder) Decode(r io.Reader, v *Message) error {
	return gob.NewDecoder(r).Decode(v)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Payload = buf[:n]
	return nil
}
