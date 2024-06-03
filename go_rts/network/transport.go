package network

import "log"

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(transport Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}

func KafkaStart() {
	log.Println("starting kafka")
}
