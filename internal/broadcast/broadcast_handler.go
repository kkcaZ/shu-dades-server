package broadcast

import (
	"encoding/json"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	routerUc "github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/kkcaz/shu-dades-server/pkg/models"
)

type BroadcastHandler struct {
	BroadcastUseCase domain.BroadcastUsecase
}

func NewBroadcastHandler(router *routerUc.RouterUseCase, uc domain.BroadcastUsecase) {
	handler := BroadcastHandler{
		BroadcastUseCase: uc,
	}

	router.AddRoute("/broadcast", "POST", handler.Publish)
	router.AddRoute("/broadcast/subscribe", "POST", handler.Subscribe)
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
