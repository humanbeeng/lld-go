package main

import (
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	sync.RWMutex
	config     TokenBucketConfig
	UserBucket map[string]*Bucket
}

func NewTokenBucketLimiter(cfg ...TokenBucketConfig) TokenBucketLimiter {
	var config TokenBucketConfig
	if len(cfg) == 0 {
		config = TokenBucketConfig{
			RefillSize:            DefaultRefillSize,
			BucketSize:            DefaultBucketSize,
			RefillIntervalSeconds: DefaultRefillIntervalSeconds,
		}
	} else {
		config = cfg[0]
	}

	if config.BucketSize == 0 {
		config.BucketSize = DefaultBucketSize
	}

	if config.RefillIntervalSeconds == 0 {
		config.RefillIntervalSeconds = DefaultRefillIntervalSeconds
	}

	if config.RefillSize == 0 {
		config.RefillSize = DefaultRefillSize
	}

	return TokenBucketLimiter{
		UserBucket: make(map[string]*Bucket),
		config:     config,
	}
}

type TokenBucketConfig struct {
	RefillSize            uint
	BucketSize            uint
	RefillIntervalSeconds uint
}

type Bucket struct {
	BucketSize uint
	Hits       uint
	StartTime  time.Time
}

type Request struct {
	Key string
}

func (lim *TokenBucketLimiter) Allow(req Request) bool {
	lim.RWMutex.Lock()
	defer lim.RWMutex.Unlock()

	ub, ok := lim.UserBucket[req.Key]
	if !ok {

		lim.UserBucket[req.Key] = &Bucket{
			// TODO: Fetch this from config store
			BucketSize: DefaultBucketSize,
			Hits:       1,
			StartTime:  time.Now(),
		}
		return true
	}

	if time.Since(ub.StartTime) > time.Duration(time.Second*DefaultRefillIntervalSeconds) {
		ub.refill()
		ub.StartTime = time.Now()
	}

	if ub.Hits >= ub.BucketSize {
		return false
	}

	ub.Hits++
	return true
}

func (b *Bucket) refill() {
	if DefaultRefillSize > b.BucketSize {
		b.Hits = 0
	} else if b.Hits < DefaultRefillSize {
		b.Hits = 0
	} else {
		b.Hits = b.Hits - DefaultRefillSize
	}
}
