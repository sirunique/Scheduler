package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/sirunique/scheduler/scheduler"
)

func main() {
	ctx := context.Background()

	worker := scheduler.Scheduler()
	worker.Add(ctx, processTransactionData, time.Second*5)
	worker.Add(ctx, sendTransactionData, time.Second*10)
	worker.Add(ctx, testFunc, time.Second*2)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}

func processTransactionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf("Transaction Data process successfuly at %s\n", time.Now())
}

func sendTransactionData(ctx context.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("Transaction Data sent at %s\n", time.Now().String())
}

func testFunc(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, time.Second*5)

	i := 0
	for {
		time.Sleep(time.Millisecond * 100)
		i++
		fmt.Printf("%d ", 1)

		select {
		case <-ctx.Done():
			fmt.Println()
			return
		default:
		}
	}
}
