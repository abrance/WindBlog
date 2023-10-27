package doc

import (
	"github.com/WindBlog/util/errors"
	"github.com/WindBlog/util/http"
	"github.com/WindBlog/util/storage/file"
	"github.com/WindBlog/util/storage/json_storage"
	"github.com/WindBlog/util/storage/sqlite"
	"github.com/WindBlog/util/storage/sqlite/table"
	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetHandler(ctx *gin.Context) {
	// curl 组织该 get 请求
	// curl http://localhost:5000/api/doc/v1/get/1 -v
	id := ctx.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err)
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	_file := table.File{}
	f := sqlite.GetDB().First(&_file, uint(_id))

	if f.Error != nil {
		logger.Error(f.Error)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	//ctx.File()
	http.Responses(ctx, errors.OK, _file)
}

func ListHandler(ctx *gin.Context) {

	// curl 组织该 get 请求
	// curl http://localhost:5000/api/doc/v1/list -v
	var fileLs []table.File
	db := sqlite.GetDB().Find(&fileLs)
	if db.Error != nil {
		logger.Error(db.Error)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}

	//f, err := json_storage.GetFileTable().List(nil)
	//if err != nil {
	//	logger.Error(err)
	//}
	http.Responses(ctx, errors.OK, fileLs)
}

func UrlHandler(ctx *gin.Context) {
	url := ctx.Query("url")
	// 根据不同的前缀去不同的地方找文件, file: 表示本地 ,后面的是文件相对路径
	if strings.HasPrefix(url, file.FileUrlPrefix) {
		filePath := strings.Replace(url, file.FileUrlPrefix, "", 1)
		realPath := file.GetRealPath(filePath)
		fileutil.Exist(realPath)
		ctx.File(realPath)
		http.Responses(ctx, errors.OK, nil)
	} else {
		http.Responses(ctx, errors.FileNotExistError, nil)
	}
}

// AddHandler
// curl -X POST "http://localhost:5000/api/doc/v1/add" -d '{"name":"test", "url":"file://test.txt", "is_archive":false, "archive_id":""}'
// 新增文件上传
func AddHandler(ctx *gin.Context) {
	//req := AddFileRequest{}
	//err := ctx.ShouldBindJSON(&req)
	//if err != nil {
	//	logger.Error(errors.ValidationException)
	//	return
	//}
	//f := &json_storage.File{
	//	Id:         "1",
	//	Name:       req.Name,
	//	Url:        req.Url,
	//	IsArchive:  req.IsArchive,
	//	ArchiveId:  "",
	//	CreateTime: time.Now().Unix(),
	//	UpdateTime: time.Now().Unix(),
	//}
	//err = json_storage.GetFileTable().Insert(f)

	req := AddFileRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(errors.ValidationException)
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	f := &table.File{
		Name:       req.Name,
		Url:        req.Url,
		IsArchive:  req.IsArchive,
		ArchiveId:  req.ArchiveId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	db := sqlite.GetDB().Create(f)
	if db.Error != nil {
		logger.Error(db.Error)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	logger.Info("req: %v", req)
	http.Responses(ctx, errors.OK, nil)
}

// UploadHandler
// curl -X POST "http://localhost:5000/api/doc/v1/upload" -F "doc=@test"
// 上传文件
func UploadHandler(ctx *gin.Context) {
	_file, err := ctx.FormFile("doc")
	if err != nil {
		logger.Error(errors.ValidationException, err)
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	// 限制文件大小
	if _file.Size > 1024*1024*10 {
		logger.Error(errors.ValidationException)
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	// 将文件保存到本地
	err = ctx.SaveUploadedFile(_file, file.GetRealPath(_file.Filename))
	if err != nil {
		logger.Error(err)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	http.Responses(ctx, errors.OK, nil)
}

// UpdateMetaHandler
// curl -X PUT "http://localhost:5000/api/doc/v1/update/meta/1" -d '{"name":"test", "url":"file://test.txt", "is_archive":false, "archive_id":""}'
// 更新文件元信息
func UpdateMetaHandler(ctx *gin.Context) {
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
	//oldFile, err := json_storage.GetFileTable().Get(id)
	//if err != nil {
	//	logger.Error(errors.FileNotExistError)
	//	return
	//}
	//f := &json_storage.File{
	//	Id:         id,
	//	Name:       req.Name,
	//	Url:        req.Url,
	//	IsArchive:  req.IsArchive,
	//	ArchiveId:  req.ArchiveId,
	//	CreateTime: oldFile.CreateTime,
	//	UpdateTime: req.UpdateTime,
	//}
	//err = json_storage.GetFileTable().Update(id, f)
	f := &table.File{}
	db := sqlite.GetDB().First(f, id)
	if db.Error != nil {
		logger.Error(db.Error)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}

	if req.Name != "" {
		f.Name = req.Name
	}
	if req.Url != "" {
		f.Url = req.Url
	}
	if req.IsArchive != false {
		f.IsArchive = req.IsArchive
	}
	if req.ArchiveId != "" {
		f.ArchiveId = req.ArchiveId
	}

	db = sqlite.GetDB().Save(f)
	if db.Error != nil {
		logger.Error(db.Error)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	http.Responses(ctx, errors.OK, nil)
}

// UpdateContentHandler
// curl -X PATCH "http://localhost:5000/api/doc/v1/update/content/1" -d '{"name":"test"}'
// 更新文件内容
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
