package chat

import (
	"encoding/json"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/kkcaz/shu-dades-server/pkg/models"
)

type ChatHandler struct {
	ChatUseCase domain.ChatUseCase
	Auth        domain.AuthUseCase
}

func NewChatHandler(router *router.RouterUseCase, chatUseCase domain.ChatUseCase, auth domain.AuthUseCase) {
	handler := ChatHandler{
		ChatUseCase: chatUseCase,
		Auth:        auth,
	}

	router.AddRoute("/chat/thumbnails", models.GET, handler.GetChatThumbnails)
	router.AddRoute("/chat", models.GET, handler.GetChat)
	router.AddRoute("/chat", models.POST, handler.CreateChat)
	router.AddRoute("/chat/message", models.POST, handler.SendMessage)
}

func (c ChatHandler) GetChatThumbnails(ctx *router.RouterContext) {
	token := ctx.GetAuthToken()
	if token == nil {
		ctx.JSON(401, models.NewErrorResponse(401, "Unauthorized"))
		return
	}

	userClaim, err := c.Auth.GetUser(*token)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	thumbnails, err := c.ChatUseCase.GetChatThumbnails(userClaim.UserId)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.ChatThumbnailsResponse{
		StatusCode: 200,
		Chats:      thumbnails,
	})
}

func (c ChatHandler) GetChat(ctx *router.RouterContext) {
	token := ctx.GetAuthToken()
	if token == nil {
		ctx.JSON(401, models.NewErrorResponse(401, "Unauthorized"))
		return
	}

	userClaim, err := c.Auth.GetUser(*token)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	var request models.RequestById
	err = json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	chat, err := c.ChatUseCase.GetChat(request.Id)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	if chat == nil {
		ctx.JSON(404, models.NewErrorResponse(404, "Chat not found"))
		return
	}

	for _, participant := range chat.Participants {
		if participant.UserId == userClaim.UserId {
			ctx.JSON(200, models.ChatResponse{
				StatusCode: 200,
				Chat:       chat,
			})
			return
		}
	}

	ctx.JSON(403, models.NewErrorResponse(403, "Forbidden"))
}

func (c ChatHandler) CreateChat(ctx *router.RouterContext) {
	var request models.CreateChatRequest
	err := json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	chat, err := c.ChatUseCase.CreateChat(request.UserIds)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.ChatResponse{
		StatusCode: 200,
		Chat:       chat,
	})
}

func (c ChatHandler) SendMessage(ctx *router.RouterContext) {
	var request models.SendMessageRequest
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

	userClaim, err := c.Auth.GetUser(*token)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	err = c.ChatUseCase.SendMessage(request.ChatId, request.Message, userClaim.UserId)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.NewSuccessResponse(200, "Message sent"))
}
