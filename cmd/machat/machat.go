package main

import (
	"flag"
	"log"
	"machat"
	"machat/chatapi"
)

func main() {
	addr := flag.String("a", "localhost:8989", "Address for the chat server to listen on")
	flag.Parse()
	api := chatapi.New()
	go func() {
		machat.RunWSChatWithExistingAPI("/mychat", "localhost:8099", 1024, 1024, api)
	}()
	if err := machat.RunTCPWithExistingAPI(*addr, api); err != nil {
		log.Fatalf("Could not listen on %s, error %s \n", *addr, err)
	}
}
