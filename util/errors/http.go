package errors

// 按照 http 请求的处理流程的不同阶段归类异常
const (
	SerializationException  = 1290
	HandleInternalException = 1220
	IdValidationException   = 1211
	ValidationException     = 1210
	HandleException         = 1200

	MiddlewareException = 1100

	ApiTodoException = 1002

	ApiAboveException = 1000
)
