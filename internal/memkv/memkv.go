package memkv

import "swan/pkg/memkv"

var kv *memkv.MemKv

//GetKv 获取MemKv的实例
func GetKv() *memkv.MemKv {
	if kv == nil {
		kv = memkv.New()
	}
	return kv
}
