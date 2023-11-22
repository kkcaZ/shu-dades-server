package domain

type EncryptionUseCase interface {
	Encrypt(data string) (string, error)
	Decrypt(data []byte) ([]byte, error)
}
