package engine

import (
	"fmt"
	memkv2 "swan/internal/memkv"
	"swan/internal/storage"
	"swan/internal/template"
	"swan/pkg/log"
	"swan/pkg/memkv"
)

type engine struct {
	kv      *memkv.MemKv
	storage storage.Storage
	keys    []string
	stopCh  chan bool
	isStop  bool
	watcher *templateProcessorWatcher
}

//New 创建engine实例
func New(storage storage.Storage, configs []template.Config) (*engine, error) {
	watcher := newTemplateProcessorWatcher()

	for _, cfg := range configs {
		processor, err := template.NewProcessor(cfg)
		if err != nil {
			return nil, err
		}
		for _, key := range cfg.Keys {
			watcher.Add(key, processor)
		}
	}

	keys := getTemplateKeys(configs)

	return &engine{
		kv:      memkv2.GetKv(),
		storage: storage,
		keys:    keys,
		stopCh:  make(chan bool),
		isStop:  false,
		watcher: watcher,
	}, nil
}

func (e *engine) process(data []storage.Data) {
	keys := make([]string, 0)

	for _, v := range data {
		if e.kv.Equal(v.Key, v.Value) {
			continue
		}

		e.kv.Set(v.Key, v.Value)

		if !elementInStrSlice(v.Key, e.keys) {
			keys = append(keys, v.Key)
		}
	}

	if len(keys) != 0 {
		result := e.watcher.Notify(keys)
		e.storage.Report(result)
	}
}

//Run 启动engine
func (e *engine) Run() error {
	dataCh, errCh, err := e.storage.Get(e.keys)
	if err != nil {
		return err
	}
	for {
		select {
		case data, ok := <-dataCh:
			if !ok {
				return fmt.Errorf("engine: storage channel is closed")
			}
			e.process(data)
		case err = <- errCh:
			log.Errorf("storage error: %v", err)
		case <- e.stopCh:
			e.isStop = true
			return nil
		}
	}
}

//Close 停止engine
func (e *engine) Close() error {
	if e.isStop {
		return nil
	}

	err := e.storage.Stop()
	if err != nil {
		return err
	}
	e.stopCh <- true

	return nil
}

func getTemplateKeys(cfgs []template.Config) []string {
	keys := make([]string, 0)
	for _, cfg := range cfgs {
		keys = append(keys, cfg.Keys...)
	}
	return removeRepElement(keys)
}

func removeRepElement(s []string) []string {
	result := make([]string, 0)

	for _, v := range s {
		if !elementInStrSlice(v, result) {
			result = append(result, v)
		}
	}

	return result
}

func elementInStrSlice(s string, e []string) bool {
	for _, v := range e {
		if s == v {
			return true
		}
	}
	return false
}
