package main

import t1websocket "RTS/01_websocket"

func main() {
	//StartWebSocketService()
	//Channels1Start()
	//StartChannels2()
	//KafkaStart()
	println("hello")

	//t1 := network.NewLocalTransport("1")
	//t2 := network.NewLocalTransport("2")
	//
	//t1.Connect(t2)
	//
	//t1.SendMessage("2", []byte("hello"))

	t1websocket.StartWebSocketService()
}
