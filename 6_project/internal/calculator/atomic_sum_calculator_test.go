package calculator

import (
	"testing"
)

// TestCalculateSum проверяет корректность параллельного суммирования массива с атомарными операциями.
func TestCalculateSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedSum := int64(55)

	atomicSumCalculator := NewAtomicSumCalculator()
	totalSum := atomicSumCalculator.CalculateSum(numbers, 2)

	if totalSum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, totalSum)
	}
}
