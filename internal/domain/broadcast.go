package domain

type BroadcastUsecase interface {
	Publish(message string, eventType string) error
	PublishToUsers(message string, eventType string, users []string) error
	AddConnection(subscribeAddress string, publishAddress string)
	RegisterUser(addr string, userId string)
	RemoveUser(addr string)
}
