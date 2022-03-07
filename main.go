package main

import (
	"net/http"
	"sktbind/process"
)

func main() {

	p := process.NewProcess()
	p.Startprocess()

	// select {}
	// 因為 client 採用 http.HandleFunc，所以要使用 ListenAdnServe 來服務
	http.ListenAndServe(":3000", nil)
}
