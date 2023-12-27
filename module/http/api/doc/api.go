package doc

import "github.com/gin-gonic/gin"

func SetDocRouterGroup(router *gin.RouterGroup) {
	docRouterGroup := router.Group("/doc")

	// v1 doc api
	v1DocRouterGroup := docRouterGroup.Group("/v1")
	{
		// 单个文件获取元信息
		// 文件元信息列表
		// 单个文件获取内容
		// 新增文件元信息
		// 文件上传
		// 更新文件元信息
		// 更新文件内容
		// 删除文件元信息
		// 删除文件内容
		v1DocRouterGroup.GET("/get/:id", GetHandler)
		v1DocRouterGroup.GET("/list", ListHandler)
		v1DocRouterGroup.GET("/url", UrlHandler)

		v1DocRouterGroup.POST("/add", AddHandler)
		v1DocRouterGroup.POST("/upload", UploadHandler)

		v1DocRouterGroup.PUT("/update/meta/:id", UpdateMetaHandler)
		v1DocRouterGroup.PATCH("/update/content/:id", UpdateContentHandler)
		v1DocRouterGroup.DELETE("/remove/:id", RemoveHandler)
		v1DocRouterGroup.DELETE("/remove_url/", RemoveUrlHandler)
	}
}
