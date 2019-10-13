package data

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// SocketRoom rooms for socket connection
type SocketRoom struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	Forward chan []byte
	// join is a channel for clients wishing to join the room.
	Join chan *SocketClient
	// leave is a channel for clients wishing to leave the room.
	Leave chan *SocketClient
	// clients holds all current clients in this room.
	Clients map[*SocketClient]bool
}

func (r *SocketRoom) Run() {
	for {
		select {
		case client := <-r.Join:
			// joining
			r.Clients[client] = true
		case client := <-r.Leave:
			// leaving
			delete(r.Clients, client)
			close(client.send)
		case msg := <-r.Forward:
			// forward message to all clients
			for client := range r.Clients {
				select {
				case client.send <- msg:
					// send the message
				default:
					// failed to send
					delete(r.Clients, client)
					close(client.send)
				}
			}
		}
	}
}

func (r *SocketRoom) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: remove in production
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	vars := mux.Vars(req)
	user := vars["user"]
	isHex := bson.IsObjectIdHex(user)

	if !isHex {
		return
	}

	client := &SocketClient{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
		user:   user,
	}
	r.Join <- client
	defer func() { r.Leave <- client }()
	go client.write()
	client.read()
}
