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

<<<<<<< HEAD
func GenerateGenericResponseSucc(result *any.Any) (resp *GenericResponse) {
	resp = GenerateGenericResponse(SUCCESS)
=======
func (resp *GenericResponse) GenerateGenericResponseSucc(result *any.Any) {
	resp.GenerateGenericResponse(SUCCESS)
>>>>>>> c031e13b04ab4db207aacb7d2dea9b9c576d4ff3
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
