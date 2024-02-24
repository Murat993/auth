package tests

// Импортируем необходимые пакеты: стандартные библиотеки Go, библиотеки для генерации данных, мокирования, утверждений,
// а также внутренние модули для работы с API, моделями и сервисами.
import (
	"context"
	"fmt"
	"github.com/Murat993/auth/internal/api/user"
	"github.com/Murat993/auth/internal/dto"
	"github.com/Murat993/auth/internal/service"
	"github.com/Murat993/auth/internal/service/mocks"
	desc "github.com/Murat993/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v6"
	"testing"

	"github.com/gojuno/minimock/v3"       // Для мокирования зависимостей
	"github.com/stretchr/testify/require" // Для утверждений в тестах
	// Внутренние пакеты проекта
)

// TestCreate - функция для тестирования создания заметки.
func TestCreate(t *testing.T) {
	t.Parallel() // Запускаем тесты параллельно для ускорения выполнения.

	// Определяем типы и переменные для использования в тестах.
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	//  Структура для аргументов, которые будут переданы в тестируемую функцию.
	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	// Подготавливаем контекст, минимок контроллер и тестовые данные.
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id              = gofakeit.UUID() // Генерируем тестовые данные
		title           = gofakeit.Animal()
		email           = gofakeit.Email()
		password        = gofakeit.Password(true, true, true, true, true, 64)
		passwordConfirm = gofakeit.Password(true, true, true, true, true, 64)
		role            = desc.Role_USER

		serviceErr = fmt.Errorf("service error") // Ошибка для имитации сбоя в сервисе

		// Создаем запрос на создание заметки
		req = &desc.CreateRequest{
			UseCreate: &desc.UserCreate{
				Name:            title,
				Email:           email,
				Password:        password,
				PasswordConfirm: passwordConfirm,
				Role:            role,
			},
		}

		// Информация для создания заметки
		userCreateDto = &dto.UserCreate{
			Name:            title,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role,
		}

		// Ожидаемый ответ от сервиса
		res = &desc.CreateResponse{
			Id: id,
		}
	)
	defer t.Cleanup(mc.Finish) // Очищаем ресурсы после выполнения каждого теста.

	// Определяем тестовые случаи.
	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		// Тест на успешное создание заметки
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, userCreateDto).Return(id, nil) // Настраиваем мок для имитации успешного создания
				return mock
			},
		},
		// Тест на случай ошибки в сервисе
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, userCreateDto).Return("", serviceErr) // Настраиваем мок для имитации ошибки сервиса
				return mock
			},
		},
	}

	// Итерируем по тестовым случаям и запускаем каждый тест.
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Запускаем тесты в параллельном режиме для каждого тестового случая.

			// Создаем мок сервиса заметок с помощью переданной функции.
			userServiceMock := tt.userServiceMock(mc)
			// Инициализируем API с моком сервиса.
			api := user.NewImplementation(userServiceMock)

			// Вызываем метод Create API и получаем результат.
			newID, err := api.Create(tt.args.ctx, tt.args.req)

			// Проверяем, соответствует ли полученная ошибка ожидаемой ошибке.
			require.Equal(t, tt.err, err)
			// Проверяем, соответствует ли возвращенный результат ожидаемому результату.
			require.Equal(t, tt.want, newID)
		})
	}
}
