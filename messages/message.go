package messages

type Message struct {
	RId        int64       `json:"r_id"`
	FromUserId string      `json:"from_user_id"`
	ToUserId   string      `json:"to_user_id"`
	GroupId    string      `json:"group_id"`
	MsgId      int64       `json:"msg_id"`
	MimeType   string      `json:"mime_type"`
	Content    string      `json:"content"`
	SendTime   int64       `json:"send_time"`
	Type       MessageType `json:"type"`
}
