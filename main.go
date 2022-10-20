package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/cr00z/parser/work"
	"golang.org/x/time/rate"
)

const (
	workNum   = 10
	RPS       = 1
	startLink = "http://185.204.3.165"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	limiter := rate.NewLimiter(rate.Limit(RPS), 1)
	workQueue := make(chan *work.Work, workNum)
	workFinished := int32(0)

	// init
	for idx := 1; idx <= workNum; idx++ {
		workQueue <- work.NewWork(idx, ctx, startLink)
	}

	for {
		select {
		case work := <-workQueue:
			log.Printf("Work %d step %d started", work.Idx, work.Step)
			go func() {
				err := work.Do()
				if err != nil {
					log.Println(err)
					work.Finished = true
				}
				if work.Finished {
					atomic.AddInt32(&workFinished, 1)
				} else {
					workQueue <- work
				}
			}()
		case <-termChan:
			log.Println("shutdown signal received")
			cancelFn()
			return // or break loop if needed
		default:
		}

		err := limiter.Wait(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if workFinished == workNum {
			log.Println("all works completed")
			return
		}
	}
}
