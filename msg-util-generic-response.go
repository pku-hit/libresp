package libresp

import any "github.com/golang/protobuf/ptypes/any"

func (resp *GenericResponse) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

func GenerateGenericResponseSucc(result *any.Any) (resp *GenericResponse) {
	resp = GenerateGenericResponse(SUCCESS)
	resp.Result = result
	return
}

func GenerateGenericResponse(code *Response) (resp *GenericResponse) {
	resp = &GenericResponse{
		Code: code.Code,
		Info: code.Info,
	}
	return
}

func GenerateGenericResponseWithInfo(code *Response, info string) (resp *GenericResponse) {
	if code == nil {
		resp = nil
	} else {
		resp = &GenericResponse{
			Code: code.Code,
			Info: info,
		}
	}
	return
}
