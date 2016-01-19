package messages

type PullOldMsgRequest struct {
	RId               int64       `json:"r_id"`
	UserId            string      `json:"user_id"`
	RemoteId          string      `json:"remote_id"`
	GroupId           string      `json:"group_id"`
	MaxMsgId          int64       `json:"max_msg_id"`
	MinMsgId          int64       `json:"min_msg_id"`
	Limit             int         `json:"limit"`
	Type              MessageType `json:"type"`
	MaxMsgIdInRequest int64       `json:"max_msg_id_in_request"`
}
