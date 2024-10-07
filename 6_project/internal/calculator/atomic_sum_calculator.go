package calculator

import (
	"sync"
	"sync/atomic"
)

// AtomicSumCalculator управляет параллельным суммированием массива с использованием атомарных операций.
type AtomicSumCalculator struct {
	totalSum int64
}

// NewAtomicSumCalculator создает новый экземпляр калькулятора.
func NewAtomicSumCalculator() *AtomicSumCalculator {
	return &AtomicSumCalculator{}
}

// CalculateSum параллельно суммирует элементы массива с помощью атомарных операций.
func (asc *AtomicSumCalculator) CalculateSum(numbers []int, numGoroutines int) int64 {
	chunkSize := len(numbers) / numGoroutines
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(numbers) // Последняя горутина обрабатывает остаток массива
		}

		go func(nums []int) {
			defer wg.Done()
			localSum := int64(0)
			for _, num := range nums {
				localSum += int64(num)
			}

			atomic.AddInt64(&asc.totalSum, localSum)
		}(numbers[start:end])
	}

	wg.Wait()
	return atomic.LoadInt64(&asc.totalSum)
}
