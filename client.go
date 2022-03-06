package main

import "github.com/gorilla/websocket"

// Client : a middleman between the websocket connection and the hub.
type Client struct {
	ClientID  string
	conn      *websocket.Conn // The websocket connection.
	send      chan []byte     // Buffered channel of outbound messages.
	haveLogin bool            // check whether the client is Validated
	lastAck   int             // the last ack from client .. TODO: ..
}

type ClientManager struct {
	Broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	clients map[string]*Client
	// OnMessage : trigger while receive and client message
	/*
		e.g.
		1. LoginValidation
		2. last ack sequence number
		3. ...
	*/
	OnMessage OnMessageFunc
}
