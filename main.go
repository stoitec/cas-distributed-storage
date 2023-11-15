package main

import (
	"fmt"
	"log"

	"github.com/stoitec/cas-distributed-storage/p2p"
)

func onPeer(p p2p.Peer) error {
	fmt.Println("Doing some good shit during peering")
	p.Close()
	return nil
	// return fmt.Errorf("failed the onPeer func")
}
func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		Decoder:    p2p.DefaultDecoder{},
		HandshakeFunc: func(p p2p.Peer) error {
			return nil
		},
		OnPeer: onPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
