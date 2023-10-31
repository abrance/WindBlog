package tag

import (
	"github.com/WindBlog/util/errors"
	"github.com/WindBlog/util/http"
	"github.com/WindBlog/util/storage/sqlite"
	"github.com/WindBlog/util/storage/sqlite/table"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"strconv"
	"time"
)

// GetHandler 获取单个 tag
// curl http://localhost:5000/api/tag/v1/get/1 -v
func GetHandler(ctx *gin.Context) {
	_id := ctx.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	tag := &table.Tag{}
	db := sqlite.GetDB().First(tag, id)
	if db.Error != nil {
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	http.Responses(ctx, errors.OK, tag)
}

// ListHandler 获取 tag 列表
// curl http://localhost:5000/api/tag/v1/list -v
func ListHandler(ctx *gin.Context) {
	var tags []table.Tag
	db := sqlite.GetDB().Find(&tags)
	if db.Error != nil {
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	http.Responses(ctx, errors.OK, tags)
}

// AddHandler 添加 tag
// curl -X POST http://localhost:5000/api/tag/v1/add -d '{"name":"test_tag", "isDir": false, "nice": 1}'   -v
func AddHandler(ctx *gin.Context) {
	a := AddRequest{}
	err := ctx.ShouldBindJSON(&a)
	if err != nil {
		logger.Error(err)
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	tag := &table.Tag{
		Name:       a.Name,
		IsDir:      a.IsDir,
		Nice:       a.Nice,
		CreateTime: time.Now().Unix(),
	}
	db := sqlite.GetDB().Create(tag)
	if db.Error != nil {
		logger.Error(db.Error)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	http.Responses(ctx, errors.OK, tag)
}

// UpdateHandler 更新 tag
// curl -X PUT http://localhost:5000/api/tag/v1/update/1 -d '{"name":"test_tag_1", "is_dir": false, "nice": 1}'   -v
func UpdateHandler(ctx *gin.Context) {
	_id := ctx.Param("id")
	id, err := strconv.Atoi(_id)
	r := UpdateRequest{}
	err = ctx.ShouldBindJSON(&r)
	if err != nil {
		logger.Error(err)
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	tag := &table.Tag{}
	db := sqlite.GetDB().First(tag, id)
	if db.Error != nil {
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	if r.Name != "" {
		tag.Name = r.Name
	}
	// is dir 必传
	tag.IsDir = r.IsDir
	tag.Nice = r.Nice
	db = sqlite.GetDB().Save(tag)
	if db.Error != nil {
		logger.Error(db.Error)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	http.Responses(ctx, errors.OK, tag)
}

// RemoveHandler 删除 tag
// curl -X DELETE http://localhost:5000/api/tag/v1/remove/1 -v
func RemoveHandler(ctx *gin.Context) {
	_id := ctx.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	tag := &table.Tag{}
	db := sqlite.GetDB().Delete(tag, id)
	if db.Error != nil {
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	http.Responses(ctx, errors.OK, tag)
}
