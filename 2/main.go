package main

import (
	"fmt"
	"sync"
	"time"
)

// Пример более сложного сценария с использованием TryLock через канал

type Resource struct {
	name string
}

var resource1 = Resource{name: "Resource 1_project"}
var resource2 = Resource{name: "Resource 2"}

var mu1 = sync.Mutex{}
var mu2 = sync.Mutex{}

// TryLock — попытка захватить мьютекс с тайм-аутом
func TryLock(mu *sync.Mutex, timeout time.Duration) bool {
	locked := make(chan struct{}, 1)

	go func() {
		mu.Lock()
		locked <- struct{}{}
	}()

	select {
	case <-locked:
		return true
	case <-time.After(timeout):
		return false
	}
}

func safeProcess(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Попытка захватить mu1")
	if !TryLock(&mu1, 1*time.Second) {
		fmt.Println("Не удалось захватить mu1, прерывание операции")
		return
	}
	defer mu1.Unlock()
	fmt.Println("mu1 захвачен")

	fmt.Println("Попытка захватить mu2")
	if !TryLock(&mu2, 1*time.Second) {
		fmt.Println("Не удалось захватить mu2, освобождение mu1")
		return
	}
	defer mu2.Unlock()
	fmt.Println("mu2 захвачен")

	fmt.Printf("Обработка ресурсов %s и %s завершена\n", resource1.name, resource2.name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go safeProcess(&wg)
	go safeProcess(&wg)

	wg.Wait()
	fmt.Println("Все операции завершены успешно")
}
