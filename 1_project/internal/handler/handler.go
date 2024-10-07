package handler

import (
	"gameDev/1_project/internal/resource"
	"log"
	"sync"
)

// ProcessResources обрабатывает два игровых ресурса, захватывая мьютексы для безопасности.
func ProcessResources(wg *sync.WaitGroup, id string, res1, res2 *resource.Resource) {
	defer wg.Done()
	log.Printf("Горутина %s: Начало обработки игровых ресурсов\n", id)

	// Блокировка первого игрового ресурса
	res1.Lock()
	defer res1.Unlock()

	// Обработка первого ресурса
	res1.Process()

	// Блокировка второго игрового ресурса
	res2.Lock()
	defer res2.Unlock()

	// Обработка второго ресурса
	res2.Process()

	log.Printf("Горутина %s: Завершение обработки игровых ресурсов", id)
}
