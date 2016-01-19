package messages

type SyncUnreadRequest struct {
	RId      int64       `json:"r_id"`
	UserId   string      `json:"user_id"`
	RemoteId string      `json:"remote_id"`
	GroupId  string      `json:"group_id"`
	Offset   int64       `json:"offset"`
	Limit    int         `json:"limit"`
	Type     MessageType `json:"type"`
}
