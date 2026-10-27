package main

import (
	"math"
	"sync"
)

type CustomWaitGroup struct {
	mu      sync.Mutex
	count   int             // счётчик
	waiters []chan struct{} // вызовы wg.Wait()
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	// не можем увеличивать счётчик бесконечно
	if delta > 0 && wg.count > math.MaxInt-delta {
		panic("CustomWaitGroup: counter overflow")
	}

	wg.count = wg.count + delta
	// не можем уменьшать счётчик ниже нуля
	if wg.count < 0 {
		panic("CustomWaitGroup: negative counter")
	}

	// если достигли нуля, отпускаем всех ожидающих
	if wg.count == 0 {
		for _, semaphore := range wg.waiters {
			semaphore <- struct{}{}
		}
		wg.waiters = nil
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()

	// если счётчик уже равен нулю, то сразу отпускаем ожидающего
	if wg.count == 0 {
		wg.mu.Unlock()
		return
	}

	// иначе начинаем его держать
	semaphore := make(chan struct{}, 1)
	wg.waiters = append(wg.waiters, semaphore)
	wg.mu.Unlock()
	<-semaphore
}
