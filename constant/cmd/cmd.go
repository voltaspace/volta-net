package cmd

import "time"

const (

	STRIKE = "volta-http/wspl"
	DEVIATION_LEN = len(STRIKE) + 1

	GATEWAY_PADDING string = "#volta-net#" //.ws通信数据包填充字符

	SEND_TO_UID      int = 14
	SEND_TO_UID_LIST int = 30
	SEND_TO_ALL      int = 1
	ON_CLOSE         int = 4
	IS_ONLINE        int = 11
	BIND_UID         int = 12
	UNBIND_UID       int = 13
	JOIN_GROUP       int = 20
	SEND_TO_GROUP    int = 22
	LEAVE_GROUP      int = 23
	PING             int = 201

	PING_INTERVAL  int64         = 20 //.TCP心跳超时时间
	W_CHAN_TIMEOUT time.Duration = 2  //.发送消息Gateway处理超时时间(s)

	SUCCESS int = 0   //.成功
	ERROR   int = 200 //.失败

	SESSION_HEAD    string = "session_"

	DEFAULT_FUNC	string = "Handle"
)