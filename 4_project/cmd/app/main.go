package main

import (
	"gameDev/4_project/internal/log"
	"gameDev/4_project/internal/processor"
	log2 "log"
)

func main() {
	// Инициализация логгера
	log.InitLogger()

	// Создание большого слайса для обработки
	largeSlice := make([]int, 1_000_000)
	for i := 0; i < len(largeSlice); i++ {
		largeSlice[i] = i
	}

	// Число горутин для параллельной обработки
	numGoroutines := 10

	// Создание процессора для работы со слайсами
	sliceProcessor := processor.NewSliceProcessor()

	// Параллельная обработка слайса
	totalSum := sliceProcessor.ProcessSlice(largeSlice, numGoroutines)

	// Логирование результатов
	log2.Printf("Общая сумма элементов слайса: %d\n", totalSum)
	log2.Println("Все части обработаны.")
}
