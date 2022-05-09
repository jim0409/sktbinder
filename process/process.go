package process

import (
	"log"
	"sktbind/api"
	"sktbind/client"
)

type Process struct {
	Cm *client.ClientManager
}

func NewProcess() *Process {
	// TODO: ...
	// initModule()
	return &Process{}
}

func (p *Process) Startprocess() {
	cm := client.ClientCenter(connectHandler)
	client.WsServer(cm)
	api.ApiServer(cm)
}

func connectHandler(msg []byte) error {
	// 1. decode jwt
	// 2. handle each time
	log.Printf("client connect to server")
	log.Printf("%v\n", string(msg))
	return nil
}
