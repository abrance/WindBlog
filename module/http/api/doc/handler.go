package doc

import (
	"github.com/WindBlog/util/errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"strconv"
)

// Response
// Json 响应基类
type Response struct {
	Data any
}

func GetHandler(ctx *gin.Context) {
	_id := ctx.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		logger.Error(errors.IdValidationException)
	}

	//ctx.File()
	//ctx.JSON(200, "get")
}

func AddHandler(ctx *gin.Context) {
	ctx.JSON(200, "get")
}

func UpdateMetaHandler(ctx *gin.Context) {
	ctx.JSON(200, "get")
}

func UpdateContentHandler(ctx *gin.Context) {
	ctx.JSON(200, "get")
}

func RemoveHandler(ctx *gin.Context) {
	ctx.JSON(200, "get")
}
