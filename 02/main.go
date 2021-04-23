package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type MyError struct {
	s    string
	code int
}

func NewMyError(s string, code int) MyError {
	return MyError{s, code}
}

func (me MyError) Error() string {
	return me.s
}
func (me MyError) Code() int {
	return me.code
}

func main() {
	//e := errors.New("原始错误e")
	//w := fmt.Errorf("warp 了一个错误:%w", e)
	//fmt.Println(w)
	//fmt.Println(errors.Unwrap(w))
	myErr := NewMyError("not found", 404)
	wrapp1Err := fmt.Errorf("this is a wrapping1 error:%w", myErr)
	wrapp2Err := fmt.Errorf("this is a wrapping2 error:%w", wrapp1Err)
	fmt.Println(errors.Unwrap(wrapp2Err))
	fmt.Println(errors.Is(wrapp2Err, myErr))
	if target, ok := wrapp2Err.(MyError); ok {
		fmt.Println("type assert success:", target.Error())
	}
	var target MyError
	if errors.As(wrapp2Err, &target) {
		fmt.Println("as success:", target.Code())
	}
}
