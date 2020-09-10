package libresp

import "github.com/pku-hit/libresp/proto"

var (
	SUCCESS       = &proto.Response{Code: "RC00000", Info: "成功"}
	GENERAL_ERROR = &proto.Response{Code: "RC60000", Info: "服务错误"}
	ERROR         = &proto.Response{Code: "RC99999", Info: "失败"}
)
