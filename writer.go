package response

import (
	"github.com/gin-gonic/gin"
)

type writer struct {
	Code    int         `json:"-"`
	Error   error       `json:"-"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *writer) Pending() *writer {
	res.Status = StatusPENDING
	return res
}

func (res *writer) Msg(msg string) *writer {
	res.Message = msg
	return res
}

// JSON godoc.
func (res *writer) JSON(ctx *gin.Context) {
	if res.Error != nil {
		_ = ctx.Error(res.Error)
	}

	ctx.JSON(res.Code, res)
	ctx.Abort()
}
