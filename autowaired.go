package voltanet

var AutoWaired autowaired

type autowaired struct {
	WsEvents EventsInterface `volta:"wsEvents"`
}

