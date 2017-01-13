package gorocket

import (
	"fmt"
	"bytes"
	"net/http"
	"html"
)

type messagesResponse struct {
	statusResponse
	ChannelName string `json:"channel"`
	Messages []Message `json:"messages"`
}

type messageResponse struct {
	statusResponse
	ChannelName string `json:"channel"`
	Message Message `json:"message"`
}

type Message struct {
	Id string `json:"_id"`
	Rid string `json:"rid"`
	Text string `json:"msg"`
	Timestamp string `json:"ts"`
	User User `json:"u"`
}

type Page struct {
	Count int
}

// Sends a message to a channel. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (r *Rocket) Send(channel *Channel, msg string) error {
	body := fmt.Sprintf(`{ "channel": "%s", "text": "%s"}`, channel.Name, html.EscapeString(msg))
	request, _ := http.NewRequest("POST", r.getUrl() + "/api/v1/chat.postMessage", bytes.NewBufferString(body))

	response := new(messageResponse)

	return r.doRequest(request, response)
}

// Get messages from a channel. The channel id has to be not nil. Optionally a
// count can be specified to limit the size of the returned messages.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/history
func (r *Rocket) GetMessages(channel *Channel, page *Page) ([]Message, error) {
	u := fmt.Sprintf("%s/api/v1/channels.history?roomId=%s", r.getUrl(), channel.Id)

	if page != nil {
		u = fmt.Sprintf("%s&count=%d", u, page.Count)
	}

	request, _ := http.NewRequest("GET", u, nil)
	response := new(messagesResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}