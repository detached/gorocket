package main

import (
	"github.com/detached/gorocket/realtime"
	"github.com/detached/gorocket/api"
	"fmt"
)

func main() {
	// Create new realtime client
	c, _ := realtime.NewClient("127.0.0.1", "3000", false)

	// Login an existing user
	c.Login(&api.UserCredentials{Email: "reatimeTest@mail.com", Name: "realtime", Password: "realtime"})

	// Subscribe to the general channel
	general := api.Channel{Id: "GENERAL"}
	messageChannel, _ := c.SubscribeToMessageStream(&general)

	// Send some messages
	c.SendMessage(&general, "This")
	c.SendMessage(&general, "is")
	c.SendMessage(&general, "sparta!")

	// Get messages from channel
	fmt.Println(<-messageChannel)
	fmt.Println(<-messageChannel)
	fmt.Println(<-messageChannel)

	// Don't forget to close the client
	c.Close()
}
