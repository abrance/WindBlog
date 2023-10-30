package table

type Tag struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string // tag name
	IsDir      bool   // 是否是目录
	Nice       uint
	CreateTime int64
}
