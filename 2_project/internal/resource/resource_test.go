package resource

import (
	"testing"
	"time"
)

// Тестирует создание и блокировку игрового ресурса.
func TestResource(t *testing.T) {
	res := NewResource("Test Game Resource")

	// Проверка блокировки и разблокировки
	res.Lock()
	go func() {
		time.Sleep(100 * time.Millisecond)
		res.Unlock()
	}()

	// Если ресурс заблокирован, вызов метода снова вызовет deadlock.
	locked := make(chan bool)
	go func() {
		res.Lock()
		locked <- true
		res.Unlock()
	}()

	select {
	case <-locked:
		t.Log("Игровой ресурс успешно заблокирован и разблокирован")
	case <-time.After(1 * time.Second):
		t.Error("Ошибка блокировки/разблокировки игрового ресурса")
	}
}
