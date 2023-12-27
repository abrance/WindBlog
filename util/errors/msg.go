package errors

var Msg = map[int]string{
	OK: "OK",
	// 按照 http 请求的处理流程的不同阶段归类异常
	SerializationException:  "序列化异常",
	HandleInternalException: "内部处理异常",
	IdValidationException:   "ID 验证异常",
	ValidationException:     "验证异常",
	HandleException:         "处理异常",

	MiddlewareException: "中间件异常",

	ApiTodoException:  "API 待完成异常",
	ApiAboveException: "API 异常",

	// 通用异常
	FileNotExistError: "文件不存在",
	FileDeleteError:   "文件删除异常",
	FileStorageError:  "文件存储异常",

	JsonInsertError: "JSON 插入异常",
	JsonSaveError:   "JSON 保存异常",
	JsonDBInitError: "JSON 数据库初始化异常",

	StorageError: "存储异常",
}
