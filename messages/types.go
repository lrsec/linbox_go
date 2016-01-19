package messages

type RequestResponseType uint16

const (
	INVALID_REQUEST_RESPONSE_TYPE RequestResponseType = iota

	// 认证信息
	AUTH_REQUEST_MSG
	AUTH_RESPONSE_MSG

	// 客户端同步未读消息数
	SYNC_UNREAD_REQUEST_MSG
	SYNC_UNREAD_RESPONSE_MSG

	// 客户端确认未读消息已被读取
	READ_ACK_REQUEST_MSG
	READ_ACK_RESPONSE_MSG

	// 以反序分页拉取消息
	PULL_OLD_MSG_REQUEST_MSG
	PULL_OLD_MSG_RESPONSE_MSG

	// 发送消息
	SEND_MSG_REQUEST_MSG
	SEND_MSG_RESPONSE_MSG

	// 新消息通知
	NEW_MSG_INFO

	// 用户离线通知
	OFFLINE_INFO

	// heart beat
	PING
	PONG
)

func (rrt RequestResponseType) Name() string {
	switch rrt {
	case INVALID_REQUEST_RESPONSE_TYPE:
		return "INVALID_REQUEST_RESPONSE_TYPE"
	case AUTH_REQUEST_MSG:
		return "AUTH_REQUEST_MSG"
	case AUTH_RESPONSE_MSG:
		return "AUTH_RESPONSE_MSG"
	case SEND_MSG_REQUEST_MSG:
		return "SEND_MSG_REQUEST_MSG"
	case SEND_MSG_RESPONSE_MSG:
		return "SEND_MSG_RESPONSE_MSG"
	case SYNC_UNREAD_REQUEST_MSG:
		return "SYNC_UNREAD_REQUEST_MSG"
	case SYNC_UNREAD_RESPONSE_MSG:
		return "SYNC_UNREAD_RESPONSE_MSG"
	case READ_ACK_REQUEST_MSG:
		return "READ_ACK_REQUEST_MSG"
	case READ_ACK_RESPONSE_MSG:
		return "READ_ACK_RESPONSE_MSG"
	case PULL_OLD_MSG_REQUEST_MSG:
		return "PULL_OLD_MSG_REQUEST_MSG"
	case PULL_OLD_MSG_RESPONSE_MSG:
		return "PULL_OLD_MSG_RESPONSE_MSG"
	case NEW_MSG_INFO:
		return "NEW_MSG_INFO"
	case OFFLINE_INFO:
		return OFFLINE_INFO
	case PING:
		return "PING"
	case PONG:
		return "PONG"
	default:
		return ""
	}
}

type MessageType int

const (
	MESSAGE_TYPE_ALL     MessageType = 1
	MESSAGE_TYPE_SESSION MessageType = 2
	MESSAGE_TYPE_GROUP   MessageType = 3
)
