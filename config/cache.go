package config

import (
	"sync"
	"sync/atomic"
	"time"
)

// cache will update in blow time
const updateDuration = time.Minute

// traceOpen is special it use atomic to define the whether open-trace is allow or not
var isTraceOpen int32

// IsTraceOpen indicate whether the open-trace open or not
func IsTraceOpen() bool {
	return atomic.LoadInt32(&isTraceOpen) == 1
}

// UpdateOpenTraceConfig force config update asynchronous
func UpdateOpenTraceConfig() {
	// open trace
	isTraceOpenInt := GetInt("openTrace.isTraceOpen")
	atomic.SwapInt32(&isTraceOpen, int32(isTraceOpenInt))
}

// cache cache config and will update in sure time,
// and there is not need for multiple cache implement
type cache struct {
	UpdateTime time.Duration
	Map        sync.Map
}

// Update will find resource where resource is comming and fresh config cache
func (c *cache) Update() {
	c.Map.Range(func(key, value interface{}) bool {
		c.Map.Store(key, Get(key.(string)))
		return true
	})
}

// Get will find cache first; if result is not exit
// then it will find it from reflect of config
func (c *cache) Get(key string) interface{} {
	if value, ok := c.Map.Load(key); ok {
		return value
	}

	value := Get(key)
	go c.Map.Store(key, value)
	return value
}

// QueryCache will cache the config result and will getUpdate when time is come
var QueryCache cache

func init() {
	// 定期更新配置
	QueryCache = cache{
		UpdateTime: updateDuration,
	}
	go func() {
		for range time.NewTicker(updateDuration).C {
			UpdateOpenTraceConfig()
			QueryCache.Update()
		}
	}()
}
