package backend

import (
	"database/sql"
	"encoding/json"
	"se/database"
)

type Message struct {
	fn *database.MessageDB
}

// NewMessage init database
func NewMessage(db *sql.DB) (m *Message) {
	m = new(Message)
	m.fn = database.MessageDBInit(db)
	return
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

func (m *Message) GetAll() string {
	ms := m.fn.GetAll()

	str, err := json.Marshal(ms)
	if err != nil {
		return err.Error()
	}
	return string(str)
}
