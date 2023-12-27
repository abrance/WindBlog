package doc

type AddFileRequest struct {
	Name      string `json:"name" binding:"required"`
	Url       string `json:"url" binding:"required"`        // 地址, file://  表示本地
	IsArchive bool   `json:"is_archive" binding:"required"` // 是否已归档
	ArchiveId string `json:"archive_id" binding:"required"` //归档id
}

type UpdateFileMetaRequest struct {
	Name       string `json:"name"`
	Url        string `json:"url"`        // 地址, file://  表示本地
	IsArchive  bool   `json:"is_archive"` // 是否已归档
	ArchiveId  string `json:"archive_id"` //归档id
	UpdateTime int64  `json:"update_time"`
}

// 	Id         string // unique key, 数字整型
//	Name       string // 书名
//	Url        string // 地址, file://  表示本地
//	IsArchive  bool   // 是否已归档
//	ArchiveId  string //归档id
//	CreateTime timestamp.Timestamp
