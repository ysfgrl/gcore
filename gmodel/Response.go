package gmodel

import "github.com/ysfgrl/gcore/gerror"

type Response struct {
	Code    int           `json:"code"`
	Content any           `json:"content"`
	Error   *gerror.Error `json:"error"`
}
type ResponseOk struct {
	Code    int           `json:"code"`
	Content Ok            `json:"content"`
	Error   *gerror.Error `json:"error"`
}
type ResponseStr struct {
	Code    int           `json:"code"`
	Content Str           `json:"content"`
	Error   *gerror.Error `json:"error"`
}

type Ok struct {
	IsOk bool `json:"isOk"`
}
type Str struct {
	Value string `json:"value"`
}
