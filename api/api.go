package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sktbind/client"
)

func ApiServer(cm *client.ClientManager) {
	ac := newClient(cm)
	http.HandleFunc("/login", ac.login)
	http.HandleFunc("/kickout", ac.kickout)
}

type ApiClient struct {
	Client *client.ClientManager
}

func newClient(cm *client.ClientManager) *ApiClient {
	return &ApiClient{
		Client: cm,
	}
}

type Login struct {
	ClientId string `json:"id"`
}

func (c *ApiClient) login(res http.ResponseWriter, r *http.Request) {
	var lg Login

	if err := json.NewDecoder(r.Body).Decode(&lg); err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	c.Client.GetClientMap()[lg.ClientId].Login()

	_, err := io.WriteString(res, "ok")
	if err != nil {
		log.Fatal(err)
	}
}

func (c *ApiClient) kickout(res http.ResponseWriter, r *http.Request) {
	var lg Login

	if err := json.NewDecoder(r.Body).Decode(&lg); err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	c.Client.GetClientMap()[lg.ClientId].Close()

	_, err := io.WriteString(res, "ok")

	if err != nil {
		log.Fatal(err)
	}
}
