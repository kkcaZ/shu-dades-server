package models

type BroadcastRequest struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type BroadcastSubscribeRequest struct {
	PublishAddress   string `json:"publishAddress"`
	SubscribeAddress string `json:"subscribeAddress"`
}

type BroadcastConnection struct {
	PublishAddress   string `json:"publishAddress"`
	SubscribeAddress string `json:"subscribeAddress"`
	UserId           string `json:"userId"`
}
