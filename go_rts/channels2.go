package main

import (
	"log"
	"time"
)

type C2Hub struct {
	Peers map[string]*NodeConnection
}

type NodeConnection struct {
	Node *Node
	Conn chan string
}

type Node struct {
	Id   string
	Conn chan string
}

func NewHub() *C2Hub {
	return &C2Hub{
		Peers: make(map[string]*NodeConnection),
	}
}

func (hub *C2Hub) Join(peer *Node) {
	nodeConnection := &NodeConnection{
		Node: peer,
		Conn: make(chan string),
	}
	hub.Peers[peer.Id] = nodeConnection
	nodeConnection.Connect()
}

func (hub *C2Hub) Broadcast(message string) {
	for _, peer := range hub.Peers {
		peer.Conn <- message
	}
}

func (node *NodeConnection) Connect() {
	log.Println("Client ", node.Node.Id, " connecting")
	go func() {
		for {
			select {
			case message := <-node.Conn:
				log.Println("Client ", node.Node.Id, " received a message ", message)
				log.Println(message)
				break
			}

		}
	}()
}

func StartChannels2() {
	hub := NewHub()

	node1 := &Node{
		Id: "1",
	}

	node2 := Node{
		Id: "2",
	}

	hub.Join(node1)
	hub.Join(&node2)

	hub.Broadcast("Hello")

	time.Sleep(10 * time.Second)
}
