package messages

import "bottled/utils"

type Message struct {
	SenderID   int
	ReceiverID int
	BottleID   int
	MessageID  int

	Text string

	Sequence int
}

func NewMessage(sid int, rid int, bid int, text string, sequence int) *Message {
	return &Message{
		SenderID:   sid,
		ReceiverID: rid,
		BottleID:   bid,
		Text:       text,
		Sequence:   sequence,
		MessageID:  utils.GenInt(0, 99999999),
	}
}
