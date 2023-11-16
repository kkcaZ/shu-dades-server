package domain

type BroadcastUsecase interface {
	Publish(message string, sender string) error
	AddConnection(subscribeAddress string, publishAddress string)
}
