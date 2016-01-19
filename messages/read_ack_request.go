package messages

type ReadAckRequest struct {
	RId      int64       `json:"r_id"`
	UserId   string      `json:"user_id"`
	RemoteId string      `json:"remote_id"`
	GroupId  string      `json:"group_id"`
	MsgId    int64       `json:"msg_id"`
	Type     MessageType `json:"type"`
}
