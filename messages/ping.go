package messages

type Ping struct {
	RId    int64  `json:"r_id"`
	UserId string `json:"user_id"`
}
