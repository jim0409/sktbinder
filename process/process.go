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

func (p Process) Startprocess() {
	f := func(msg []byte, hub *client.ClientManager) error {
		log.Printf("每次處理訊息都會跑這行!!\n")
		log.Printf("Add Client to pools")
		log.Printf("%v\n", string(msg))
		return nil
	}
	cm := client.ClientCenter(f)
	client.WsServer(cm)
	api.ApiServer(cm)
}
