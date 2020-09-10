package libresp

import any "github.com/golang/protobuf/ptypes/any"

func (resp *ListResponse) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

func GenerateListResponseSucc(result []*any.Any) (resp *ListResponse) {
	resp = GenerateListResponse(SUCCESS)
	resp.Result = result
	return
}

func GenerateListResponse(code *Response) (resp *ListResponse) {
	resp = &ListResponse{
		Code: code.Code,
		Info: code.Info,
	}
	return
}

func GenerateListResponseWithInfo(code *Response, info string) (resp *ListResponse) {
	if code == nil {
		resp = nil
	} else {
		resp = &ListResponse{
			Code: code.Code,
			Info: info,
		}
	}
	return
}
