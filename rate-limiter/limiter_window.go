package main

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	sync.RWMutex
	config     FixedWindowConfig
	UserWindow map[string]*Window
}

func NewFixedWindowLimiter(cfg ...FixedWindowConfig) FixedWindowLimiter {
	var config FixedWindowConfig
	if len(cfg) == 0 {
		config = FixedWindowConfig{
			WindowSize:           DefaultBucketSize,
			ResetIntervalSeconds: DefaultRefillIntervalSeconds,
		}
	} else {
		config = cfg[0]
	}

	if config.WindowSize == 0 {
		config.WindowSize = DefaultBucketSize
	}

	if config.ResetIntervalSeconds == 0 {
		config.ResetIntervalSeconds = DefaultRefillIntervalSeconds
	}

	return FixedWindowLimiter{
		UserWindow: make(map[string]*Window),
		config:     config,
	}
}

type FixedWindowConfig struct {
	WindowSize           uint
	ResetIntervalSeconds uint
}

type Window struct {
	WindowSize uint
	Hits       uint
	StartTime  time.Time
}

func (lim *FixedWindowLimiter) Allow(req Request) bool {
	lim.RWMutex.Lock()
	defer lim.RWMutex.Unlock()

	ub, ok := lim.UserWindow[req.Key]
	if !ok {

		lim.UserWindow[req.Key] = &Window{
			// TODO: Fetch this from config store
			WindowSize: lim.config.WindowSize,
			Hits:       1,
			StartTime:  time.Now(),
		}
		return true
	}

	if time.Since(ub.StartTime) > time.Second*time.Duration(lim.config.ResetIntervalSeconds) {
		ub.reset()
		ub.StartTime = time.Now()
	}

	if ub.Hits >= ub.WindowSize {
		return false
	}

	ub.Hits++
	return true
}

func (w *Window) reset() {
	fmt.Println("resetting")
	w.Hits = 0
}
