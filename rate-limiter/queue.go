package main

import "fmt"

type RequestQueue struct {
	MaxBufferSize int
	requests      []*Request
}

func NewRequestQueue(maxBufferSize int) RequestQueue {
	return RequestQueue{
		MaxBufferSize: maxBufferSize,
		requests:      make([]*Request, 0),
	}
}

func (rq *RequestQueue) Push(req *Request) error {
	if len(rq.requests) > rq.MaxBufferSize {
		return fmt.Errorf("requests buffer limit reached")
	}
	rq.requests = append(rq.requests, req)
	return nil
}

func (rq *RequestQueue) Pop() (*Request, error) {
	if len(rq.requests) == 0 {
		return nil, fmt.Errorf("no requests enqueued")
	}
	req := rq.requests[0]
	rq.requests = rq.requests[1:]

	return req, nil
}

func (rq *RequestQueue) Size() int {
	return len(rq.requests)
}
