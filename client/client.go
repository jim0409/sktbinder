package client

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/xid"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

// Client : a middleman between the websocket connection and the hub.
type Client struct {
	ClientID string
	conn     *websocket.Conn // The websocket connection.
	send     chan []byte     // Buffered channel of outbound messages.
	login    bool
	close    chan struct{}
}

func newClient(unregistChan chan *Client, recvMsg chan []byte, Id string, conn *websocket.Conn) *Client {
	client := &Client{
		ClientID: xid.New().String(),
		conn:     conn,
		send:     make(chan []byte, 256),
		close:    make(chan struct{}),
		login:    true,
	}

	go client.writePump()
	go client.readPump()
	go client.loginTimer()

	return client
}

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

// timeout second for 10s
func (c *Client) loginTimer() {
	timer := time.NewTimer(10 * time.Second)

	select {
	case <-c.close:
		return
	case <-timer.C:
		if !c.login {
			c.Close()
		}
	}
}
func (c *Client) Close() {
	c.conn.Close()
}
