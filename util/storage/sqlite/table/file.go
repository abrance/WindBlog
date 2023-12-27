package table

type File struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string // 书名
	Url        string // 地址, file://  表示本地
	IsArchive  bool   // 是否已归档
	ArchiveId  string //归档id
	CreateTime int64
	UpdateTime int64
}
