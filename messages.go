package gorocket

import (
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"
)

type messagesResponse struct {
	statusResponse
	Messages []Message `json:"messages"`
}

type Message struct {
	Id string `json:"_id"`
	Rid string `json:"rid"`
	Text string `json:"msg"`
	Timestamp string `json:"ts"`
	User User `json:"u"`
}

type Page struct {
	Skip int
	Limit int
}

func (r *Rocket) Send(room *Room, msg string) error {
	b, _ := json.Marshal(&Message{Text:msg})
	request, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/rooms/%s/send", r.getUrl(), room.Id), bytes.NewBuffer(b))

	response := new(statusResponse)

	return r.doRequest(request, response)
}

func (r *Rocket) GetMessages(room *Room, page *Page) ([]Message, error) {
	u := fmt.Sprintf("%s/api/rooms/%s/messages", r.getUrl(), room.Id)

	if page != nil {
		u = fmt.Sprintf("%s?skip=%d&limit=%d", u, page.Skip, page.Limit)
	}

	request, _ := http.NewRequest("GET", u, nil)
	response := new(messagesResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}