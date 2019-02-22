package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func add(numbers []int) int {
	var v int
	for _, n := range numbers {
		v += n
	}
	return v
}

func addConcurrent(goroutines int, numbers []int) int {
	var v int64
	totalNumbers := len(numbers)
	lastGoroutine := goroutines - 1
	stride := totalNumbers / goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastGoroutine {
				end = totalNumbers
			}

			var lv int
			for _, n := range numbers[start:end] {
				lv += n
			}

			atomic.AddInt64(&v, int64(lv))
			wg.Done()
		}(g)
	}

	wg.Wait()

	return int(v)
}
func BenchmarkSequential(b *testing.B) {
	n := 1000000
	numbers := make([]int, 0, n)
	for i := 0; i < n; i++ {
		numbers = append(numbers, i)
	}
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	n := 1000000
	numbers := make([]int, 0, n)
	for i := 0; i < n; i++ {
		numbers = append(numbers, i)
	}
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}

//func main() {
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
//	addConcurrent(nums, 8)
//	add(nums)
//}
