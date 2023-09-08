package http

import (
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

func Responses(ctx *gin.Context, code int, msg string, data interface{}) {
	resp := ResponseData{
		Code: code,
		Data: data,
	}
	if msg != "" {
		resp.Msg = msg
	} else {
		resp.Msg = "error format"
	}

	ctx.JSON(http.StatusOK, resp)
}
