package router

import (
	"encoding/json"
)

type RouterContext struct {
	Body     string
	Headers  map[string]string
	Response *string
	Sender   string
}

func (rc *RouterContext) JSON(code int, i interface{}) {
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	respStr := string(bytes)
	rc.Response = &respStr
}

func (rc *RouterContext) GetAuthToken() *string {
	authHeader := rc.Headers["Authorization"]
	if authHeader == "" {
		return nil
	}

	return &authHeader
}
