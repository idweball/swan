package engine

import (
	"swan/internal/storage"
	"swan/internal/template"
	"swan/pkg/log"
	"sync"
)

type templateProcessorWatcher struct {
	watch map[string][]*template.Processor
	lock  *sync.RWMutex
}

func newTemplateProcessorWatcher() *templateProcessorWatcher {
	return &templateProcessorWatcher{
		watch: make(map[string][]*template.Processor),
		lock:  new(sync.RWMutex),
	}
}

//Add 关联key与模板处理器
func (watcher *templateProcessorWatcher) Add(key string, processor *template.Processor) {
	watcher.lock.Lock()
	defer watcher.lock.Unlock()

	processors, ok := watcher.watch[key]
	if !ok {
		processors = make([]*template.Processor, 0)
	}

	watcher.watch[key] = append(processors, processor)
}

//Remove 解除key与模板处理器的关联
func (watcher *templateProcessorWatcher) Remove(key string, processor *template.Processor) {
	watcher.lock.Lock()
	defer watcher.lock.Unlock()

	processors, ok := watcher.watch[key]
	if !ok {
		return
	}

	for idx, value := range processors {
		if value == processor {
			watcher.watch[key] = append(processors[:idx], processors[idx+1:]...)
			break
		}
	}
}

//Notify 通知与key关联的模板处理器进行模板渲染
func (watcher *templateProcessorWatcher) Notify(keys []string) []storage.Result {
	watcher.lock.RLock()
	defer watcher.lock.RUnlock()

	processors := make([]*template.Processor, 0)

	for _, key := range keys {
		value, ok := watcher.watch[key]
		if !ok {
			continue
		}
		processors = append(processors, value...)
	}

	processors = removeRepeatProcessor(processors)

	result := make([]storage.Result, 0)

	for _, processor := range processors {
		log.Infof("render template: %s", processor.GetTemplate())
		err := processor.Render()
		if err != nil {
			result = append(result, storage.Result{
				Status:   false,
				Error:    err.Error(),
				Keys:     keys,
				Template: processor.GetTemplate(),
			})
			continue
		}
		result = append(result, storage.Result{
			Status:   true,
			Error:    "",
			Keys:     keys,
			Template: processor.GetTemplate(),
		})
	}

	return result
}

func removeRepeatProcessor(processors []*template.Processor) []*template.Processor {
	result := make([]*template.Processor, 0)
	for _, processor := range processors {
		if !processorInSlice(processor, result) {
			result = append(result, processor)
		}
	}
	return result
}

func processorInSlice(processor *template.Processor, processors []*template.Processor) bool {
	for _, v := range processors {
		if v == processor {
			return true
		}
	}
	return false
}
