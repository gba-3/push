package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
}

func (s Scheduler) Wait(ctx context.Context, wg *sync.WaitGroup, t time.Time) {
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		select {
		case <-cancelCtx.Done():
			return
		default:
		}
		if time.Now().After(t) {
			cancel()
			wg.Done()
		}
	}
}

func (s Scheduler) Remind(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}
	fmt.Println(num)
	return nil
}
