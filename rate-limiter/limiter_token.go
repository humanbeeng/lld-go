package main

import (
	"sync"
	"time"
)

type TokenBasedLimiter struct {
	IntervalSecs uint
	clientMap    map[string]*Bucket
	sync.RWMutex
}

type Bucket struct {
	Size         uint
	FilledTokens uint
	LastHitOn    time.Time
}

func NewTokenBasedLimiter() TokenBasedLimiter {
	return TokenBasedLimiter{}
}

func (lim TokenBasedLimiter) Allow(req Request) bool {
	return false
}

func (lim *TokenBasedLimiter) refillTokens() {
}
