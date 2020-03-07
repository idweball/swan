package template

import (
	"reflect"
	"sync"
)

var (
	funcMap  = newFuncMap()
	funcLock = new(sync.RWMutex)
)

//Register 注册模板渲染函数
func Register(name string, fn interface{}) {
	funcLock.Lock()
	defer funcLock.Unlock()

	if !isFn(fn) {
		panic("template:" + name + "is not a function type")
	}

	_, dup := funcMap[name]
	if dup {
		panic("template: register called twice for " + name)
	}

	funcMap[name] = fn
}

//List 已注册的渲染函数列表
func List() []string {
	funcLock.Lock()
	defer funcLock.Unlock()

	s := make([]string, 0)

	for k := range funcMap {
		s = append(s, k)
	}

	return s
}

func isFn(fn interface{}) bool {
	return reflect.TypeOf(fn).Kind() == reflect.Func
}