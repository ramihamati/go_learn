package network

import (
	"log"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	lock      sync.RWMutex
	consumeCh chan RPC
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr *LocalTransport) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.peers[tr.addr] = tr
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}

func (t *LocalTransport) SendMessage(addr NetAddr, message []byte) {
	t.lock.Lock()
	defer t.lock.Unlock()
	peer, ok := t.peers[addr]
	if !ok {
		log.Fatalln("could not find peer ", addr)
	}

	peer.consumeCh <- RPC{
		Payload: message,
		From:    addr,
	}
}
