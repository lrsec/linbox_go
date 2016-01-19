package messages

type OfflineInfo struct {
	UserId string `json:"user_id"`
	Reason string `json:"reason"`
}
