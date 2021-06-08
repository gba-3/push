package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	wg := sync.WaitGroup{}
	now := time.Now()
	num := 5

	c := make(chan time.Time)
	s := Scheduler{c}
	for i := 1; i <= num; i++ {
		wg.Add(1)
		t := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), i, 0, time.Local)
		go s.Wait(ctx, &wg, t)
	}
	go s.Remind(ctx)
	wg.Wait()
}
