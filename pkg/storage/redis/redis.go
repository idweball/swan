package redis

import (
	"fmt"
	"swan/pkg/log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
	"swan/internal/storage"
)

type redisStorageConfig struct {
	DB       int
	Hostname string
	Port     uint
	Interval int64
	Password string
	BuffSize int64
}

//DefaultBuffSize storage与engine交互的channel的默认大小
const DefaultBuffSize = 1024

//DefaultInterval 向redis取数据的默认间隔
const DefaultInterval = 300

//Driver 用于实现storage.Driver接口
type Driver struct{}

//Open 创建redis 类型的storage
func (d *Driver) Open(cfg map[string]interface{}) (storage.Storage, error) {
	var redisOpt redisStorageConfig
	err := mapstructure.Decode(cfg, &redisOpt)
	if err != nil {
		return nil, err
	}

	rc := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisOpt.Hostname, redisOpt.Port),
		DB:       redisOpt.DB,
		Password: redisOpt.Password,
	})

	_, err = rc.Ping().Result()
	if err != nil {
		return nil, err
	}

	var buffSize int64
	if redisOpt.BuffSize <= 0 {
		buffSize = DefaultBuffSize
	} else {
		buffSize = redisOpt.BuffSize
	}

	var interval int64
	if redisOpt.Interval < 1 {
		interval = DefaultInterval
	} else {
		interval = redisOpt.Interval
	}

	return &Storage{
		rc:       rc,
		buffSize: buffSize,
		stopCh:   make(chan int),
		interval: time.Duration(interval) * time.Second,
		errCh: make(chan error),
	}, nil
}

//Storage redis存储器实现
type Storage struct {
	rc       *redis.Client
	isStop   bool
	stopCh   chan int
	errCh    chan error
	buffSize int64
	interval time.Duration
}

//Get 获取指定keys数据
func (s *Storage) Get(keys []string) (ch chan []storage.Data, errCh chan error, err error) {
	ch = make(chan []storage.Data, s.buffSize)
	s.poll(keys, ch)
	return ch,  s.errCh, nil
}

//Report 报告渲染结果
func (s *Storage) Report(result []storage.Result) {
	for _, r := range result {
		if !r.Status {
			log.Errorf("failed to render template %s, error: %s", r.Template, r.Error)
		}
	}
}

func (s *Storage) getData(keys []string) []storage.Data {
	data := make([]storage.Data, 0)
	values, err := s.rc.MGet(keys...).Result()
	if err != nil {
		s.errCh <- err
	} else {
		for idx, key := range keys {
			if values[idx] == redis.Nil {
				data = append(data, storage.Data{
					Key:   key,
					Value: nil,
				})
			} else {
				data = append(data, storage.Data{
					Key:   key,
					Value: values[idx],
				})
			}
		}
	}
	return data
}

func (s *Storage) poll(keys []string, ch chan []storage.Data) {
	go func() {
		ticker := time.NewTicker(s.interval)
		defer ticker.Stop()

		data := s.getData(keys)
		ch <- data

		for {
			select {
			case <-s.stopCh:
				_ = s.rc.Close()
				s.isStop = true
				close(ch)
				return
			case <-ticker.C:
				data := s.getData(keys)
				ch <- data
			}
		}
	}()
}

//Stop 停止从Storage获取数据
func (s *Storage) Stop() error {
	if s.isStop {
		return nil
	}
	s.stopCh <- 1
	return nil
}

func init() {
	storage.Register("redis", &Driver{})
}
