package client

import "fmt"

const workNum = 10

type DealMsg struct {
	recvMsgChan     chan *MsgPkg
	clientCloseChan chan *Client
}

func NewDealMsg(recvMsgChan chan *MsgPkg, clientCloseChan chan *Client) *DealMsg {
	d := &DealMsg{
		recvMsgChan:     recvMsgChan,
		clientCloseChan: clientCloseChan,
	}
	d.Run()

	return d
}

// Run: 作為處理 Message package 的一個後台執行 worker pool
func (d *DealMsg) Run() {
	for i := 0; i < workNum; i++ {
		go func(id int) {
			workRun(id, d.recvMsgChan, d.clientCloseChan)
		}(i)
	}
}

/*
id: 識別 worker 的狀態
msg: 傳送給 worker 要做的 message
closeEvt: 傳送 client 要關閉的事件
*/
func workRun(id int, msg chan *MsgPkg, closeEvt chan *Client) {
	fmt.Printf("enter worker id %d\n", id)
	for {
		select {
		// TODO: 考慮 global goroutine 停止
		case m, ok := <-msg:
			if !ok {
				return
			}
			go func(pkg *MsgPkg) {
				fmt.Printf("inside worker pool execute %d.. %v\n", m.MessageType, string(m.Message))
			}(m)

		// TODO: 考慮 server 關閉 client 連線前要提示 client
		case evt := <-closeEvt:
			evt.Close()
		}
	}
}
