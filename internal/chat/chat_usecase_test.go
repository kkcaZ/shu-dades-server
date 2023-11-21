package chat

import (
	"github.com/kkcaz/shu-dades-server/internal/domain/mocks"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

func TestChatUseCase_GetChatThumbnails(t *testing.T) {
	thumbnails := []models.ChatThumbnail{
		{
			ChatId: "test",
			Participants: []models.Participant{
				{
					UserId: "test",
				},
			},
		},
	}

	testCases := []struct {
		name           string
		userId         string
		repoThumbnails []models.ChatThumbnail
		expected       []models.ChatThumbnail
		err            error
	}{
		{
			name:           "Happy path",
			userId:         "test",
			repoThumbnails: thumbnails,
			expected:       thumbnails,
		},
		{
			name:           "Happy path",
			userId:         "john",
			repoThumbnails: thumbnails,
			expected:       []models.ChatThumbnail{},
		},
		{
			name:           "Sad path",
			userId:         "test",
			repoThumbnails: thumbnails,
			expected:       nil,
			err:            assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := mocks.NewChatRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			logger := slog.Default()
			testUc := NewChatUseCase(repo, auth, broadcast, *logger)

			repo.On("GetAllChatThumbnails").Return(tc.repoThumbnails, tc.err)

			result, err := testUc.GetChatThumbnails(tc.userId)
			assert.Equal(t, tc.expected, result)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestChatUseCase_GetChat(t *testing.T) {
	testCases := []struct {
		name     string
		chatId   string
		expected *models.Chat
		err      error
	}{
		{
			name:   "Happy path",
			chatId: "test",
			expected: &models.Chat{
				Id: "test",
			},
		},
		{
			name:     "Sad path",
			chatId:   "test",
			expected: nil,
			err:      assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := mocks.NewChatRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			logger := slog.Default()
			testUc := NewChatUseCase(repo, auth, broadcast, *logger)

			repo.On("GetChat", tc.chatId).Return(tc.expected, tc.err)

			result, err := testUc.GetChat(tc.chatId)
			assert.Equal(t, tc.expected, result)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestChatUseCase_GetChatParticipantIds(t *testing.T) {
	testCases := []struct {
		name     string
		chatId   string
		chat     *models.Chat
		expected []string
		err      error
	}{
		{
			name:   "Happy path",
			chatId: "test",
			chat: &models.Chat{
				Id: "test",
				Participants: []models.Participant{
					{
						UserId: "test",
					},
				},
			},
			expected: []string{"test"},
		},
		{
			name:     "Sad path",
			chatId:   "test",
			chat:     nil,
			expected: nil,
			err:      assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := mocks.NewChatRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			logger := slog.Default()
			testUc := NewChatUseCase(repo, auth, broadcast, *logger)

			repo.On("GetChat", tc.chatId).Return(tc.chat, tc.err)

			result, err := testUc.GetChatParticipantIds(tc.chatId)
			assert.Equal(t, tc.expected, result)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestChatUseCase_CreateChat(t *testing.T) {
	testCases := []struct {
		name       string
		userIds    []string
		user       *models.User
		authErr    error
		expected   *models.Chat
		createErr  error
		publishErr error
	}{
		{
			name:    "Happy path",
			userIds: []string{"test"},
			user: &models.User{
				Id:       "test",
				Username: "john",
			},
			authErr: nil,
			expected: &models.Chat{
				Id: "test",
				Participants: []models.Participant{
					{
						UserId:   "test",
						Username: "john",
					},
				},
			},
		},
		{
			name:       "Sad path - auth error",
			userIds:    []string{"test"},
			user:       nil,
			authErr:    assert.AnError,
			createErr:  nil,
			publishErr: nil,
			expected:   nil,
		},
		{
			name:    "Sad path - creation error",
			userIds: []string{"test"},
			user: &models.User{
				Id:       "test",
				Username: "john",
			},
			authErr:    nil,
			createErr:  assert.AnError,
			publishErr: nil,
			expected:   nil,
		},
		{
			name:    "Sad path - publish error",
			userIds: []string{"test"},
			user: &models.User{
				Id:       "test",
				Username: "john",
			},
			authErr:    nil,
			createErr:  nil,
			publishErr: assert.AnError,
			expected:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := mocks.NewChatRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			logger := slog.Default()
			testUc := NewChatUseCase(repo, auth, broadcast, *logger)

			auth.On("GetUserById", mock.AnythingOfType("string")).Return(tc.user, tc.authErr)
			if tc.authErr == nil {
				repo.On("CreateChat", mock.AnythingOfType("models.Chat")).Return(tc.createErr)
				if tc.createErr == nil {
					broadcast.On("PublishToUsers", mock.AnythingOfType("string"), "newChat", tc.userIds).Return(tc.publishErr)
				}
			}

			result, err := testUc.CreateChat(tc.userIds)
			if tc.createErr == nil && tc.authErr == nil && tc.publishErr == nil {
				assert.Equal(t, tc.expected.Participants, result.Participants)
				assert.Equal(t, 0, len(result.Messages))
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestChatUseCase_SendMessage(t *testing.T) {
	testCases := []struct {
		name      string
		chatId    string
		content   string
		userId    string
		addErr    error
		notifyErr error
	}{
		{
			name:      "Happy path",
			chatId:    "test",
			content:   "test",
			userId:    "test",
			addErr:    nil,
			notifyErr: nil,
		},
		{
			name:      "Sad path",
			chatId:    "test",
			content:   "test",
			userId:    "test",
			addErr:    assert.AnError,
			notifyErr: nil,
		},
		{
			name:      "Sad path",
			chatId:    "test",
			content:   "test",
			userId:    "test",
			addErr:    nil,
			notifyErr: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := mocks.NewChatRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			logger := slog.Default()
			testUc := NewChatUseCase(repo, auth, broadcast, *logger)

			auth.On("GetUserById", mock.AnythingOfType("string")).Return(&models.User{
				Id: tc.userId,
			}, nil)
			repo.On("AddMessage", tc.chatId, mock.AnythingOfType("models.Message")).Return(tc.addErr)
			if tc.addErr == nil {
				repo.On("GetChat", tc.chatId).Return(&models.Chat{
					Id: tc.chatId,
					Participants: []models.Participant{
						{
							UserId: tc.userId,
						},
					},
				}, nil)
				broadcast.On("PublishToUsers", mock.AnythingOfType("string"), "message", mock.AnythingOfType("[]string")).Return(tc.notifyErr)
			}

			err := testUc.SendMessage(tc.chatId, tc.content, tc.userId)
			assert.Equal(t, tc.addErr, err)
			if tc.notifyErr != nil || tc.addErr != nil {
				assert.Error(t, err)
			}
		})
	}
}
