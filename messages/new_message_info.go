package messages

type NewMessageInfo struct {
	UserId   string      `json:"user_id"`
	RemoteId string      `json:"remote_id"`
	GroupId  string      `json:"group_id"`
	Type     MessageType `json:"type"`
}
