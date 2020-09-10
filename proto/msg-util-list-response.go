package proto

import "github.com/golang/protobuf/ptypes/any"

func (resp *ListResponse) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

func (resp *ListResponse) GenerateListResponseSucc(result []*any.Any) {
	resp.GenerateListResponse(SUCCESS)
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
