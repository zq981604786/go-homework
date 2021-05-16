package config

import "errors"

var (
	ErrNotFound   = errors.New("key not found")
	ErrTypeAssert = errors.New("type assert error")
)

//type Observer func(string, Value)
//
//type Config interface {
//	Load() error
//	Scan(v interface{}) error
//	Value(key string) Value
//}
