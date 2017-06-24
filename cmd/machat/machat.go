package main

import (
	"flag"
	"log"

	"github.com/minaandrawos/machat"
	"github.com/minaandrawos/machat/chatapi"
)

func main() {
	tcpAddr := flag.String("tcp", "localhost:8989", "Address for the TCP chat server to listen on")
	wsAddr := flag.String("ws", "localhost:8099", "Address for the websocket chat server to listen on")
	flag.Parse()
	api := chatapi.New()
	go func() {
		machat.RunWSChatWithExistingAPI("/mychat", *wsAddr, 1024, 1024, api)
	}()
	if err := machat.RunTCPWithExistingAPI(*tcpAddr, api); err != nil {
		log.Fatalf("Could not listen on %s, error %s \n", *tcpAddr, err)
	}
}
