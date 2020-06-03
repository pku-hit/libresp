package libresp

import "github.com/golang/protobuf/ptypes/any"

func (resp *GenericResponse) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

func (resp *GenericResponse) GenerateGenericResponseSucc(result *any.Any) {
	resp.GenerateGenericResponse(SUCCESS)
	resp.Result = result
	return
}

func (resp *GenericResponse) GenerateGenericResponse(code *Response) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = code.Info
	}
	return
}

func (resp *GenericResponse) GenerateGenericResponseWithInfo(code *Response, info string) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = info
	}
	return
}
