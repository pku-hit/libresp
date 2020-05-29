package libresp

func (resp *Response) IsSuccess() (re bool) {
	if resp == nil {
		re = false
	} else {
		re = resp.Code == SUCC.Code
	}
	return
}

func GenerateResponseSucc() (resp *Response) {
	resp = GenerateResponse(SUCC)
	return
}

func GenerateResponseWithInfo(code *Response, info string) (resp *Response) {
	if code == nil {
		resp = nil
	} else {
		resp = &Response{
			Code: code.Code,
			Info: info,
		}
	}
	return
}

func GenerateResponse(code *Response) (resp *Response) {
	if code == nil {
		resp = nil
	} else {
		resp = &Response{
			Code: code.Code,
			Info: code.Info,
		}
	}
	return
}
