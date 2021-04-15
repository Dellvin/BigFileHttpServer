package errors

import "encoding/json"

func NotPost() []byte {
	err := &Response{
		Code:        400,
		Description: "Do not require request's method, expected POST",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func NotGet() []byte {
	err := &Response{
		Code:        400,
		Description: "Do not require request's method, expected GET",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func UploadResponse(text string, code uint16) []byte {
	err := &Response{
		Code:        code,
		Description: text,
	}
	ans, _ := json.Marshal(err)
	return ans
}

func DownloadResponse(text string, code uint16) []byte {
	err := &Response{
		Code:        code,
		Description: text,

	}
	ans, _ := json.Marshal(err)
	return ans
}
