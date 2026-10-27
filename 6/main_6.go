package main

import (
	"context"
	"math/rand"
)

func generate(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- rand.Int():
			}
		}
	}()
	return ch
}
