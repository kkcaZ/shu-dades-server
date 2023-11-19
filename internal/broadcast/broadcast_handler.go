package broadcast

import (
	"encoding/json"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	routerUc "github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/kkcaz/shu-dades-server/pkg/models"
)

type BroadcastHandler struct {
	BroadcastUseCase domain.BroadcastUsecase
	Auth             domain.AuthUseCase
}

func NewBroadcastHandler(router *routerUc.RouterUseCase, uc domain.BroadcastUsecase, auth domain.AuthUseCase) {
	handler := BroadcastHandler{
		BroadcastUseCase: uc,
		Auth:             auth,
	}

	router.AddRoute("/broadcast", "POST", handler.Publish)
	router.AddRoute("/broadcast/subscribe", "POST", handler.Subscribe)
	router.AddRoute("/broadcast/user", "POST", handler.RegisterUser)
	router.AddRoute("/broadcast/user", "DELETE", handler.UnregisterUser)
}

func (b *BroadcastHandler) Publish(ctx *routerUc.RouterContext) {
	var request models.BroadcastRequest
	err := json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	err = b.BroadcastUseCase.Publish(request.Message, ctx.Sender)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.NewSuccessResponse(200, "Message published"))
}

func (b *BroadcastHandler) Subscribe(ctx *routerUc.RouterContext) {
	var request models.BroadcastSubscribeRequest
	err := json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	b.BroadcastUseCase.AddConnection(request.SubscribeAddress, request.PublishAddress)

	ctx.JSON(200, models.NewSuccessResponse(200, "Subscribed to broadcast"))
}

func (b *BroadcastHandler) RegisterUser(ctx *routerUc.RouterContext) {
	token := ctx.GetAuthToken()
	if token == nil {
		ctx.JSON(401, models.NewErrorResponse(401, "Unauthorized"))
		return
	}

	userClaim, err := b.Auth.GetUser(*token)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	b.BroadcastUseCase.RegisterUser(ctx.Sender, userClaim.UserId)
	ctx.JSON(200, models.NewSuccessResponse(200, "Registered user"))
}

func (b *BroadcastHandler) UnregisterUser(ctx *routerUc.RouterContext) {
	b.BroadcastUseCase.RemoveUser(ctx.Sender)
	ctx.JSON(200, models.NewSuccessResponse(200, "Removed user"))
}
