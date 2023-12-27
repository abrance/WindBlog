package api

import (
	"github.com/WindBlog/module/http/api/doc"
	"github.com/WindBlog/module/http/api/tag"
	"github.com/gin-gonic/gin"
)

func SetRouterGroup(router *gin.Engine) {
	ApiRouteGroup := router.Group("/api")
	doc.SetDocRouterGroup(ApiRouteGroup)
	tag.SetTagRouterGroup(ApiRouteGroup)
}
