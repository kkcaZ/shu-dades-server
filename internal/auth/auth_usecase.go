package auth

import (
	"encoding/json"
	"fmt"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"math/rand"
	"os"
)

type userData struct {
	Users []models.User `json:"users"`
}

type authUseCase struct {
	users       []models.User
	validTokens []string
}

func NewAuthUseCase() domain.AuthUseCase {
	users, err := readUsers()
	if err != nil {
		panic(err)
	}

	return &authUseCase{
		users: users,
	}
}

func readUsers() ([]models.User, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dat, err := os.ReadFile(fmt.Sprintf("%s/internal/data/auth/users.json", currentDir))
	if err != nil {
		return nil, err
	}

	var userData userData
	err = json.Unmarshal(dat, &userData)
	if err != nil {
		return nil, err
	}

	return userData.Users, nil
}

func (a *authUseCase) Authenticate(username string, password string) (*models.UserClaim, error) {
	for _, user := range a.users {
		if user.Username == username && user.Password == password {
			token := generateToken()
			a.validTokens = append(a.validTokens, token)
			return &models.UserClaim{
				Token: token,
				Role:  user.Role,
			}, nil
		}
	}

	return nil, nil
}

func (a *authUseCase) TokenIsValid(token string) bool {
	for _, validToken := range a.validTokens {
		if validToken == token {
			return true
		}
	}

	return false
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateToken() string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
