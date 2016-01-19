package messages

type SyncUnreadResponse struct {
	RId      int64           `json:"r_id"`
	UserId   string          `json:"user_id"`
	RemoteId string          `json:"remote_id"`
	GroupId  string          `json:"group_id"`
	Type     MessageType     `json:"type"`
	Unreads  []UnreadMessage `json:"unreads"`
	Offset   int64           `json:"offset"`
	Status   int             `json:"status"`
	ErrMsg   string          `json:"err_msg"`
}
