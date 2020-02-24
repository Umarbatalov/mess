package model

type Message struct {
	Sender    int    `json:"sender"`
	Receiver  int    `json:"receiver"`
	Data      string `json:"data"`
}

func NewMessage() *Message {
	return &Message{}
}
