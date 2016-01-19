package messages

type SendMsgRequest struct {
	RId      int64   `json:"r_id"`
	UserId   string  `json:"user_id"`
	RemoteId string  `json:"remote_id"`
	GroupId  string  `json:"group_id"`
	Msg      Message `json:"msg"`
	Type     int     `json:"type"`
}
