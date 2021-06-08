package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
	c chan time.Time
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
		// 現在時刻を受けとった時間がすぎていたら
		if time.Now().After(t) {
			s.c <- t // チャネルに値をセット セット後、Remindのチャネルの処理が行われる
			cancel()
			wg.Done()
		}
	}
}

func (s Scheduler) Remind(ctx context.Context) {
	for {
		select {
		case t := <-s.c:
			fmt.Println("Remind", t)
		default:
		}
	}
}
