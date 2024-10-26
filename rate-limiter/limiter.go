package main

const (
	DefaultRefillSize             uint = 5
	DefaultBucketSize                  = 10
	DefaultRefillIntervalSeconds       = 1
	DefaultProcessIntervalSeconds      = 1
	DefaultRequestsPickupCount         = 1
)

type Limiter interface {
	Allow(req Request) bool
}
