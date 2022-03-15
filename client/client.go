package client

import (
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

// Client : a middleman between the websocket connection and the hub.
type Client struct {
	ID    string
	conn  *websocket.Conn // The websocket connection.
	send  chan []byte     // Buffered channel of outbound messages.
	login bool
	close chan struct{}
}

func newClient(unregistChan chan *Client, recvMsg chan *MsgPkg, Id string, conn *websocket.Conn) *Client {
	client := &Client{
		ID:    Id,
		conn:  conn,
		send:  make(chan []byte, 256),
		close: make(chan struct{}),
		login: false, // 預設燈入為 false
	}

	go client.readPump(unregistChan, recvMsg)
	go client.writePump()
	go client.loginTimer()

	return client
}

func (c *Client) Login() {
	c.login = true
}

// readPump: Client to Server
func (c *Client) readPump(unregistChan chan *Client, recvMsgChan chan *MsgPkg) {
	log.Println("start read message")
	defer func() {
		log.Println("websocket Close")
		c.conn.Close()
		unregistChan <- c // 刪除 clientMap 裡面的 client
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
		// log.Printf("%v", string(message))
		log.Printf("client %v pass message to msg chain ..\n", c)

		// 將 message 送往 訊息通道 做處理
		id, _ := strconv.Atoi(c.ID)
		msg := &MsgPkg{
			MessageType: id,
			Message:     message,
		}
		recvMsgChan <- msg
	}
}

// writePump: Server to Client
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case <-c.close:
			return
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

// timeout second for 10s
func (c *Client) loginTimer() {
	log.Println("execute loginTimer")
	timer := time.NewTimer(10 * time.Second)

	select {
	case <-c.close:
		return
	case <-timer.C:
		if !c.login {
			log.Println("time out")
			c.Close()
		}
	}
}

// Close: 關閉 websocket 的連線
func (c *Client) Close() {
	c.conn.Close()
}
