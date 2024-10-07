package handler

import (
	"gameDev/2_project/internal/resource"
	"log"
	"sync"
	"time"
)

// SafeProcess обрабатывает ресурсы, используя TryLock с тайм-аутом.
func SafeProcess(wg *sync.WaitGroup, res1, res2 *resource.Resource) {
	defer wg.Done()
	log.Printf("Попытка захватить первый ресурс: %s", res1.Name)
	if !res1.TryLock(1 * time.Second) {
		log.Printf("Не удалось захватить первый ресурс: %s, прерывание операции", res1.Name)
		return
	}
	defer res1.Unlock()

	log.Printf("Попытка захватить второй ресурс: %s", res2.Name)
	if !res2.TryLock(1 * time.Second) {
		log.Printf("Не удалось захватить второй ресурс: %s, освобождение первого ресурса", res2.Name)
		return
	}
	defer res2.Unlock()

	// Обработка ресурсов
	res1.Process()
	res2.Process()
	log.Printf("Обработка ресурсов %s и %s завершена", res1.Name, res2.Name)
}
