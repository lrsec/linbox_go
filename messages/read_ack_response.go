package messages

type ReadAckResponse struct {
	RId      int64       `json:"r_id"`
	UserId   string      `json:"user_id"`
	RemoteId string      `json:"remote_id"`
	GroupId  string      `json:"group_id"`
	Type     MessageType `json:"type"`
	Status   int         `json:"status"`
	ErrCode  string      `json:"err_code"`
}
