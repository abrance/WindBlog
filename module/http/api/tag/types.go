package tag

type AddRequest struct {
	Name  string `json:"name" binding:"required"` // tag name
	IsDir bool   `json:"isDir"`                   // 是否是目录
	Nice  uint   `json:"nice" binding:"required"`
}

type UpdateRequest struct {
	Name  string `binding:"required"` // tag name
	IsDir bool   `binding:"required"` // 是否是目录
	Nice  uint   `binding:"required"`
}
