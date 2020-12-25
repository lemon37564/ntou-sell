package backend

import (
	"encoding/json"
	"se/database"
)

// Message is a module that handle messages
type Message struct {
	fn *database.Data
}

// MessageInit inits database
func MessageInit(data *database.Data) (m *Message) {
	return &Message{fn: data}
}

// AddMessage add single message between two users
func (m *Message) AddMessage(senderUID, recieverUID int, text string) string {

	if senderUID == recieverUID {
		return "cannot send message to yourself!"
	}

	err := m.fn.AddMessage(senderUID, recieverUID, text)
	if err != nil {
		return err.Error()
	}

	return "ok"
}

// GetMessages return all messages between two users
func (m *Message) GetMessages(senderUID, recieverUID int, ascend bool) string {

	if senderUID == recieverUID {
		return "cannot get message which sent to yourself!"
	}

	ms := m.fn.GetMessages(senderUID, recieverUID, ascend)

	str, err := json.Marshal(ms)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

// GetAll return all messages(for debugging only)
func (m *Message) GetAll() string {
	ms := m.fn.GetAllMessages()

	str, err := json.Marshal(ms)
	if err != nil {
		return err.Error()
	}
	return string(str)
}
