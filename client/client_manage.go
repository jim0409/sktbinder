package client

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/xid"
)

type ClientManager struct {
	Broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	clients map[string]*Client
	// OnMessage : trigger while receive and client message
	/*
		e.g.
		1. LoginValidation
		2. last ack sequence number
		3. ...
	*/
	OnMessage OnMessageFunc
}

func ClientCenter(onEvent OnMessageFunc) *ClientManager {
	return &ClientManager{
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
		OnMessage:  onEvent,
	}
}

// OnMessageFunc : 會檢查每次傳送訊息過來時的 msg []byte
type OnMessageFunc func(msg []byte, cm *ClientManager) error

func (h *ClientManager) Run() {
	for {
		select {
		case client := <-h.register:
			log.Printf("new client was add to map ... %v\n", client)
			log.Println("---------", client.ClientID)
			h.clients[client.ClientID] = client

			for i, j := range h.clients {
				log.Printf("%v___%v\n", i, j)
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client.ClientID]; ok {
				delete(h.clients, client.ClientID)
				close(client.send)
			}

		case message := <-h.Broadcast:
			for _, client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client.ClientID)
				}
			}
		}
	}
}

func SocketServer(cm *ClientManager, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// new client added ...
	client := newClient(cm.unregister, cm.Broadcast, xid.New().String(), conn)
	cm.register <- client

}

func StartServer(cm *ClientManager) {
	go cm.Run()
	http.HandleFunc("/conn", func(res http.ResponseWriter, r *http.Request) {
		SocketServer(cm, res, r)
	})
}
