package database

import (
	"database/sql"
	"log"
)

const messageTable = `
CREATE TABLE IF NOT EXISTS message(
	message_id int NOT NULL,
	sender_uid int NOT NULL,
	receiver_uid int NOT NULL,
	text varchar(128) NOT NULL,
	PRIMARY KEY(message_id),
	FOREIGN KEY(sender_uid) REFERENCES user,
	FOREIGN KEY(receiver_uid) REFERENCES user
);`

// Messages struct contain all message with contactor name
type Messages struct {
	ContactorName string
	Content       []message
}

// message struct store data of a single message
type message struct {
	// Status is "s" or "r" to represent sender or receiver
	Status string
	Text   string
}

// MessID is a struct only for getting all messages(debugging)
type MessID struct {
	messageID   int
	senderUID   int
	receiverUID int
	text        string
}

type messageStmt struct {
	all     *sql.Stmt
	add     *sql.Stmt
	getNew  *sql.Stmt
	getOld  *sql.Stmt
	maxID   *sql.Stmt
	getName *sql.Stmt
}

func messagePrepare(db *sql.DB) *messageStmt {
	var err error
	message := new(messageStmt)

	const (
		all    = "SELECT * FROM message;"
		add    = "INSERT INTO message VALUES(?,?,?,?);"
		getOld = `
			SELECT text, sender_uid
			FROM message
			WHERE (sender_uid=? AND receiver_uid=?) OR (receiver_uid=? AND sender_uid=?)
			ORDER BY message_id ASC;
			`
		getNew = `
			SELECT text, sender_uid
			FROM message
			WHERE (sender_uid=? AND receiver_uid=?) OR (receiver_uid=? AND sender_uid=?)
			ORDER BY message_id DESC;
			`
		maxID   = "SELECT max(message_id) FROM message;"
		getName = "SELECT name FROM user WHERE uid=? AND uid>0;"
	)

	if message.all, err = db.Prepare(all); err != nil {
		log.Println(err)
	}

	if message.add, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if message.getOld, err = db.Prepare(getOld); err != nil {
		log.Println(err)
	}

	if message.getNew, err = db.Prepare(getNew); err != nil {
		log.Println(err)
	}

	if message.maxID, err = db.Prepare(maxID); err != nil {
		log.Println(err)
	}

	if message.getName, err = db.Prepare(getName); err != nil {
		log.Println(err)
	}

	return message
}

// AddMessage record a new message between two users
func (dt Data) AddMessage(senderUID, receiverUID int, messageText string) error {
	var mID int

	rows, err := dt.message.maxID.Query()
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

	_, err = dt.message.add.Exec(mID, senderUID, receiverUID, messageText)
	return err
}

// GetMessages return all messge between two users
func (dt Data) GetMessages(localUID, remoteUID int, ascend bool) Messages {

	var rows *sql.Rows
	var err error

	if ascend {
		rows, err = dt.message.getNew.Query(localUID, remoteUID, localUID, remoteUID)
	} else {
		rows, err = dt.message.getOld.Query(localUID, remoteUID, localUID, remoteUID)
	}

	if err != nil {
		log.Println(err)
		return Messages{}
	}

	var (
		all   []message
		ms    message
		tmpID int
	)

	for rows.Next() {
		err = rows.Scan(&ms.Text, &tmpID)
		if err != nil {
			log.Println(err)
			return Messages{}
		}

		// local is sender or receiver
		if tmpID == localUID {
			ms.Status = "s"
		} else {
			ms.Status = "r"
		}

		all = append(all, ms)
	}

	return Messages{ContactorName: dt.getName(remoteUID), Content: all}
}

// GetAll return all messages (debugging only)
func (dt Data) GetAll() (all []MessID) {

	rows, err := dt.message.all.Query()
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

func (dt Data) getName(uid int) (name string) {

	rows, err := dt.message.getName.Query(uid)
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
