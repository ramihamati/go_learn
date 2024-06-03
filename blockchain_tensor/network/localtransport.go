package network

import "sync"

type LocalTransport struct {
	addr      NetAddr
	peers     map[NetAddr]*LocalTransport
	lock      sync.RWMutex
	consumeCh chan RPC
}
