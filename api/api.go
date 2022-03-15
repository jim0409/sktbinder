package api

import (
	"io"
	"log"
	"net/http"
)

func ApiServer() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/kickout", kickout)
}

func login(res http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(res, "ok")
	if err != nil {
		log.Fatal(err)
	}
}

func kickout(res http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(res, "ok")

	if err != nil {
		log.Fatal(err)
	}
}
