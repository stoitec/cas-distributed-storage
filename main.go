package main

import (
	"log"

	"github.com/stoitec/cas-distributed-storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		Decoder:    p2p.DefaultDecoder{},
		HandshakeFunc: func(p p2p.Peer) error {
			return nil
		},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	// fmt.Println("Ohayo mina")
}
