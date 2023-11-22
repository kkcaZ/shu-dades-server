package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"log/slog"
)

type encryptionUseCase struct {
	logger slog.Logger
	key    string
}

func NewEncryptionUseCase(logger slog.Logger) domain.EncryptionUseCase {
	return &encryptionUseCase{
		logger: logger,
		key:    "12345678901234567890123456789012",
	}
}

func (e *encryptionUseCase) Encrypt(data string) (string, error) {
	text := []byte(data)
	key := []byte(e.key)

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := []byte("unique nonce")

	encryptedMessage := gcm.Seal(nil, nonce, text, nil)

	return string(encryptedMessage), nil
}

func (e *encryptionUseCase) Decrypt(data []byte) ([]byte, error) {
	e.logger.Info("Decrypting...", "data", data)

	ciphertext := data
	key := []byte(e.key)

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}

	nonce := []byte("unique nonce")

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, err
}
