package main

import "github.com/gorilla/websocket"

// type Client struct {
// 	ID        uint64 // 客户端ID
// 	connect   websocket.Connect
// 	SendChan  chan []byte
// 	haveLogin bool // 是否已登录

// Client : a middleman between the websocket connection and the hub.
type Client struct {
	conn      *websocket.Conn // The websocket connection.
	send      chan []byte     // Buffered channel of outbound messages.
	haveLogin bool            // check whether the client is Validated
	lastAck   int             // the last ack from client .. TODO: ..
}

type ClientManager struct {
	clients   map[*Client]bool
	Broadcast chan []byte
	// OnMessage : trigger while receive and client message
	/*
		e.g.
		1. LoginValidation
		2. last ack sequence number
		3. ...
	*/
	OnMessage OnMessageFunc

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}
