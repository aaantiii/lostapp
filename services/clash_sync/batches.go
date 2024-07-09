package clashsync

import (
	"errors"
	"log/slog"
	"time"
)

// BatchedTask is a task that splits the work into smaller batches.
type BatchedTask[T any] struct {
	nextIndex     int
	task          func(batch []T) error
	interval      time.Duration
	processAmount int
	data          []T
	cancel        chan struct{}
}

func NewBatchedTask[T any](task func([]T) error, interval time.Duration, processAmount int, data []T) *BatchedTask[T] {
	return &BatchedTask[T]{
		task:          task,
		interval:      interval,
		processAmount: processAmount,
		data:          data,
		cancel:        make(chan struct{}, 1),
	}
}

// Run runs the task in batches every interval. After the task is done, it will return.
func (bt *BatchedTask[T]) Run() error {
	for bt.nextIndex < len(bt.data) {
		select {
		case <-time.After(bt.interval):
			start := bt.nextIndex
			end := start + bt.processAmount
			if end > len(bt.data) {
				end = len(bt.data)
			}
			if start > end {
				start = end
			}

			batch := bt.data[start:end]
			bt.nextIndex += bt.processAmount

			if err := bt.task(batch); err != nil {
				slog.Error("Error while running task.", slog.Any("err", err))
			} else {
				slog.Info(
					"Processed batch of task.",
					slog.Int("start", start),
					slog.Int("end", end),
					slog.Int("amount", len(batch)),
				)
			}
		case <-bt.cancel:
			return errors.New("task was canceled")
		}
	}

	return nil
}

func (bt *BatchedTask[T]) Stop() {
	bt.cancel <- struct{}{}
}
