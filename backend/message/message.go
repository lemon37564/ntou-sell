package message

import (
	"database/sql"
	"encoding/json"
	"se/database"
)

type Message struct {
	messagedb *database.MessageDB
}

// NewMessage init database
func NewMessage(db *sql.DB) (m *Message) {
	m = new(Message)
	m.messagedb = database.MessageDBInit(db)
	return
}

// AddMessage add single message between two users
func (m *Message) AddMessage(senderUID, recieverUID int, text string) string {
	err := m.messagedb.AddMessage(senderUID, recieverUID, text)
	if err != nil {
		err.Error()
	}

	return "ok"
}

// GetMessages return all messages between two users
func (m *Message) GetMessages(senderUID, recieverUID int) string {
	pd := m.messagedb.GetMessages(senderUID, recieverUID)

	str, err := json.Marshal(pd)
	if err != nil {
		panic(err)
	}
	return string(str)
}
