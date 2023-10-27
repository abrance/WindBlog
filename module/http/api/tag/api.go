package tag

import "github.com/gin-gonic/gin"

func SetTagRouterGroup(router *gin.RouterGroup) {
	tagRouterGroup := router.Group("/tag")

	// v1 tag api
	v1TagRouterGroup := tagRouterGroup.Group("/v1")
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

		// tag 列表获取
		// tag 详情获取
		v1TagRouterGroup.GET("/get/:id")
		v1TagRouterGroup.GET("/list")
		v1TagRouterGroup.GET("/url")

		v1TagRouterGroup.POST("/add")
		v1TagRouterGroup.POST("/upload")

		v1TagRouterGroup.PUT("/update/meta/:id")
		v1TagRouterGroup.PATCH("/update/content/:id")
		v1TagRouterGroup.DELETE("/remove/:id")
		v1TagRouterGroup.DELETE("/remove_url/")
	}
}
