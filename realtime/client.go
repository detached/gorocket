package realtime
// Provides access to the Meteor ddp realtime API of Rocket.Chat

import (
	"github.com/detached/ddp"
	"fmt"
)

type Client struct {
	ddp      *ddp.Client
	loggedIn bool
}

// Creates a new instance and connects to the websocket.
func NewClient(host, port string, debug bool) (*Client, error) {
	c := new(Client)
	c.ddp = ddp.NewClient(fmt.Sprintf("ws://%v:%v/websocket", host, port), "http://" + host)

	if debug {
		c.ddp.SetSocketLogActive(true)
	}

	if err := c.ddp.Connect(); err != nil {
		return nil, err
	}

	return c, nil
}

// Closes the ddp session
func (c *Client) Close() {
	c.ddp.Close()
}