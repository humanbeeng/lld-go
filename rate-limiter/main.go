package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Rate limiter")

	var lim Limiter

	tbl := NewTokenBucketLimiter()

	lim = &tbl

	req1 := Request{
		Key: "one",
	}

	hitCount := 0
	missCount := 0

	var wg sync.WaitGroup
	for i := range 20 {
		wg.Add(1)

		time.Sleep(time.Duration(time.Millisecond * 100))
		go func(wg *sync.WaitGroup) {
			var allowed bool
			allowed = lim.Allow(req1)

			if allowed {
				hitCount++
				fmt.Println(i, "200: Request allowed")
			} else {
				missCount++
				fmt.Println(i, "429: Too many requests")
			}

			defer wg.Done()
		}(&wg)

	}

	wg.Wait()

	fmt.Println("Hit count", hitCount)
	fmt.Println("Miss count", missCount)
}
