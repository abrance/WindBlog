package doc

import (
	"github.com/WindBlog/util/errors"
	"github.com/WindBlog/util/storage/file"
	"github.com/WindBlog/util/storage/json_storage"
	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"strings"
	"time"
)

var ()

func GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	f, err := json_storage.GetFileTable().Get(id)
	if err != nil {
		logger.Error(err)
	}
	//ctx.File()
	ctx.JSON(200, *f)
}

func ListHandler(ctx *gin.Context) {
	f, err := json_storage.GetFileTable().List(nil)
	if err != nil {
		logger.Error(err)
	}
	//ctx.File()
	ctx.JSON(200, f)
}

func UrlHandler(ctx *gin.Context) {
	url := ctx.Param("url")
	if strings.HasPrefix(url, file.FileUrlPrefix) {
		filePath := strings.Replace(url, file.FileUrlPrefix, "", 1)
		realPath := file.GetRealPath(filePath)
		fileutil.Exist(realPath)
		ctx.File(realPath)
	}
}

type AddFileRequest struct {
	Name      string `json:"name" binding:"required"`
	Url       string `json:"url"`        // 地址, file://  表示本地
	IsArchive bool   `json:"is_archive"` // 是否已归档
	ArchiveId string `json:"archive_id"` //归档id
}

// 	Id         string // unique key, 数字整型
//	Name       string // 书名
//	Url        string // 地址, file://  表示本地
//	IsArchive  bool   // 是否已归档
//	ArchiveId  string //归档id
//	CreateTime timestamp.Timestamp
//	UpdateTime timestamp.Timestamp

func AddHandler(ctx *gin.Context) {
	req := AddFileRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(errors.ValidationException)
		return
	}
	f := &json_storage.File{
		Id:         "1",
		Name:       req.Name,
		Url:        req.Url,
		IsArchive:  req.IsArchive,
		ArchiveId:  "",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	err = json_storage.GetFileTable().Insert(f)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("req: %v", req)
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
