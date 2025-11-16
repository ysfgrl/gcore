package gmodel

import "github.com/ysfgrl/gcore/gerror"

type ResApi[CType any] struct {
	Code    int           `json:"code"`
	Content CType         `json:"content"`
	Error   *gerror.Error `json:"error"`
}
