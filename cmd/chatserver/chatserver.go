package main

import (
	"flag"
	"machat"
	"log"
)

func main() {
	addr := flag.String("a", "localhost:8989", "Address for the chat server to listen on")
	flag.Parse()
	if err := machat.RunTCP(*addr); err != nil {
		log.Fatalf("Could not listen on %s, error %s \n", *addr, err)
	}
}
