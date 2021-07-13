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
	FOREIGN KEY(sender_uid) REFERENCES user ON DELETE CASCADE,
	FOREIGN KEY(receiver_uid) REFERENCES user ON DELETE CASCADE
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

var (
	msgAll     *sql.Stmt
	msgAdd     *sql.Stmt
	msgGetNew  *sql.Stmt
	msgGetOld  *sql.Stmt
	msgMaxID   *sql.Stmt
	msgGetName *sql.Stmt
)

func messagePrepare(db *sql.DB) {
	var err error

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

	if msgAll, err = db.Prepare(all); err != nil {
		log.Println(err)
	}

	if msgAdd, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if msgGetOld, err = db.Prepare(getOld); err != nil {
		log.Println(err)
	}

	if msgGetNew, err = db.Prepare(getNew); err != nil {
		log.Println(err)
	}

	if msgMaxID, err = db.Prepare(maxID); err != nil {
		log.Println(err)
	}

	if msgGetName, err = db.Prepare(getName); err != nil {
		log.Println(err)
	}
}

// AddMessage record a new message between two users
func AddMessage(senderUID, receiverUID int, messageText string) error {
	var mID int

	rows, err := msgMaxID.Query()
	if err != nil {
		return err
	}

	if rows.Next() {
		err = rows.Scan(&mID)
		if err != nil {
			mID = 0
		}
	}

	_, err = msgAdd.Exec(mID+1, senderUID, receiverUID, messageText)
	return err
}

// GetMessages return all messge between two users
func GetMessages(localUID, remoteUID int, ascend bool) Messages {

	var rows *sql.Rows
	var err error

	if ascend {
		rows, err = msgGetNew.Query(localUID, remoteUID, localUID, remoteUID)
	} else {
		rows, err = msgGetOld.Query(localUID, remoteUID, localUID, remoteUID)
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

	return Messages{ContactorName: getName(remoteUID), Content: all}
}

// GetAllMessages return all messages (debugging only)
func GetAllMessages() (all []MessID) {

	rows, err := msgAll.Query()
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

func getName(uid int) (name string) {

	rows, err := msgGetName.Query(uid)
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
