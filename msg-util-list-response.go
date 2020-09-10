package libresp

import "github.com/golang/protobuf/ptypes/any"


func (resp *ListResponse) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

<<<<<<< HEAD
func GenerateListResponseSucc(result []*any.Any) (resp *ListResponse) {
	resp = GenerateListResponse(SUCCESS)
=======
func (resp *ListResponse) GenerateListResponseSucc(result []*any.Any) {
	resp.GenerateListResponse(SUCCESS)
>>>>>>> c031e13b04ab4db207aacb7d2dea9b9c576d4ff3
	resp.Result = result
	return
}

func (resp *ListResponse) GenerateListResponse(code *Response) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = code.Info
	}
	return
}

func (resp *ListResponse) GenerateListResponseWithInfo(code *Response, info string) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = info
	}
	return
}