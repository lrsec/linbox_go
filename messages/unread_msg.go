package messages

type UnreadMsg struct {
	UserId   string      `json:"user_id"`
	RemoteId string      `json:"remote_id"`
	GroupId  string      `json:"group_id"`
	Type     MessageType `json:"type"`
	MsgId    int64       `json:"msg_id"`
	Count    int64       `json:"count"`
	Msg      Message     `json:"msg"`
}
