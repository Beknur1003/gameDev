# gameDev


1/main.go - Пример кода: безопасный захват нескольких мьютексов
2/main.go - Пример более сложного сценария с использованием TryLock через канал
3/main.go - Оптимизированный пример распараллеливания работы с большим слайсом
4/main.go - как и в 3, но имеется обработка состояния гонки (race condition)
5/main.go - Если есть массив цифр и есть одно поле, где должно быть сумма цифр. Как суммировать безопасно в горутинах. При помощи мьютексов