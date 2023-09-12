package doc

import (
	"github.com/WindBlog/util/errors"
	"github.com/WindBlog/util/http"
	"github.com/WindBlog/util/storage/file"
	"github.com/WindBlog/util/storage/json_storage"
	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"os"
	"strings"
	"time"
)

func GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	f, err := json_storage.GetFileTable().Get(id)
	if err != nil {
		logger.Error(err)
	}
	//ctx.File()
	ctx.JSON(errors.OK, f)
}

func ListHandler(ctx *gin.Context) {
	f, err := json_storage.GetFileTable().List(nil)
	if err != nil {
		logger.Error(err)
	}
	http.Responses(ctx, errors.OK, f)
}

func UrlHandler(ctx *gin.Context) {
	url := ctx.Query("url")
	// 根据不同的前缀去不同的地方找文件, file: 表示本地 ,后面的是文件相对路径
	if strings.HasPrefix(url, file.FileUrlPrefix) {
		filePath := strings.Replace(url, file.FileUrlPrefix, "", 1)
		realPath := file.GetRealPath(filePath)
		fileutil.Exist(realPath)
		ctx.File(realPath)
	} else {
		http.Responses(ctx, errors.FileNotExistError, nil)
	}
}

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
	http.Responses(ctx, errors.OK, nil)
}

func UploadHandler(ctx *gin.Context) {
	// curl 组织该 post 请求
	// curl http://localhost:5000/api/doc/v1/list -v
	_file, err := ctx.FormFile("doc")
	if err != nil {
		logger.Error(errors.ValidationException, err)
		return
	}
	// 限制文件大小
	if _file.Size > 1024*1024*10 {
		logger.Error(errors.ValidationException)
		return
	}
	// 将文件保存到本地
	err = ctx.SaveUploadedFile(_file, file.GetRealPath(_file.Filename))
	if err != nil {
		logger.Error(err)
		return
	}
	http.Responses(ctx, errors.OK, nil)
}

func UpdateMetaHandler(ctx *gin.Context) {
	// curl -X PUT "http://localhost:5000/api/doc/v1/update_meta/?id=152227692589057" -d '{"name":"test"}'
	id := ctx.Param("id")
	req := &UpdateFileMetaRequest{
		UpdateTime: time.Now().Unix(),
	}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		logger.Error(errors.ValidationException)
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	oldFile, err := json_storage.GetFileTable().Get(id)
	if err != nil {
		logger.Error(errors.FileNotExistError)
		return
	}
	f := &json_storage.File{
		Id:         id,
		Name:       req.Name,
		Url:        req.Url,
		IsArchive:  req.IsArchive,
		ArchiveId:  req.ArchiveId,
		CreateTime: oldFile.CreateTime,
		UpdateTime: req.UpdateTime,
	}
	err = json_storage.GetFileTable().Update(id, f)
	http.Responses(ctx, errors.OK, nil)
}

func UpdateContentHandler(ctx *gin.Context) {
	// curl -X PATCH "http://localhost:5000/api/doc/v1/update/content/152227692589057" -d '{"name":"test"}'
	http.Responses(ctx, errors.ApiTodoException, "")
}

func RemoveHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := json_storage.GetFileTable().Delete(id)
	if err != nil {
		logger.Error(err)
		http.Responses(ctx, errors.FileDeleteError, nil)
		return
	}
	http.Responses(ctx, errors.OK, id)
}

func RemoveUrlHandler(ctx *gin.Context) {
	url := ctx.Query("url")
	// 组织一个 curl 请求这个接口
	// curl -X DELETE "http://localhost:8080/doc/v1/remove_url/?url=file:test.txt"

	// 根据不同的前缀去不同的地方找文件, file: 表示本地 ,后面的是文件相对路径
	if strings.HasPrefix(url, file.FileUrlPrefix) {
		logger.Info(url)
		filePath := strings.Replace(url, file.FileUrlPrefix, "", 1)
		realPath := file.GetRealPath(filePath)
		logger.Info(realPath)
		if fileutil.Exist(realPath) {
			// 删除文件
			err := os.Remove(realPath)
			if err != nil {
				logger.Error(err)
				return
			}
		} else {
			logger.Error(errors.FileNotExistError)
			return
		}
	}
	http.Responses(ctx, 200, nil)
}
