package main

import "fmt"

type DialOption func(options *dialOptions)

type dialOptions struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	TimeOut int
	Auth    string
}

func DialAuth(auth string) DialOption {
	return func(options *dialOptions) {
		options.Auth = auth
	}
}

func Dial(name, addr string, options ...DialOption) *dialOptions {
	do := &dialOptions{
		Name: name,
		Addr: addr,
	}
	for _, option := range options {
		option(do)
	}
	return do
}

func main() {
	d := Dial("111", "222", DialAuth("2222"))
	fmt.Println("111", d)

}
