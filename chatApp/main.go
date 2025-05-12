package main

import (
	"fmt"
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
	Clients:    make(map[string]*Client),
	Broadcast:  make(chan Message),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-manager.Register:
			manager.Clients[client.ID] = client
			fmt.Println("User Connected:", client.ID)
		case client := <-manager.Unregister:
			if _, ok := manager.Clients[client.ID]; ok {
				close(client.Send)
				delete(manager.Clients, client.ID)
				fmt.Println("User Disconnected:", client.ID)
			}
		case message := <-manager.Broadcast:
			fmt.Println("Broadcasting from", message.Sender, "to", message.Receiver)
			if message.Receiver == "all" {
				for _, client := range manager.Clients {
					client.Send <- message
				}
			} else {
				if recip, ok := manager.Clients[message.Receiver]; ok {
					recip.Send <- message
				}
			}
		}
	}
}

func (c *Client) Read() {
	defer func() {
		manager.Unregister <- c
		c.Conn.Close()
	}()
	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		manager.Broadcast <- msg
	}
}

func (c *Client) Write() {
	defer c.Conn.Close()
	for msg := range c.Send {
		fmt.Println("Sending message to", c.ID)
		err := c.Conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("Write error for", c.ID, ":", err.Error())
			break
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade", http.StatusInternalServerError)
		return
	}
	client := &Client{
		ID:   userID,
		Conn: conn,
		Send: make(chan Message),
	}

	manager.Register <- client

	go client.Read()
	go client.Write()
}

func main() {
	go manager.Start()

	http.HandleFunc("/ws", wsHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Server Started at :8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}

// 192.168.1.231
