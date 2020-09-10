package libresp

import "github.com/golang/protobuf/ptypes/any"

func (resp *MapResponse) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

<<<<<<< HEAD
func GenerateMapResponseSucc(result map[string]*any.Any) (resp *MapResponse) {
	resp = GenerateMapResponse(SUCCESS)
=======
func (resp *MapResponse) GenerateMapResponseSucc(result map[string]*any.Any) {
	resp.GenerateMapResponse(SUCCESS)
>>>>>>> c031e13b04ab4db207aacb7d2dea9b9c576d4ff3
	resp.Result = result
	return
}

func (resp *MapResponse) GenerateMapResponse(code *Response) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = code.Info
	}
	return
}

func (resp *MapResponse) GenerateMapResponseWithInfo(code *Response, info string) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = info
	}
	return
}