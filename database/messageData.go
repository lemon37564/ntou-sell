package database

import (
	"database/sql"
	"log"
)

const messageTable = `CREATE TABLE message(
						message_id int NOT NULL,
						sender_uid int NOT NULL,
						receiver_uid int NOT NULL,
						message varchar(128) NOT NULL,
						PRIMARY KEY(message_id),
						FOREIGN KEY(sender_uid) REFERENCES user,
						FOREIGN KEY(receiver_uid) REFERENCES user
					);`

// MessageDB contain funcions to use
type MessageDB struct {
	add   *sql.Stmt
	get   *sql.Stmt
	maxid *sql.Stmt
}

// Message struct store data of a single Message
type Message struct {
	SenderUID   int
	RecieverUID int
	MessageText string
}

// MessageDBInit prepare function for database using
func MessageDBInit(db *sql.DB) *MessageDB {
	var err error
	message := new(MessageDB)

	message.add, err = db.Prepare("INSERT INTO message VALUES(?,?,?,?);")
	if err != nil {
		panic(err)
	}

	message.get, err = db.Prepare("SELECT * FROM message WHERE sender_uid=? AND receiver_uid=? ORDER BY message_id DESC;")
	if err != nil {
		panic(err)
	}

	message.maxid, err = db.Prepare("SELECT max(message_id) FROM message;")
	if err != nil {
		panic(err)
	}

	return message
}

// AddMessage record a new message between two users
func (m *MessageDB) AddMessage(senderUID, receiverUID int, messageText string) error {
	var mID int

	rows, err := m.maxid.Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&mID)
		if err != nil {
			mID = 0
		}
	}

	mID++

	_, err = m.add.Exec(mID, senderUID, receiverUID, messageText)
	return err
}

// GetMessages return all messge between two users
func (m *MessageDB) GetMessages(senderUID, receiverUID int) (all []Message) {

	rows, err := m.get.Query(senderUID, receiverUID)
	if err != nil {
		log.Println(err)
		return
	}

	var mess Message
	for rows.Next() {
		err = rows.Scan(&mess.SenderUID, &mess.RecieverUID, &mess.MessageText)
		if err != nil {
			log.Println(err)
			return
		}

		all = append(all, mess)
	}

	return
}
