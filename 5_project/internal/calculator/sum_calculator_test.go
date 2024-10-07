package calculator

import (
	"testing"
)

// TestCalculateSum проверяет корректность параллельного суммирования массива.
func TestCalculateSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedSum := 55

	sumCalculator := NewSumCalculator()
	totalSum := sumCalculator.CalculateSum(numbers, 2)

	if totalSum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, totalSum)
	}
}
