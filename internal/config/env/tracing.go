package env

import (
	"context"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go/config" // Импорт пакета для конфигурации Jaeger трейсера
	"os"
)

const (
	serviceName = "SERVICE_NAME"
)

// NewTraceConfig инициализирует глобальный трейсер Jaeger для сбора трассировок выполнения программы.
// Это позволяет отслеживать и анализировать вызовы между микросервисами и операциями внутри сервиса.
func NewTraceConfig(_ context.Context) (*config.Configuration, error) {
	// Определение конфигурации трейсера
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const", // Использование константного сэмплера
			Param: 1,       // Сбор трассировок для всех запросов (1 = 100% запросов)
		},
		// Другие параметры конфигурации трейсера (например, репортеры) могут быть добавлены здесь
	}

	// Инициализация глобального трейсера с указанным именем сервиса и конфигурацией
	// serviceName используется для идентификации трассировок, принадлежащих этому сервису
	_, err := cfg.InitGlobalTracer(os.Getenv(serviceName))
	if err != nil {
		// В случае ошибки инициализации, логируем критическую ошибку и останавливаем сервис
		// Это обеспечивает, что сервис не будет работать без возможности трассировки
		return nil, errors.New("failed to init tracing")
	}
	// После успешной инициализации, глобальный трейсер Jaeger готов к сбору и отправке трассировок

	return cfg, nil
}
