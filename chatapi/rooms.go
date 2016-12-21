package chatapi

import (
	"io"
	"log"
	"sync"
)

//Room type represents a chat room
type Room struct {
	name    string
	Msgch   chan string
	clients map[string]chan<- string
	//signals the quitting of the chat room
	Quit chan struct{}
	*sync.RWMutex
}

//CreateRoom starts a new chat room with name rname
func CreateRoom(rname string) *Room {
	r := &Room{
		name:    rname,
		Msgch:   make(chan string),
		RWMutex: new(sync.RWMutex),
		clients: make(map[string]chan<- string),
		Quit:    make(chan struct{}),
	}
	r.Run()
	return r
}

//AddClient adds a new client to the chat room
func (r *Room) AddClient(c io.ReadWriteCloser, clientname string) {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.clients[clientname]; ok {
		log.Printf("Client %s already exist in chat room %s, existing...", clientname, r.name)
		return
	}
	log.Printf("Adding client %s \n", clientname)
	wc, done := StartClient(clientname, r.Msgch, c, r.name)
	r.clients[clientname] = wc

	//remove client when done is signalled
	go func() {
		<-done
		r.RemoveClientSync(clientname)
	}()
}

//ClCount returns the number of clients in a chat room
func (r *Room) ClCount() int {
	return len(r.clients)
}

//RemoveClientSync removes a client from the chat room. This is a blocking call
func (r *Room) RemoveClientSync(name string) {
	r.Lock()
	defer r.Unlock()
	log.Printf("Removing client %s \n", name)
	delete(r.clients, name)
}

//Run runs a chat room
func (r *Room) Run() {
	log.Println("Starting chat room", r.name)
	//handle the chat room main message channel
	go func() {
		for msg := range r.Msgch {
			r.broadcastMsg(msg)
		}
	}()

	//handle when the quit channel is triggered
	go func() {
		<-r.Quit
		r.CloseChatRoomSync()
	}()
}

//CloseChatRoomSync closes a chat room. This is a blocking call
func (r *Room) CloseChatRoomSync() {
	r.Lock()
	defer r.Unlock()
	close(r.Msgch)
	for name := range r.clients {
		delete(r.clients, name)
	}
}

//fan out is used to distribute the chat message
func (r *Room) broadcastMsg(msg string) {
	r.RLock()
	defer r.RUnlock()
	for _, wc := range r.clients {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}
