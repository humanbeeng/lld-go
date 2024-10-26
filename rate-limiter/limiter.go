package main

type Limiter interface {
	Allow(req Request) bool
}
