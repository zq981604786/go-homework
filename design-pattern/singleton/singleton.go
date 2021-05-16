package singleton

import "sync"

// 单例模式
type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
