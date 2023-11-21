package notification

import (
	"github.com/kkcaz/shu-dades-server/internal/domain/mocks"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

func TestNotificationUseCase_Get(t *testing.T) {
	testCases := []struct {
		name     string
		userId   string
		expected []models.Notification
		err      error
	}{
		{
			name:   "Happy path",
			userId: "test",
			expected: []models.Notification{
				{
					Id: "notif",
				},
			},
			err: nil,
		},
		{
			name:     "Sad path",
			userId:   "test",
			expected: nil,
			err:      assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewNotificationRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			testUc := NewNotificationUseCase(repo, auth, broadcast, *logger)

			repo.On("Get", tc.userId).Return(tc.expected, tc.err)

			notification, err := testUc.Get(tc.userId)
			assert.Equal(t, tc.expected, notification)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestNotificationUseCase_Add(t *testing.T) {
	testCases := []struct {
		name     string
		userId   string
		message  string
		expected error
	}{
		{
			name:     "Happy path",
			userId:   "test",
			message:  "test",
			expected: nil,
		},
		{
			name:     "Sad path",
			userId:   "test",
			message:  "test",
			expected: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewNotificationRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			testUc := NewNotificationUseCase(repo, auth, broadcast, *logger)

			repo.On("Add", mock.AnythingOfType("models.Notification")).Return(tc.expected)

			err := testUc.Add(tc.userId, tc.message)
			assert.Equal(t, tc.expected, err)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestNotificationUseCase_AddAll(t *testing.T) {
	testCases := []struct {
		name     string
		message  string
		expected error
	}{
		{
			name:     "Happy path",
			message:  "test",
			expected: nil,
		},
		{
			name:     "Sad path",
			message:  "test",
			expected: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewNotificationRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			testUc := NewNotificationUseCase(repo, auth, broadcast, *logger)

			auth.On("GetAllUserIds").Return([]string{"test"}, nil)
			repo.On("Add", mock.AnythingOfType("models.Notification")).Return(tc.expected)
			broadcast.On("PublishToUsers", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(tc.expected)

			err := testUc.AddAll(tc.message)
			assert.Equal(t, tc.expected, err)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestNotificationUseCase_Delete(t *testing.T) {
	testCases := []struct {
		name           string
		userId         string
		notificationId string
		expected       error
	}{
		{
			name:           "Happy path",
			userId:         "test",
			notificationId: "test",
			expected:       nil,
		},
		{
			name:           "Sad path",
			userId:         "test",
			notificationId: "test",
			expected:       assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewNotificationRepository(t)
			auth := mocks.NewAuthUseCase(t)
			broadcast := mocks.NewBroadcastUseCase(t)
			testUc := NewNotificationUseCase(repo, auth, broadcast, *logger)

			repo.On("Delete", tc.userId, tc.notificationId).Return(tc.expected)

			err := testUc.Delete(tc.userId, tc.notificationId)
			assert.Equal(t, tc.expected, err)
		})
	}
}
