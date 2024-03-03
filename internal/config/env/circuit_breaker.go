package env

import (
	"context"
	"github.com/sony/gobreaker"
	"log"
	"time"
)

// Circuit Breaker — это паттерн, который предотвращает распространение сбоев в системе путем временного
// прекращения выполнения операций, которые могут вызвать ошибку.
// Он имеет три основных состояния: закрытое, открытое и полуоткрытое.

// Закрытое (Closed): Все запросы выполняются. Если ошибок слишком много, переходит в открытое.

// Открытое (Open): Все запросы автоматически блокируются, чтобы дать системе время на восстановление.
// После тайм-аута переходит в полуоткрытое.

// Полуоткрытое (Half-Open): Ограниченное количество запросов проверяется на успех.
// При успехе возвращается в закрытое, при неудаче — в открытое.

type circuitBreakerConfig struct {
	circuitBreaker *gobreaker.CircuitBreaker
}

func NewCircuitBreakerConfig(_ context.Context) *gobreaker.CircuitBreaker {
	// Создание нового Circuit Breaker с конфигурацией.
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		// Имя Circuit Breaker для идентификации, может использоваться в логах.
		Name: "my-service",

		// Максимальное количество запросов, которое Circuit Breaker позволяет обработать
		// в полуоткрытом состоянии, прежде чем определить, следует ли перейти обратно в закрытое состояние.
		MaxRequests: 3,

		// Время ожидания перед переключением Circuit Breaker из открытого состояния в полуоткрытое,
		// позволяя новым запросам пройти для проверки, восстановилась ли система.
		Timeout: 5 * time.Second,

		// Функция, определяющая условия, при которых Circuit Breaker должен "сработать" и перейти в открытое состояние.
		// В данном случае, если отношение ошибок к общему числу запросов превышает 60%.
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6
		},

		// Функция обратного вызова, которая срабатывает при каждом изменении состояния Circuit Breaker.
		// Полезно для логирования и мониторинга состояний Circuit Breaker.
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit Breaker: %s, changed from %v, to %v\n", name, from, to)
		},
	})

	return cb
}
