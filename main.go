package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	s := Scheduler{}
	ctx := context.Background()

	wg := sync.WaitGroup{}
	now := time.Now()
	for i:=1; i<=5; i++ {
		wg.Add(1)
		t := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), i*10, 0, time.Local)
		go s.Wait(ctx, &wg, t)
	}
	wg.Wait()
}
