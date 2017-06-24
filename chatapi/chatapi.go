package chatapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ChatAPI struct {
	rooms map[string]*Room
	*sync.RWMutex
}

type clientInfo struct {
	Room string `json:"room"`
	Name string `json:"name"`
}

//New start a new instance of the new chat api
func New() *ChatAPI {
	api := &ChatAPI{
		rooms:   make(map[string]*Room),
		RWMutex: new(sync.RWMutex),
	}
	//handle shutdown
	go func() {
		// Handle SIGINT and SIGTERM.
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		fmt.Println("Closing tcp connection")
		api.RLock()
		defer api.RUnlock()
		for name, r := range api.rooms {
			log.Printf("Closing room %s \n", name)
			close(r.Quit)
		}
		os.Exit(0)
	}()

	return api
}

//AddClient adds a new client to the chat server. Expects a JSON file
func (cAPI *ChatAPI) AddClient(c io.ReadWriteCloser) {
	cinfo := new(clientInfo)
	if err := json.NewDecoder(c).Decode(cinfo); err != nil {
		log.Printf("Could not create chat room, invalid chat room name, error: %s \n", err)
	} else {
		cAPI.handleClient(cinfo, c)
	}
}

func (cAPI *ChatAPI) handleClient(cinfo *clientInfo, c io.ReadWriteCloser) {
	cAPI.Lock()
	defer cAPI.Unlock()
	r, ok := cAPI.rooms[cinfo.Room]
	if !ok {
		r = CreateRoom(cinfo.Room)
	}
	r.AddClient(c, cinfo.Name)
	cAPI.rooms[cinfo.Room] = r
}
