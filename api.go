package main

import (
	"io"
	"log"
	"net/http"
)

func apiServer() {
	http.HandleFunc("/api", kickout)
}

func kickout(res http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(res, "ok")

	if err != nil {
		log.Fatal(err)
	}
}
