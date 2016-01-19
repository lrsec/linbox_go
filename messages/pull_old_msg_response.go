package messages

type PullOldMsgResponse struct {
	RId      int64       `json:"r_id"`
	UserId   string      `json:"user_id"`
	RemoteId string      `json:"remote_id"`
	GroupId  string      `json:"group_id"`
	Msgs     []Message   `json:"msgs"`
	Status   int         `json:"status"`
	ErrMsg   string      `json:"err_msg"`
	Type     MessageType `json:"type"`

	// request 中携带的 max message id 信息，冗余仅用于简化客户端操作
	MaxMsgIdInRequest int64 `json:"max_msg_id_in_request"`
}
