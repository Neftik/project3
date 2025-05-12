# Project3 🚀

## О проекте 🛠️

Система для параллельного вычисления арифметических выражений в распределенной среде с авторизацией. Агент создает воркеров, которые считают бинарные выражения. Всё просто и эффективно! 😎

## Структура проекта 📂

```
api/       # gRPC и proto файлы
cmd/       # Запуск оркестратора и агента
internal/  # Логика приложения
pkg/       # Утилиты и вспомогательные пакеты
```

## Требования 🧰

### С Docker (рекомендуется) 🐳

- Docker 24.0+
- Docker Compose

### Без Docker 🛠️

- Go 1.20+
- Protocol Buffers (protoc)

## Установка и запуск 🚀

### Установка

1. Клонируйте репозиторий:

```sh
git clone https://github.com/Neftik/project3.git
cd ./project3
```

2. Установите зависимости:

```sh
go mod tidy
```

### Запуск

1. Создайте файл `.env` в корне и укажите параметры:

```sh
TIME_ADDITION_MS=<ЗНАЧЕНИЕ>
TIME_SUBTRACTION_MS=<ЗНАЧЕНИЕ>
TIME_MULTIPLICATIONS_MS=<ЗНАЧЕНИЕ>
TIME_DIVISIONS_MS=<ЗНАЧЕНИЕ>
COMPUTING_POWER=<ЗНАЧЕНИЕ>
ORCHESTRATOR_ADDRESS=<ЗНАЧЕНИЕ> # "localhost:8080" без Docker; "orchestrator:8080" с Docker
```

2. Запустите проект:

#### С помощью Go 🏃‍♂️

- Оркестратор:

```sh
go run ./cmd/orchestrator/cmd.go
```

- Агент:

```sh
go run ./cmd/agent/cmd.go
```

#### С помощью Docker 🐳

```sh
docker-compose up --build
```

## Компоненты системы 🧩

### Оркестратор 🎛️

- Принимает выражения, разбивает их на операции и отправляет агенту.
- Требуется регистрация и авторизация (JWT токен).

### Агент 🤖

- Получает задачи от оркестратора, выполняет их и возвращает результаты.
- Количество воркеров регулируется `COMPUTING_POWER`.

## Примеры запросов 📡

- **Регистрация:**

```sh
curl --location 'localhost:8080/api/v1/register' \
--header 'Content-Type: application/json' \
--data '{"login": "Neftik", "password": "123"}'
```

- **Логин:**

```sh
curl --location 'localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{"login": "Neftik", "password": "123"}'
```

- **Добавление выражения:**

```sh
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer ваш_токен' \
--data '{"expression": "2+2*2"}'
```

- **Получение списка выражений:**

```sh
curl --location 'localhost:8080/api/v1/expressions' \
--header 'Authorization: Bearer ваш_токен'
```

- **Получение выражения по ID:**

```sh
curl --location 'localhost:8080/api/v1/expressions/1' \
--header 'Authorization: Bearer ваш_токен'
```
