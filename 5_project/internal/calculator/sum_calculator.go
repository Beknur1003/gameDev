package calculator

import (
	"sync"
)

// SumCalculator управляет параллельной обработкой массива для суммирования.
type SumCalculator struct {
	mu       sync.Mutex
	totalSum int
}

// NewSumCalculator создает новый экземпляр калькулятора сумм.
func NewSumCalculator() *SumCalculator {
	return &SumCalculator{}
}

// CalculateSum параллельно суммирует элементы массива, используя указанное количество горутин.
func (sc *SumCalculator) CalculateSum(numbers []int, numGoroutines int) int {
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
			localSum := 0
			for _, num := range nums {
				localSum += num
			}

			sc.mu.Lock()
			sc.totalSum += localSum
			sc.mu.Unlock()
		}(numbers[start:end])
	}

	wg.Wait()
	return sc.totalSum
}
