package resource

import (
	"log"
	"sync"
	"time"
)

// Resource представляет структуру игрового ресурса, используемого для блокировки.
type Resource struct {
	Name string
	mu   sync.Mutex
}

// NewResource создает новый игровой ресурс.
func NewResource(name string) *Resource {
	return &Resource{Name: name}
}

// Lock блокирует доступ к игровому ресурсу.
func (r *Resource) Lock() {
	r.mu.Lock()
	log.Printf("Игровой ресурс %s заблокирован", r.Name)
}

// Unlock разблокирует доступ к игровому ресурсу.
func (r *Resource) Unlock() {
	r.mu.Unlock()
	log.Printf("Игровой ресурс %s разблокирован", r.Name)
}

// TryLock пытается заблокировать доступ к игровому ресурсу в течение указанного таймаута.
func (r *Resource) TryLock(timeout time.Duration) bool {
	locked := make(chan struct{}, 1)

	go func() {
		r.mu.Lock()
		locked <- struct{}{}
	}()

	select {
	case <-locked:
		log.Printf("Игровой ресурс %s успешно заблокирован", r.Name)
		return true
	case <-time.After(timeout):
		log.Printf("Не удалось заблокировать игровой ресурс %s в течение таймаута", r.Name)
		return false
	}
}

// Process симулирует обработку игрового ресурса.
func (r *Resource) Process() {
	time.Sleep(1 * time.Second) // Симуляция работы с игровым ресурсом
	log.Printf("Обработка игрового ресурса %s завершена", r.Name)
}
