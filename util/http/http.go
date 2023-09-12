package http

import (
	"github.com/WindBlog/util/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response
// Json 响应基类
type Response struct {
	Data any
}

type ResponseData struct {
	Msg  string      `json:"message"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func Responses(ctx *gin.Context, code int, data interface{}) {
	// 优化此段代码
	if data == nil {
		data = ""
	}

	resp := ResponseData{
		Code: code,
		Data: data,
	}

	resp.Msg = errors.Msg[code]

	ctx.JSON(http.StatusOK, resp)
}
