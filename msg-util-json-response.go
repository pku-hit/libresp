package libresp

func (resp *JsonResponse) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

func (resp *JsonResponse) GenerateJsonResponseSucc(result string) {
	resp.GenerateJsonResponse(SUCCESS)
	resp.Result = result
	return
}

func (resp *JsonResponse) GenerateJsonResponse(code *Response) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = code.Info
	}
	return
}

func (resp *JsonResponse) GenerateJsonResponseWithInfo(code *Response, info string) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = info
	}
	return
}
