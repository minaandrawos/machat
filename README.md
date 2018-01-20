***MaChat***

----------
![MaChat in action](https://github.com/minaandrawos/machat/blob/master/machatarch.jpg)

An open source chat server implemented in pure Go.  The chat server listens on a TCP port and/or a websocket. It supports the following concepts:

 - Chat clients: A chat client represents the person who is sending and receiving messages
 - Chat rooms: A chat room is a place where multiple chat clients can be added. All clients connected to the same chat room will be able to see their messages

Pull requests are more than welcome

----------


**How it works?**

 - The MaChat server listens on a TCP server as well as a websocket server specified by you for
   incoming messages
 - A MaChat client starts the connection by:
	 -  Establishing a TCP or a websocket connection to the MaChat server port
	 -  Sending a JSON string message to indicate that chat client name as well as the name of the chat room that the client would like to join. The JSON format looks like this:
	 `{"name": "Mina", "room":"devroom"}`
     - Chat rooms can be shared between TCP clients and websocket clients based on how the chat API is initialized
	 
Once the name and chat room is sent to the chat server, the client can then proceed to send string messages via TCP or websockets,  the server will route the messages to the appropriate chat room and share them with other clients in the chat room. 


----------


**How to run the server?**

The MaChat server supports two user flags:
 - '-tcp' which is used to indicate the TCP server listening address. The default value is "localhost:8989"
 - '-ws' which is used to indicate the websocket server listening address. The default value is "localhost:8099"

Example command:

    <executablename> -tcp localhost:8989 -ws localhost:8099

'executablename' here is the name of the executable you build for the chat server. If you would like to run the code directly from the go code:

    #Run the chat server with the default port
    go run cmd/machat/machat.go
  
  There is an example TCP chat client application code provided which supports three user flags:
  

 - '-client' : for client name
 - '-room': for chat room name
 - '-address': for the TCP address of the chat server,  defaulted to localhost:8989

To run the example client application with client name 'Mina' and chat room 'dev':

    go run cmd/examplechatclient/chattcpclient.go -client Mina -room dev

  There is also an example for a websocket client application which supports four user flags:
 - '-client' : for client name
 - '-room': for chat room name
 - '-origin': for origin header field of the websocket 
 - '-address': for the TCP address of the chat server,  defaulted to localhost:8099  


----------


**What's checked in?**

 - *./gotcpchat.go*: entry point to the chat TCP server code
 - *./gowschat.go*: entry point to the chat websocket server code
 - *./chatapi/clients.go*: code to handle the chat clients, the code is exposed as an API to be used by gotcpchat.go
 - .*/chatapi/rooms.go*: code to handle chat rooms, the code is  exposed as an API to be used by gotcpchat.go
 - *./cmd/machat/machat.go*: main package for the MaChat TCP server, you need to build this file to get the executable for the chat server
 - *./cmd/examplechatclients/chattcpclient.go*: main package for an example tcp client that can connect to the chat server
 - *./cmd/examplechatclients/chatwebsocketclient.go*: main package for an example tcp client that can connect to the chat server

 ----------
 
 For other articles and tutorials, check my website at www.minaandrawos.com
