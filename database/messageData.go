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
	all          *sql.Stmt
	add          *sql.Stmt
	getnew       *sql.Stmt
	getold       *sql.Stmt
	maxid        *sql.Stmt
	getNameByUID *sql.Stmt
}

// Messages struct contain all message with contactor name
type Messages struct {
	ContactorName string
	Content       []message
}

// message struct store data of a single message
type message struct {
	// Status is 's' or 'r' to represent sender or receiver
	status rune
	text   string
}

// MessID is a struct only for getting all messages(debugging)
type MessID struct {
	messageID   int
	senderUID   int
	receiverUID int
	text        string
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
		SELECT text, sender_uid
		FROM message
		WHERE (sender_uid=? AND receiver_uid=?) OR (receiver_uid=? AND sender_uid=?)
		ORDER BY message_id ASC;
	`)
	if err != nil {
		panic(err)
	}

	message.getnew, err = db.Prepare(`
		SELECT text, sender_uid
		FROM message
		WHERE (sender_uid=? AND receiver_uid=?) OR (receiver_uid=? AND sender_uid=?)
		ORDER BY message_id DESC;
	`)
	if err != nil {
		panic(err)
	}

	message.maxid, err = db.Prepare("SELECT max(message_id) FROM message;")
	if err != nil {
		panic(err)
	}

	message.getNameByUID, err = db.Prepare("SELECT name FROM user WHERE uid=? AND uid>0;")
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
func (m *MessageDB) GetMessages(localUID, remoteUID int, ascend bool) Messages {

	var rows *sql.Rows
	var err error

	if ascend {
		rows, err = m.getnew.Query(localUID, remoteUID, localUID, remoteUID)
		if err != nil {
			log.Println(err)
			return Messages{}
		}
	} else {
		rows, err = m.getold.Query(localUID, remoteUID, localUID, remoteUID)
		if err != nil {
			log.Println(err)
			return Messages{}
		}
	}

	var all []message
	var ms message
	var tmpID int
	for rows.Next() {
		err = rows.Scan(&ms.text, &tmpID)
		if err != nil {
			log.Println(err)
			return Messages{}
		}

		// local is sender or receiver
		if tmpID == localUID {
			ms.status = 's'
		} else {
			ms.status = 'r'
		}

		all = append(all, ms)
	}

	return Messages{ContactorName: m.getName(localUID), Content: all}
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
		err = rows.Scan(&mess.messageID, &mess.senderUID, &mess.receiverUID, &mess.text)
		if err != nil {
			log.Println(err)
			return
		}

		all = append(all, mess)
	}

	return
}

func (m *MessageDB) getName(uid int) (name string) {

	rows, err := m.getNameByUID.Query(uid)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}
