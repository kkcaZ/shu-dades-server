package notification

import (
	"encoding/json"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/kkcaz/shu-dades-server/pkg/models"
)

type NotificationHandler struct {
	UseCase domain.NotificationUseCase
	Auth    domain.AuthUseCase
}

func NewNotificationHandler(router *router.RouterUseCase, uc domain.NotificationUseCase, auc domain.AuthUseCase) {
	handler := NotificationHandler{
		UseCase: uc,
		Auth:    auc,
	}

	router.AddRoute("/notification", models.GET, handler.Get)
	router.AddRoute("/notification", models.DELETE, handler.Delete)
}

func (n NotificationHandler) Get(ctx *router.RouterContext) {
	token := ctx.GetAuthToken()
	if token == nil {
		ctx.JSON(401, models.NewErrorResponse(401, "Unauthorized"))
		return
	}

	userClaim, err := n.Auth.GetUser(*token)

	notifications, err := n.UseCase.Get(userClaim.UserId)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.NotificationListResponse{
		StatusCode:    200,
		Notifications: notifications,
	})
}

func (n NotificationHandler) Delete(ctx *router.RouterContext) {
	var request models.RequestById
	err := json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	token := ctx.GetAuthToken()
	if token == nil {
		ctx.JSON(401, models.NewErrorResponse(401, "Unauthorized"))
		return
	}

	userClaim, err := n.Auth.GetUser(*token)

	err = n.UseCase.Delete(userClaim.UserId, request.Id)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.NewSuccessResponse(200, "Notification deleted"))
}
