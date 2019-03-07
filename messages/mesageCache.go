package messages

import (
	"bottled/database"
	"fmt"
)

type MessageCache struct {
	Messages    map[int]map[int][]*Message
	DB          *database.DatabaseConnection
	AllMessages map[int]*Message
}

func NewMessageCache(d *database.DatabaseConnection) *MessageCache {
	return &MessageCache{
		Messages: make(map[int]map[int][]*Message),
		DB:       d,
	}
}

func (mc *MessageCache) CreateMailbox(userID int, bottleID int) {
	//if user isn't in system, add user
	if mc.Messages[userID] == nil {
		mc.Messages[userID] = make(map[int][]*Message)
		mc.Messages[userID][bottleID] = []*Message{}
	}
}

func (mc *MessageCache) DeleteMessages(messageIDs ...int) {
	for ids := range messageIDs {
		mc.AllMessages[ids] = nil
	}
}

func (mc *MessageCache) StartConversation(message *Message) {
	mc.CreateMailbox(message.ReceiverID, message.BottleID)
	mc.CreateMailbox(message.SenderID, message.BottleID)

	mc.Messages[message.ReceiverID][message.BottleID] = append(mc.Messages[message.ReceiverID][message.BottleID], message)

}

func (mc *MessageCache) ContinueConversation(message *Message) {
	mc.Messages[message.ReceiverID][message.BottleID] = append(mc.Messages[message.ReceiverID][message.BottleID], message)
}

func (mc *MessageCache) GetNewMessages(userID int) []*Message {
	messages := []*Message{}

	for _, v := range mc.Messages[userID] {
		fmt.Printf("\n%v", len(v))

		for _, y := range v {
			if y != nil {
				fmt.Printf("\n%v", y)
				messages = append(messages, y)
			}
		}
	}

	return messages
}
