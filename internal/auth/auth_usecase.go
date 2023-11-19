package auth

import (
	"encoding/json"
	"fmt"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/pkg/errors"
	"math/rand"
	"os"
)

type userData struct {
	Users []models.User `json:"users"`
}

type authUseCase struct {
	users       []models.User
	validTokens map[string]*models.User
}

func NewAuthUseCase() domain.AuthUseCase {
	users, err := readUsers()
	if err != nil {
		panic(err)
	}

	return &authUseCase{
		users:       users,
		validTokens: make(map[string]*models.User),
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
			a.validTokens[token] = &user

			return &models.UserClaim{
				UserId: user.Id,
				Token:  token,
				Role:   user.Role,
			}, nil
		}
	}

	return nil, nil
}

func (a *authUseCase) TokenIsValid(token string) bool {
	return a.validTokens[token] != nil
}

func (a *authUseCase) GetUser(token string) (*models.UserClaim, error) {
	user := a.validTokens[token]
	if user == nil {
		return nil, errors.New("invalid token")
	}

	return &models.UserClaim{
		UserId: user.Id,
		Token:  token,
		Role:   user.Role,
	}, nil
}

func (a *authUseCase) GetAllUserIds() []string {
	ids := make([]string, 0)
	for _, user := range a.users {
		ids = append(ids, user.Id)
	}
	return ids
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateToken() string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (a *authUseCase) GetUserById(userId string) (*models.User, error) {
	for _, user := range a.users {
		if user.Id == userId {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (a *authUseCase) GetAllUsersInfo() ([]models.UserInfo, error) {
	userInfos := make([]models.UserInfo, 0)
	for _, user := range a.users {
		userInfo := models.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		}
		userInfos = append(userInfos, userInfo)
	}
	return userInfos, nil
}
