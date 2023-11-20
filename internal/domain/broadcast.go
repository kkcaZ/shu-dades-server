package domain

type BroadcastUseCase interface {
	PublishToUsers(message string, eventType string, users []string) error
	AddConnection(subscribeAddress string, publishAddress string)
	RegisterUser(addr string, userId string)
	RemoveUser(addr string)
}
