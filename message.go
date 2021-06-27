package response

import (
	"fmt"
	"net/http"
	"runtime"
)

func Ok(data interface{}) *writer {
	code := http.StatusOK
	return newWriter(code, StatusSUCCESS, http.StatusText(code), nil, data)
}

func Success(code int, data interface{}) *writer {
	return newWriter(code, StatusSUCCESS, http.StatusText(code), nil, data)
}

func Pending(data interface{}) *writer {
	code := http.StatusOK
	return newWriter(code, StatusPENDING, http.StatusText(code), nil, data)
}

func Failure(code int, err error) *writer {
	return newWriter(code, StatusFAILURE, http.StatusText(code), err, nil)
}

func newWriter(code int, status, message string, err error, data interface{}) *writer {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(2)
		err = fmt.Errorf("%s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	}
	return &writer{
		Code:    code,
		Status:  status,
		Message: message,
		Error:   err,
		Data:    data,
	}
}
