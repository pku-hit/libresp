package libresp

func (resp *Response) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCCESS.Code
	}
	return
}

func (resp *Response) GenerateResponseSucc() {
	resp.GenerateResponse(SUCCESS)
	return
}

func (resp *Response) GenerateResponseWithInfo(code *Response, info string) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = info
	}
	return
}

func (resp *Response) GenerateResponse(code *Response) {
	if code == nil || resp == nil {
		return
	} else {
		resp.Code = code.Code
		resp.Info = code.Info
	}
	return
}
