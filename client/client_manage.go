package client

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type ClientManager struct {
	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// 儲存 client 在記憶體內
	// clientMap map[string]*Client
	clientMap map[*Client]bool

	// 廣播訊息
	Broadcast chan []byte

	// 接收訊息
	RecvMsgChan chan *MsgPkg

	// 當 client 取消註冊時，做 Client 關閉動作
	ClientCloseChan chan *Client

	// 連線建立時觸發
	OnMessage OnMessageFunc

	// 處理訊息
	dealMsg interface{}
}

func ClientCenter(onEvent OnMessageFunc) *ClientManager {
	c := &ClientManager{
		Broadcast:   make(chan []byte),
		RecvMsgChan: make(chan *MsgPkg),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clientMap:   make(map[*Client]bool),
		OnMessage:   onEvent,
	}

	c.dealMsg = NewDealMsg(c.RecvMsgChan, c.ClientCloseChan)
	c.Run()

	return c
}

// OnMessageFunc : 會檢查每次傳送訊息過來時的 msg []byte
type OnMessageFunc func(msg []byte, cm *ClientManager) error

func (h *ClientManager) Run() {
	go func() {
		for {
			select {
			case client, ok := <-h.register:
				if !ok {
					return
				}
				h.clientMap[client] = true

			case client := <-h.unregister:
				if _, ok := h.clientMap[client]; ok {
					delete(h.clientMap, client)
					close(client.send)
					h.ClientCloseChan <- client
				}

			case message := <-h.Broadcast:
				for client, _ := range h.clientMap {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clientMap, client)
					}
				}
			}
		}
	}()
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
		return
	}

	// new client added ...
	// client := newClient(cm.unregister, cm.Broadcast, xid.New().String(), conn)
	// client := newClient(cm.unregister, cm.RecvMsgChan, xid.New().String(), conn)
	client := newClient(cm.unregister, cm.RecvMsgChan, "123", conn)
	cm.register <- client

}

func StartServer(cm *ClientManager) {
	http.HandleFunc("/conn", func(res http.ResponseWriter, r *http.Request) {
		SocketServer(cm, res, r)
	})
}
