package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Info struct {
	sonCount int
}

type Worker struct {
	cancelFunc context.CancelFunc
	ctx        context.Context
	id         int
}

func (w *Worker) Do() {
	w.ctx.Value("info").(*Info).sonCount++
	for {
		select {
		case <-w.ctx.Done():
			fmt.Println("阿我死了", w.id)
			w.ctx.Value("info").(*Info).sonCount--
			return
		default:
			fmt.Println("Do", w.id)
		}
		time.Sleep(time.Duration(rand.Intn(5) * int(time.Second)))
	}
}

type Boss struct {
	workers []*Worker
	ctx     context.Context
}

func main() {
	boss := &Boss{
		ctx: context.WithValue(context.Background(), "info", &Info{}),
	}
	for i := 0; i < 10; i++ {
		ctx, cancelFunc := context.WithCancel(boss.ctx)
		worker := &Worker{
			ctx:        ctx,
			cancelFunc: cancelFunc,
			id:         i,
		}
		boss.workers = append(boss.workers, worker)
		go worker.Do()
	}

	go func() {
		for {
			fmt.Println("sonCount", boss.ctx.Value("info").(*Info).sonCount)
			time.Sleep(1 * time.Second)
		}
	}()
	for _, worker := range boss.workers {
		time.Sleep(1 * time.Second)
		fmt.Println("Go to bed~",worker.id)
		worker.cancelFunc()
	}
	time.Sleep(1000000 * time.Second)
}

/*
	总结:
		1. select 会等待一个case满足，满足后便执行，执行后便退出select.
*/
