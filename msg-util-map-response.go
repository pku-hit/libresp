package libresp

import any "github.com/golang/protobuf/ptypes/any"

func (resp *MapResponse) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

func GenerateMapResponseSucc(result map[string]*any.Any) (resp *MapResponse) {
	resp = GenerateMapResponse(SUCCESS)
	resp.Result = result
	return
}

func GenerateMapResponse(code *Response) (resp *MapResponse) {
	resp = &MapResponse{
		Code: code.Code,
		Info: code.Info,
	}
	return
}

func GenerateMapResponseWithInfo(code *Response, info string) (resp *MapResponse) {
	if code == nil {
		resp = nil
	} else {
		resp = &MapResponse{
			Code: code.Code,
			Info: info,
		}
	}
	return
}
