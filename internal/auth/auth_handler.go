package auth

import (
	"encoding/json"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/kkcaz/shu-dades-server/pkg/models"
)

type AuthHandler struct {
	AuthUseCase domain.AuthUseCase
}

func NewAuthHandler(router *router.RouterUseCase, uc domain.AuthUseCase) {
	handler := AuthHandler{
		AuthUseCase: uc,
	}

	router.AddRoute("/auth", handler.Authenticate)
}

func (a AuthHandler) Authenticate(ctx *router.RouterContext) {
	var authRequest models.AuthRequest
	err := json.Unmarshal([]byte(ctx.Body), &authRequest)
	if err != nil {
		ctx.JSON(500, models.NewErrorResponse(500, "Internal server error"))
		return
	}

	if authRequest.Username == "" || authRequest.Password == "" {
		ctx.JSON(400, models.NewErrorResponse(400, "Missing username or password"))
		return
	}

	token, err := a.AuthUseCase.Authenticate(authRequest.Username, authRequest.Password)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	if token == nil {
		ctx.JSON(401, models.NewErrorResponse(401, "Invalid username or password"))
		return
	}

	ctx.JSON(200, models.AuthResponse{
		StatusCode: 200,
		Token:      *token,
	})
}
