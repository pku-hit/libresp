package proto

var (
	SUCCESS       = &Response{Code: "RC00000", Info: "成功"}
	GENERAL_ERROR = &Response{Code: "RC60000", Info: "服务错误"}
	ERROR         = &Response{Code: "RC99999", Info: "失败"}
)
