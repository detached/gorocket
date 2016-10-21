package gorocket

type logonResponse struct {
	statusResponse
	Data struct {
		Token string `json:"authToken"`
		UserId string `json:"userId"`
	     } `json:"data"`
}

type roomsResponse struct {
	statusResponse
	Rooms []Room `json:"rooms"`
}

type Room struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	Topic string `json:"topic"`
	MessageCount int `json:"msgs"`
	UserNames []string `json:"usernames"`

	Owner struct {
		Id string `json:"_id"`
		UserName string `json:"username"`
	      } `json:"u"`

	ReadOnly bool `json:"ro"`
	Timestamp string `json:"ts"`
	T string `json:"t"`
	UpdatedAt string `json:"_updatedAt"`
	LastMessage string `json:"lm"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type message struct {
	Text string `json:"msg"`
}