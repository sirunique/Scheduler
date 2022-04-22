package scheduler

import (
	"context"
	"sync"
	"time"
)

type SchedulerStruct struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

func Scheduler() *SchedulerStruct {
	return &SchedulerStruct{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

type Job func(ctx context.Context)

// Add function init goroutine that constantly calls provided job with interval delay
func (s *SchedulerStruct) Add(ctx context.Context, j Job, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)

	s.wg.Add(1)
	// go s.process(t)
}

// Stop cancel all running jobs
func (s *SchedulerStruct) Stop() {
	for _, cancel := range s.cancellations {
		cancel()
	}
	s.wg.Wait()
}

func (s *SchedulerStruct) process(ctx context.Context, j Job, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			j(ctx)
		case <-ctx.Done():
			s.wg.Done()
			ticker.Stop()
			return
		}
	}
}
