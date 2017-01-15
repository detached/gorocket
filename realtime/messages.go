package realtime

import (
	"github.com/gopackage/ddp"
	"github.com/detached/gorocket/api"
	"time"
	"log"
	"fmt"
)

const (
	// RocketChat doesn't send the `added` event for new messages by default, only `changed`.
	send_added_event = true
	default_buffer_size = 100
)

type MessageListener interface {
	Notify(*api.Message)
}

func fromEvent(update ddp.Update) []api.Message {
	args := update["args"].([]interface{})

	if args == nil {
		log.Println("Got update without messages")
		return make([]api.Message, 0)
	}

	log.Printf("Got %v updates", len(args))
	messages := make([]api.Message, len(args))

	for i, arg := range args {
		mData := arg.(map[string]interface {})
		messages[i] = api.Message{
			Id:        mData["_id"].(string),
			ChannelId: mData["rid"].(string),
			Text:      mData["msg"].(string),
			Timestamp: toTimestamp(mData["ts"].(map[string]interface{})),
			User:      toUser(mData["u"].(map[string]interface{}))}
	}

	return messages
}
func toTimestamp(tsData map[string]interface{}) string {
	return fmt.Sprintf("%f", tsData["$date"].(float64))
}

func toUser(userData map[string]interface{}) api.User {
	return api.User{
		Id:      userData["_id"].(string),
		UserName:userData["username"].(string)}
}

type messageExtractor struct {
	messageChannel chan api.Message
}

func (u messageExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	log.Printf("CollectionUpdate for %s, operation %s, id %s", collection, operation, id)
	if operation == "update" {
		msgs := fromEvent(doc)
		for _, m := range msgs {
			u.messageChannel <- m
		}
	}
}

func getNextMessageId() string {
	return time.Now().Format("20170102150405")
}

func (c *Client) SendMessage(channel *api.Channel, text string) (api.Message, error) {
	m := api.Message{
		Id:        getNextMessageId(),
		ChannelId: channel.Id,
		Text:      text}
	_, err := c.ddp.Call("sendMessage", m)
	return m, err
}

// Subscribes to the message updates of a channel
// The channel has to be buffered and is not allowed to
func (c *Client) SubscribeToMessageStream(channel *api.Channel) (chan api.Message, error) {

	if err := c.ddp.Sub("stream-room-messages", channel.Id, send_added_event); err != nil {
		return nil, err
	}

	m := make(chan api.Message, default_buffer_size)
	c.ddp.CollectionByName("stream-room-messages").AddUpdateListener(messageExtractor{messageChannel:m})

	return m, nil
}
