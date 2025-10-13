package telegram

type Update struct {
	UpdateId int64    `json:"update_id"`
	Message  *Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	Id int64 `json:"id"`
}
