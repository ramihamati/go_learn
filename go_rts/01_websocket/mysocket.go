package t1_websocket

// wscat -c ws://172.20.18.78:8085/ws
import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan *Payload
	register   chan *Client
	unregister chan *Client
}

type Payload struct {
	PayloadType    string
	PayloadContent []byte
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan *Payload
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				client.send <- message
			}
			//for client := range h.clients {
			//	select {
			//	//case client.send <- message:
			//
			//	default:
			//		delete(h.clients, client)
			//		close(client.send)
			//	}
			//}
		}
	}
}

func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {

		_, message, err := c.conn.ReadMessage()

		if err != nil {
			break
		}

		c.hub.broadcast <- &Payload{
			PayloadType:    "text",
			PayloadContent: message,
		}
		log.Println("client: just read something")
	}
}

func (c *Client) write() {
	defer c.conn.Close()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}
			err := c.conn.WriteMessage(websocket.TextMessage, message.PayloadContent)
			if err != nil {
				return
			}
			err1 := c.conn.WriteMessage(websocket.TextMessage, []byte("hi"))
			if err1 != nil {
				return
			}
		}
	}
}

func StartWebSocketService() {
	log.Println("Hello")
	hub := Hub{
		broadcast:  make(chan *Payload),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err)
			return
		}

		client := Client{
			hub:  &hub,
			conn: conn,
			send: make(chan *Payload),
		}

		hub.register <- &client
		go client.read()
		go client.write()
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}
