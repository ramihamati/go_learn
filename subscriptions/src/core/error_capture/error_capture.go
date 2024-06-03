package error_capture

import (
	"log"
	"subscriptions/core/errors"
)

type ErrorCapture struct {
	Error interface{}
}

func Capture(err interface{}) ErrorCapture {
	return ErrorCapture{
		Error: err,
	}
}

func (capture *ErrorCapture) PrintError() {
	if v, ok := capture.Error.(errors.CodeError); ok {
		println(v.Error())
		return
	}
	if v, ok := capture.Error.(errors.StateCodeError); ok {
		println(v.Error())
		return
	}
	if v, ok := capture.Error.(errors.DatabaseError); ok {
		println(v.Error())
		return
	}
	log.Println("Unknown panic occurred:", capture.Error)
}

func CaptureDefer() {
	if err := recover(); err != nil {
		capture := Capture(err)
		capture.PrintError()
	}
}
