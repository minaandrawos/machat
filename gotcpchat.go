package machat

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/minaandrawos/machat/chatapi"
)

type chatRoomRWMutexMap struct {
	rooms map[string]*chatapi.Room
	*sync.RWMutex
}

type clientInfo struct {
	Room string `json:"room"`
	Name string `json:"name"`
}

//RunTCP will start hydra chat tcp server on the provided connection string
func RunTCP(connection string) error {
	l, err := net.Listen("tcp", connection)
	if err != nil {
		log.Println("Error connecting to chat client", err)
		return err
	}
	roomSync := chatRoomRWMutexMap{
		rooms:   make(map[string]*chatapi.Room),
		RWMutex: new(sync.RWMutex),
	}
	//handle shutdown
	go func() {
		// Handle SIGINT and SIGTERM.
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		l.Close()
		fmt.Println("Closing tcp connection")
		roomSync.RLock()
		defer roomSync.RUnlock()
		for name, r := range roomSync.rooms {
			log.Printf("Closing room %s \n", name)
			close(r.Quit)
		}
		os.Exit(0)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			break
		}
		go func() {
			cinfo := new(clientInfo)
			if err = json.NewDecoder(conn).Decode(cinfo); err != nil {
				log.Printf("Could not create chat room, invalid chat room name, error: %s \n", err)
			} else {
				handleClient(cinfo, &roomSync, conn)
			}
		}()
	}

	return err
}

func handleClient(cinfo *clientInfo, rMap *chatRoomRWMutexMap, c net.Conn) {
	rMap.Lock()
	defer rMap.Unlock()
	r, ok := rMap.rooms[cinfo.Room]
	if !ok {
		r = chatapi.CreateRoom(cinfo.Room)
	}
	r.AddClient(c, cinfo.Name)
	rMap.rooms[cinfo.Room] = r
}
