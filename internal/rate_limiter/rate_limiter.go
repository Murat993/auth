package rate_limiter

import (
	"context"
	"time"
)

// TokenBucketLimiter представляет собой структуру ограничителя скорости,
// который использует механизм токенного ведра.
type TokenBucketLimiter struct {
	tokenBucketCh chan struct{} // Канал для хранения токенов
}

// NewTokenBucketLimiter создает новый экземпляр TokenBucketLimiter.
// ctx - контекст для управления жизненным циклом горутины пополнения токенов.
// limit - максимальное количество токенов, которые могут быть использованы за период.
// period - временной период, за который токены должны быть использованы.
func NewTokenBucketLimiter(ctx context.Context, limit int, period time.Duration) *TokenBucketLimiter {
	limiter := &TokenBucketLimiter{
		tokenBucketCh: make(chan struct{}, limit), // Инициализация канала с размером ведра равным limit
	}

	// Начальное заполнение канала токенами до лимита
	for i := 0; i < limit; i++ {
		limiter.tokenBucketCh <- struct{}{}
	}

	// Вычисление интервала времени между добавлениями токенов в ведро,
	// чтобы обеспечить равномерное распределение в течение заданного периода
	replenishmentInterval := period.Nanoseconds() / int64(limit)
	// Запуск горутины для периодического пополнения токенов
	go limiter.startPeriodicReplenishment(ctx, time.Duration(replenishmentInterval))

	return limiter
}

// startPeriodicReplenishment запускает периодическое пополнение токенов.
// Эта горутина работает в фоне и добавляет новый токен в канал через равные промежутки времени.
func (l *TokenBucketLimiter) startPeriodicReplenishment(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval) // Создание тикера с заданным интервалом
	defer ticker.Stop()                // Остановка тикера при выходе из функции

	for {
		select {
		case <-ctx.Done(): // При получении сигнала об отмене контекста завершаем функцию
			return
		case <-ticker.C:
			l.tokenBucketCh <- struct{}{} // Успешное добавление токена
		}
	}
}

// Allow проверяет, доступен ли токен для выполнения операции.
// Возвращает true, если операция может быть выполнена сразу (токен доступен),
// и false, если в текущий момент токены закончились.
func (l *TokenBucketLimiter) Allow() bool {
	select {
	case <-l.tokenBucketCh: // Попытка взять токен из канала
		return true // Токен взят, операция разрешена
	default:
		return false // Токенов нет, операция не разрешена
	}
}
