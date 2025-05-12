package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan Message
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var manager = ClientManager{
	Clients: make(map[string]*Client),
	Broadcast : make(chan Message),
	Register:  make(chan *Client),
	Unregister: make(chan *Client),
}
