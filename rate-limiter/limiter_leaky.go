package main

import (
	"sync"
	"time"
)

type LeakyBucketLimiter struct {
	sync.RWMutex
	config       LeakyBucketConfig
	RequestQueue RequestQueue
	startTime    time.Time
}

type LeakyBucketConfig struct {
	MaxBufferSize          int
	ProcessIntervalSeconds int
	RequestsPickupCount    int
}

func NewLeakyBucketLimiter(config ...LeakyBucketConfig) LeakyBucketLimiter {
	var cfg LeakyBucketConfig

	if len(config) == 0 {
		cfg = LeakyBucketConfig{
			MaxBufferSize:          DefaultBucketSize,
			ProcessIntervalSeconds: DefaultProcessIntervalSeconds,
			RequestsPickupCount:    DefaultRequestsPickupCount,
		}
	} else {
		cfg = config[0]
	}

	// TODO: Set default values for nil values in config

	return LeakyBucketLimiter{
		config:       cfg,
		RequestQueue: NewRequestQueue(int(cfg.MaxBufferSize)),
		startTime:    time.Now(),
	}
}

func (lim *LeakyBucketLimiter) Allow(req Request) bool {

	lim.RWMutex.Lock()
	defer lim.RWMutex.Unlock()

	if lim.RequestQueue.Size() < lim.config.MaxBufferSize {
		lim.RequestQueue.Push(&req)
		return true
	}

	if time.Since(lim.startTime) > time.Second*time.Duration(lim.config.ProcessIntervalSeconds) {
		// start sending requests to the server
		for i := range lim.config.RequestsPickupCount {

		}
	}

	return true
}
