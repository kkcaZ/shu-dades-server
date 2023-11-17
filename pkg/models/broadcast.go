package models

type BroadcastRequest struct {
	Message string `json:"message"`
}

type BroadcastSubscribeRequest struct {
	PublishAddress   string `json:"publishAddress"`
	SubscribeAddress string `json:"subscribeAddress"`
}

type BroadcastConnection struct {
	PublishAddress   string `json:"publishAddress"`
	SubscribeAddress string `json:"subscribeAddress"`
}

type BroadcastMessage struct {
	Message string `json:"message"`
	Sender  string `json:"sender"`
	Type    string `json:"type"`
}