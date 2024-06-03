package network

type NetAddr string

type RPC struct {
	Payload []byte
	From    NetAddr
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
