package realtime

import (
	"github.com/detached/ddp"
	"github.com/Jeffail/gabs"
	"github.com/detached/gorocket/api"
	"time"
	"log"
	"fmt"
)

const (
	// RocketChat doesn't send the `added` event for new messages by default, only `changed`.
	send_added_event    = true
	default_buffer_size = 100
)

// Send a message to a channel
func (c *Client) SendMessage(channel *api.Channel, text string) (api.Message, error) {
	m := api.Message{
		Id:        getNextMessageId(),
		ChannelId: channel.Id,
		Text:      text}
	_, err := c.ddp.Call("sendMessage", m)
	return m, err
}

// Subscribes to the message updates of a channel
// Returns a buffered channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/subscriptions/stream-room-messages/
func (c *Client) SubscribeToMessageStream(channel *api.Channel) (chan api.Message, error) {

	if err := c.ddp.Sub("stream-room-messages", channel.Id, send_added_event); err != nil {
		return nil, err
	}

	msgChannel := make(chan api.Message, default_buffer_size)
	c.ddp.CollectionByName("stream-room-messages").AddUpdateListener(messageExtractor{msgChannel, "update"})

	return msgChannel, nil
}

func getNextMessageId() string {
	return time.Now().Format("20170102150405")
}

func fromEvent(update ddp.Update) []api.Message {
	document, _ := gabs.Consume(update["args"])
	args, err := document.Children()

	if err != nil {
		log.Printf("Event arguments are in an unexpected format: %v", err)
		return make([]api.Message, 0)
	}

	messages := make([]api.Message, len(args))

	for i, arg := range args {
		messages[i] = api.Message{
			Id: stringOrZero(arg.Path("_id").Data()),
			ChannelId: stringOrZero(arg.Path("rid").Data()),
			Text:                             stringOrZero(arg.Path("msg").Data()),
			Timestamp:                        stringOrZero(arg.Path("ts.$date").Data()),
			User:      api.User{
				Id:       stringOrZero(arg.Path("u._id").Data()),
				UserName: stringOrZero(arg.Path("u.username").Data())}}
	}

	return messages
}
func stringOrZero(i interface{}) string {
	if i == nil {
		return ""
	}

	switch i.(type) {
	case string:
		return i.(string)
	case float64:
		return fmt.Sprintf("%f", i.(float64))
	default:
		return ""
	}
}

type messageExtractor struct {
	messageChannel chan api.Message
	operation      string
}

func (u messageExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	if operation == u.operation {
		msgs := fromEvent(doc)
		for _, m := range msgs {
			u.messageChannel <- m
		}
	}
}
