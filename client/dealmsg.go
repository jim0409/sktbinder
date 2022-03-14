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
			workRun(id, d.recvMsgChan)
		}(i)
	}
}

func workRun(id int, msg chan *MsgPkg) {
	fmt.Printf("enter worker id %d\n", id)
	// for {
	// 	select {
	// 	// TODO: 考慮 global goroutine 停止
	// 	case m, ok := <-msg:
	// 		if !ok {
	// 			return
	// 		}
	// 		go func(pkg MsgPkg) {
	// 			fmt.Println(pkg)
	// 		}(m)
	// 		// TODO: 考慮 client 關閉連線
	// 	}
	// }
	for m := range msg {
		fmt.Printf("inside worker pool execute %d.. %v\n", m.MessageType, string(m.Message))
	}
}
