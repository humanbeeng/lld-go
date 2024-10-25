package main

import "fmt"

func main() {
	fmt.Println("Rate limiter")

	var lim Limiter

	lim = NewTokenBasedLimiter()
	req := Request{
		Key: "one",
	}

	fmt.Println(lim.Allow(&req))
}
