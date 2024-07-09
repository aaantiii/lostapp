package clashsync

import (
	"log/slog"
	"time"
)

type Task interface {
	Run() error
	Stop()
}

// Scheduler is a simple task scheduler that runs a task based on time.
type Scheduler struct {
	taskFactory func() Task
	cancel      chan struct{}
}

// NewScheduler creates a new scheduler. The factory function returns a new task instance every time the task is run.
func NewScheduler(factory func() Task) *Scheduler {
	return &Scheduler{
		taskFactory: factory,
		cancel:      make(chan struct{}),
	}
}

// RunEvery runs the task in its own goroutine every interval, until Stop is called.
func (s *Scheduler) RunEvery(interval time.Duration) {
	go func() {
		for {
			select {
			case <-time.After(interval):
				if err := s.taskFactory().Run(); err != nil {
					slog.Error("Error while running task.", slog.Any("err", err))
				}
			case <-s.cancel:
				return
			}
		}
	}()
}

// Stop stops the scheduler immediately.
func (s *Scheduler) Stop() {
	s.cancel <- struct{}{}
}
