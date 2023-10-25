package benchmark_worker

import (
	"testing"
	"time"

	"github.com/gammazero/workerpool"

	"github.com/devchat-ai/gopool"
)

const (
	PoolSize = 1e4
	TaskNum  = 1e6
)

// 1	1235387792 ns/op	 1952832 B/op	   14179 allocs/op
func BenchmarkGoPool(b *testing.B) {
	pool := gopool.NewGoPool(PoolSize)
	defer pool.Release()

	taskFunc := func() (any, error) {
		time.Sleep(10 * time.Millisecond)
		return nil, nil
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for num := 0; num < TaskNum; num++ {
			pool.AddTask(taskFunc)
		}
	}
	pool.Wait()
	b.StopTimer()
}

// 1	1290887458 ns/op	34749896 B/op	   60666 allocs/op
func BenchmarkWorkerPool(b *testing.B) {
	pool := workerpool.New(PoolSize)

	taskFunc := func() {
		time.Sleep(10 * time.Millisecond)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for num := 0; num < TaskNum; num++ {
			pool.Submit(taskFunc)
		}
	}
	pool.StopWait()
	b.StopTimer()
}
