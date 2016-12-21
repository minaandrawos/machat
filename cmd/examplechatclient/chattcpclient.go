package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("User%d", rand.Intn(400))
	clientName := flag.String("client", name, "name of the client to connect")
	addr := flag.String("address", "localhost:8989", "address of the chat server")
	roomName := flag.String("room", "lounge", "name of the chat room")
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatal("Could not connect to the chat server ", err)
	}
	log.Println("Connected to the tcp chat system")

	defer conn.Close()
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		//connection is closed or there is an error, exit
		log.Println("Connection seem to be closed or error occure", scanner.Err())
		os.Exit(0)
	}()
	clientInfo := struct {
		Name string `json:"name"`
		Room string `json:"room"`
	}{*clientName, *roomName}
	data, _ := json.Marshal(clientInfo)
	conn.Write(data)
	fmt.Println("Start typing your messages")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, msg+"\n")
	}
}
