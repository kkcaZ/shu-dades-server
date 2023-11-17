package models

type Notification struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	UserId  string `json:"userId"`
}

type NotificationListResponse struct {
	StatusCode    int            `json:"statusCode"`
	Notifications []Notification `json:"notifications"`
}

type DeleteNotificationRequest struct {
	NotificationId string `json:"notificationId"`
}
