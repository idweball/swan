package storage

import (
	"fmt"
	"sync"
)

//Data 存储器与Engine之间的交流数据
type Data struct {
	Key   string
	Value interface{}
}

//Result 模板渲染后的结果, 此结果交由Storage进行处理
type Result struct {
	Status   bool     //true:渲染成功，false:渲染失败
	Error    string   //渲染失败的错误消息
	Keys     []string //此次渲染涉及哪些key
	Template string   //此次渲染的所使用的模板文件
}

//Driver 自定义Storage需要实现此接口进行注册
type Driver interface {
	Open(map[string]interface{}) (Storage, error)
}

//Storage 配置存储器的接口
type Storage interface {
	Get(keys []string) (ch chan []Data, ech chan error, err error)
	Stop() error
	Report([]Result)
}

var (
	drivers     = make(map[string]Driver)
	driversLock sync.RWMutex
)

//Register 注册storage名称, storage的构造函数
func Register(name string, driver Driver) {
	driversLock.Lock()
	defer driversLock.Unlock()

	if _, dup := drivers[name]; dup {
		panic("storage: register called twice for storage " + name)
	}

	drivers[name] = driver
}

//List 已注册的存储器列表
func List() (s []string) {
	driversLock.Lock()
	defer driversLock.Unlock()

	for k := range drivers {
		s = append(s, k)
	}
	return
}

//New 根据name从已注册的drivers获取对应的构造函数，创建storage
func New(name string, cfg map[string]interface{}) (Storage, error) {
	driversLock.Lock()
	defer driversLock.Unlock()

	driver, ok := drivers[name]
	if !ok {
		return nil, fmt.Errorf("storage: not found storage %s", name)
	}

	return driver.Open(cfg)
}
