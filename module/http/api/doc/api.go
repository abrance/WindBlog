package doc

import "github.com/gin-gonic/gin"

func SetDocRouterGroup(router *gin.RouterGroup) {
	docRouterGroup := router.Group("/doc")

	// v1 doc api
	v1DocRouterGroup := docRouterGroup.Group("/v1")
	{
		v1DocRouterGroup.GET("/get/:id", GetHandler)
		v1DocRouterGroup.GET("/list", ListHandler)
		v1DocRouterGroup.GET("/url", UrlHandler)
		// 暂时不提供删除文档的接口

		v1DocRouterGroup.POST("/add", AddHandler)
		v1DocRouterGroup.PUT("/update/meta/:id", UpdateMetaHandler)
		v1DocRouterGroup.PATCH("/update/content/:id", UpdateContentHandler)
		v1DocRouterGroup.DELETE("/remove/:id", RemoveHandler)
	}
}
