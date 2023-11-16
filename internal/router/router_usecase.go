package router

import (
	"encoding/json"
	"fmt"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/pkg/errors"
	"log/slog"
)

type HandlerKey struct {
	Route  string
	Method models.RequestType
}

type RouterUseCase struct {
	Logger   slog.Logger
	Handlers map[HandlerKey]func(ctx *RouterContext)
}

func NewRouterUseCase(logger slog.Logger) *RouterUseCase {
	return &RouterUseCase{
		Logger:   logger,
		Handlers: make(map[HandlerKey]func(ctx *RouterContext)),
	}
}

func (r *RouterUseCase) Handle(buffer []byte, mLen int) (*string, error) {
	r.Logger.Info("received message: " + string(buffer[:mLen]))

	req, err := r.parseMessage(buffer[:mLen])
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse message: ")
	}

	handlerKey := HandlerKey{
		Route:  req.Route,
		Method: req.Type,
	}
	handler, ok := r.Handlers[handlerKey]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no handler found for key: %v", handlerKey))
	}

	reqBody, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}

	ctx := &RouterContext{
		Body:    string(reqBody),
		Headers: req.Headers,
	}

	handler(ctx)

	return ctx.Response, nil
}

func (r *RouterUseCase) parseMessage(message []byte) (*models.Request, error) {
	var request models.Request
	err := json.Unmarshal(message, &request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (r *RouterUseCase) AddRoute(route string, method models.RequestType, handler func(ctx *RouterContext)) {
	key := HandlerKey{
		Route:  route,
		Method: method,
	}
	r.Handlers[key] = handler
}
