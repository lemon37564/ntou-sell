package backend

import (
	"encoding/json"
	"se/database"
	"strconv"
)

// AddMessage add single message between two users
func AddMessage(senderUID int, rawReceiverUID, text string) (string, error) {
	receiverUID, err := strconv.Atoi(rawReceiverUID)
	if err != nil {
		return "cannot convert " + rawReceiverUID + " into integer", err
	}

	if senderUID == receiverUID {
		return "failed", beError{text: "cannot send message to yourself!"}
	}

	err = database.AddMessage(senderUID, receiverUID, text)
	if err != nil {
		return "failed", err
	}

	return "ok", nil
}

// GetMessages return all messages between two users
func GetMessages(localUID int, rawRemoteUID, ascend string) (string, error) {
	remoteUID, err := strconv.Atoi(rawRemoteUID)
	if err != nil {
		return "cannot convert " + rawRemoteUID + " into integer", err
	}

	if localUID == remoteUID {
		return "failed", beError{text: "cannot get message which sent to yourself!"}
	}

	ms := database.GetMessages(localUID, remoteUID, ascend == "true")

	str, err := json.Marshal(ms)
	if err != nil {
		return "failed", err
	}
	return string(str), nil
}

// GetAll return all messages(for debugging only)
func GetAllMessages() string {
	ms := database.GetAllMessages()

	str, err := json.Marshal(ms)
	if err != nil {
		return err.Error()
	}
	return string(str)
}
