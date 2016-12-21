package chatapi

import (
	"bufio"
	"io"
	"log"
)

//client type represents a chat client
type client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

/*
	StartClient starts a chat client. This function uses the channel generator pattern
	first argument is the main channel of chat room which the client belongs to. The client will be sending messages to this channel
	second argument is a readwritercloser representing a connection at which the client is communicating
	third argument is a quit channel. If a signal is passed through this channel, the client closes.
*/
func StartClient(name string, msgCh chan<- string, cn io.ReadWriteCloser, roomName string) (chan<- string, <-chan struct{}) {
	c := new(client)
	c.Reader = bufio.NewReader(cn)
	c.Writer = bufio.NewWriter(cn)
	c.wc = make(chan string)
	channelDone := make(chan struct{})

	//setup the reader. When the client sends a message, we will send it to the chat room
	go func() {
		scanner := bufio.NewScanner(c.Reader)
		for scanner.Scan() {
			msg := name + ":" + scanner.Text() + "\n"
			log.Printf("%s|%s", roomName, msg)
			msgCh <- msg
		}
		close(channelDone)
		cn.Close()
	}()

	//setup the writer
	c.writeMonitor()

	return c.wc, channelDone
}

func (c *client) writeMonitor() {
	go func() {
		for s := range c.wc {
			c.WriteString(s)
			c.Flush()
		}
	}()
}
