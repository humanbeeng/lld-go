package main

type Limiter interface {
	Allow(req Request) bool
}

type Request struct {
	Key string
}
