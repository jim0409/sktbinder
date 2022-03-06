package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

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

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Client) readPump() {
	log.Println("start read message")
	defer func() {
		c.conn.Close()
		log.Println("websocket Close")
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println("websocket.IsUnexpectedCloseError:", err)
			}
			log.Println("websocket read message have err")
			break
		}
		log.Printf("%v", string(message))
	}
}

func (h *ClientManager) Run() {
	for {
		select {
		case client := <-h.register:
			log.Printf("new client was add to map ... %v\n", client)
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

func ServeWs(hub *ClientManager, w http.ResponseWriter, r *http.Request) {
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
	log.Println("ServeWs is star")
	client := &Client{conn: conn, send: make(chan []byte, 256)}
	go client.writePump()

	client.readPump()
}

func initClientCenter(f OnMessageFunc) *ClientManager {
	return ClientCenter(f)
}

func startWebsocket(cm *ClientManager) {
	go cm.Run()
	http.HandleFunc("/conn", func(res http.ResponseWriter, r *http.Request) {
		ServeWs(cm, res, r)
	})

}
