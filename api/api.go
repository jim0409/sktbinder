package api

import (
	"io"
	"log"
	"net/http"
)

func ApiServer() {
	http.HandleFunc("/api", kickout)
}

func kickout(res http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(res, "ok")

	if err != nil {
		log.Fatal(err)
	}
}
