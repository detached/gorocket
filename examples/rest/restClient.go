package main

import (
	"github.com/detached/gorocket/rest"
	"github.com/detached/gorocket/api"
	"log"
)

func main() {
	// Create a gorocket rest client
	c := gorocket.Client{Protocol: "http", Host: "127.0.0.1", Port: "3000"}

	// No login needed to get the server versions
	info, _ := c.GetServerInfo()
	log.Println("Server version: ", info.Version)

	// Login an existing user
	if err := c.Login(api.UserCredentials{Email: "test@mail.com", Password: "test"}); err != nil {
		log.Fatal("Error while login: ", err)
	}

	// Get all visible channels
	channels, _ := c.GetPublicChannels()
	log.Println("All visible channels: ", channels)

	// Join the general channel
	general := getChannelById(channels, "GENERAL")
	c.JoinChannel(general)

	// Get all joined channels
	joined, _ := c.GetJoinedChannels()
	log.Println("We are in the following channels: ", joined)

	// Send a message
	c.Send(general, "I am a go program!")

	// Get the last messages from the general channel
	messages, _ := c.GetMessages(general, nil)
	log.Println("Last messages: ", messages)

	// Leave the general channel
	c.LeaveChannel(general)

	// Logout the user
	c.Logout()
}

func getChannelById(channels []api.Channel, id string) *api.Channel {
	for _, c := range channels {
		if c.Id == id {
			return &c
		}
	}

	return nil
}