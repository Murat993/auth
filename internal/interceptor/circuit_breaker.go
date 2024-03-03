package interceptor

import (
	"context"
	// Импорт пакетов для работы с контекстом, Circuit Breaker и gRPC
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CircuitBreakerInterceptor - структура интерсептора, содержащая Circuit Breaker.
type CircuitBreakerInterceptor struct {
	cb *gobreaker.CircuitBreaker // Ссылка на экземпляр Circuit Breaker.
}

// NewCircuitBreakerInterceptor создает новый экземпляр интерсептора с Circuit Breaker.
func NewCircuitBreakerInterceptor(cb *gobreaker.CircuitBreaker) *CircuitBreakerInterceptor {
	return &CircuitBreakerInterceptor{
		cb: cb, // Инициализация с предоставленным Circuit Breaker.
	}
}

// Unary - метод интерсептора для унарных вызовов gRPC.
func (c *CircuitBreakerInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Выполнение обработчика запроса через Circuit Breaker.
	res, err := c.cb.Execute(func() (interface{}, error) {
		// Вызов реального обработчика gRPC запроса.
		return handler(ctx, req)
	})

	// Обработка ошибок после выполнения запроса через Circuit Breaker.
	if err != nil {
		// Проверка, если ошибка связана с открытым состоянием Circuit Breaker (слишком много ошибок в прошлом).
		if err == gobreaker.ErrOpenState {
			// Возвращение ошибки gRPC клиенту, указывающей на временную недоступность сервиса.
			return nil, status.Error(codes.Unavailable, "service unavailable")
		}

		// Возвращение других ошибок, возникших во время выполнения запроса.
		return nil, err
	}

	// Возвращение результата выполнения запроса, если ошибок не возникло.
	return res, nil
}
