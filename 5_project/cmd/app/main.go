package main

import (
	"gameDev/5_project/internal/calculator"
	"gameDev/5_project/internal/log"
	log2 "log"
)

func main() {
	// Инициализация логгера
	log.InitLogger()

	// Инициализация массива для суммирования
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Число горутин для параллельной обработки
	numGoroutines := 2

	// Создание экземпляра калькулятора
	sumCalculator := calculator.NewSumCalculator()

	// Параллельное суммирование элементов массива
	totalSum := sumCalculator.CalculateSum(numbers, numGoroutines)

	// Логирование результата
	log2.Printf("Итоговая сумма: %d\n", totalSum)
}
