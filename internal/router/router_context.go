package router

import "encoding/json"

type RouterContext struct {
	Body     string
	Headers  string
	Response *string
}

func (rc *RouterContext) JSON(code int, i interface{}) {
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	respStr := string(bytes)
	rc.Response = &respStr
}
