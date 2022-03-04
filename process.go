package main

import "log"

type Process struct {
	Cm *ClientManager
}

func newProcess() *Process {
	// TODO: ...
	// initModule()
	return &Process{}
}

func (p Process) startprocess() {
	f := func(msg []byte, hub *ClientManager) error {
		log.Printf("每次處理訊息都會跑這行!!\n")
		log.Printf("Add Client to pools")
		log.Printf("%v\n", string(msg))
		return nil
	}
	hub := initClientCenter(f)
	startWebsocket(hub)
}
