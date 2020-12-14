package database

import (
	"database/sql"
	"log"
)

const messageTable = `CREATE TABLE message(
						message_id int NOT NULL,
						sender_uid int NOT NULL,
						receiver_uid int NOT NULL,
						text varchar(128) NOT NULL,
						PRIMARY KEY(message_id),
						FOREIGN KEY(sender_uid) REFERENCES user,
						FOREIGN KEY(receiver_uid) REFERENCES user
					);`

// MessageDB contain funcions to use
type MessageDB struct {
	all    *sql.Stmt
	add    *sql.Stmt
	getnew *sql.Stmt
	getold *sql.Stmt
	maxid  *sql.Stmt
}

// Message struct store data of a single Message
type Message struct {
	SenderName   string
	ReceiverName string
	Text         string
}

type MessID struct {
	MessageID   int
	SenderUID   int
	ReceiverUID int
	Text        string
}

// MessageDBInit prepare function for database using
func MessageDBInit(db *sql.DB) *MessageDB {
	var err error
	message := new(MessageDB)

	message.all, err = db.Prepare("SELECT * FROM message;")
	if err != nil {
		panic(err)
	}

	message.add, err = db.Prepare("INSERT INTO message VALUES(?,?,?,?);")
	if err != nil {
		panic(err)
	}

	message.getold, err = db.Prepare(`
		SELECT T.name, S.name, text
		FROM user as T, user as S, message
		WHERE ((sender_uid=? AND receiver_uid=?) OR (receiver_uid=? AND sender_uid=?))
			AND (sender_uid=T.uid AND receiver_uid=S.uid)
		ORDER BY message_id ASC;
	`)
	if err != nil {
		panic(err)
	}

	message.getnew, err = db.Prepare(`
		SELECT T.name, S.name, text
		FROM user as T, user as S, message
		WHERE ((sender_uid=? AND receiver_uid=?) OR (receiver_uid=? AND sender_uid=?))
			AND (sender_uid=T.uid AND receiver_uid=S.uid)
		ORDER BY message_id DESC;
	`)
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
func (m *MessageDB) GetMessages(senderUID, receiverUID int, ascend bool) (all []Message) {

	var rows *sql.Rows
	var err error

	if ascend {
		rows, err = m.getnew.Query(senderUID, receiverUID, senderUID, receiverUID)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		rows, err = m.getold.Query(senderUID, receiverUID, senderUID, receiverUID)
		if err != nil {
			log.Println(err)
			return
		}
	}

	var mess Message
	for rows.Next() {
		err = rows.Scan(&mess.SenderName, &mess.ReceiverName, &mess.Text)
		if err != nil {
			log.Println(err)
			return
		}

		all = append(all, mess)
	}

	return
}

// GetAll return all messages (debugging only)
func (m *MessageDB) GetAll() (all []MessID) {

	rows, err := m.all.Query()
	if err != nil {
		log.Println(err)
		return
	}

	var mess MessID
	for rows.Next() {
		err = rows.Scan(&mess.MessageID, &mess.SenderUID, &mess.ReceiverUID, &mess.Text)
		if err != nil {
			log.Println(err)
			return
		}

		all = append(all, mess)
	}

	return
}
