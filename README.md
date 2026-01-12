# RecruitAI

Система автоматической оценки соответствия кандидатов вакансиям с использованием искусственного интеллекта (GPT-4.1).

## Описание

RecruitAI — это backend-приложение на Go, которое анализирует резюме кандидатов и определяет их соответствие требованиям вакансии с помощью OpenAI GPT API. Система загружает резюме, анализирует его вместе с описанием вакансии и возвращает результат оценки соответствия.

## Технологии

- **Go 1.24+** — основной язык программирования
- **PostgreSQL 17** — база данных
- **GORM** — ORM для работы с базой данных
- **Gorilla Mux** — HTTP роутер
- **OpenAI GPT-4.1 API** — для анализа резюме
- **Docker & Docker Compose** — для контейнеризации и оркестрации

## Структура проекта

```
RecruitAI/
├── back/                    # Backend приложение
│   ├── cmd/
│   │   └── back/
│   │       └── main.go      # Точка входа
│   ├── internal/
│   │   ├── app/
│   │   │   └── server.go    # HTTP сервер и роутинг
│   │   ├── config/
│   │   │   └── config.go    # Конфигурация приложения
│   │   ├── db/
│   │   │   └── db.go        # Инициализация БД
│   │   ├── endpoints/
│   │   │   └── candidate_resume/
│   │   │       ├── handler.go   # HTTP обработчики
│   │   │       ├── service.go   # Бизнес-логика
│   │   │       ├── repository.go # Репозиторий для БД
│   │   │       └── gpt.go        # Интеграция с OpenAI API
│   │   └── middlewares/
│   │       └── auth.go       # Middleware для аутентификации
│   ├── go.mod
│   └── go.sum
├── buildfiles/
│   └── back.dockerfile       # Dockerfile для сборки
├── data/
│   └── db/                   # Данные PostgreSQL
├── docker-compose.yml        # Docker Compose конфигурация
└── README.md
```

## Требования

- Docker и Docker Compose
- Go 1.24+ (для локальной разработки)
- OpenAI API ключ

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone <repository-url>
cd RecruitAI
```

### 2. Создание файла конфигурации

Создайте файл `.env` в корне проекта со следующим содержимым:

```env
# Порт приложения
BACK_PORT=8080

# Настройки PostgreSQL
BACK_POSTGRES_HOST=db
BACK_POSTGRES_PORT=5432
BACK_POSTGRES_USER=postgres
BACK_POSTGRES_PASSWORD=your_password
BACK_POSTGRES_DB=recruitai

# Токен для аутентификации API
BACK_SECRET_TOKEN=your_secret_token

# OpenAI API ключ
KEY_GPT=your_openai_api_key
```

### 3. Запуск через Docker Compose

```bash
docker-compose up -d
```

Приложение будет доступно по адресу `http://localhost:8080`

### 4. Локальная разработка (без Docker)

1. Убедитесь, что PostgreSQL запущен и доступен
2. Обновите переменные окружения в `.env` (особенно `BACK_POSTGRES_HOST=localhost`)
3. Установите зависимости:

```bash
cd back
go mod download
```

4. Запустите приложение:

```bash
go run cmd/back/main.go
```

## API

### Проверка соответствия кандидата

Оценивает соответствие резюме кандидата требованиям вакансии.

**Endpoint:** `POST /api/v1/resume`

**Аутентификация:** Требуется заголовок `Authorization: Bearer <BACK_SECRET_TOKEN>`

**Параметры запроса (multipart/form-data):**
- `resume` (file) — файл резюме (обязательно)
- `vacancy` (string) — описание вакансии (обязательно)

**Пример запроса:**

```bash
curl -X POST http://localhost:8080/api/v1/resume \
  -H "Authorization: Bearer your_secret_token" \
  -F "resume=@/path/to/resume.pdf" \
  -F "vacancy=Требуется разработчик Go с опытом работы от 3 лет..."
```

**Пример ответа:**

```json
{
  "match": true
}
```

**Описание:**
- Система загружает резюме в OpenAI API
- GPT-4.1 анализирует резюме и описание вакансии
- Оценка производится по шкале от 1 до 10
- Если оценка >= 7, возвращается `match: true`, иначе `match: false`

## Переменные окружения

| Переменная | Описание | Обязательно |
|------------|----------|-------------|
| `BACK_PORT` | Порт для HTTP сервера | Да |
| `BACK_POSTGRES_HOST` | Хост PostgreSQL | Да |
| `BACK_POSTGRES_PORT` | Порт PostgreSQL | Да |
| `BACK_POSTGRES_USER` | Пользователь PostgreSQL | Да |
| `BACK_POSTGRES_PASSWORD` | Пароль PostgreSQL | Да |
| `BACK_POSTGRES_DB` | Имя базы данных | Да |
| `BACK_SECRET_TOKEN` | Токен для аутентификации API | Да |
| `KEY_GPT` | OpenAI API ключ | Да |

## Разработка

### Структура кода

Проект следует чистой архитектуре с разделением на слои:
- **Handlers** — обработка HTTP запросов
- **Services** — бизнес-логика
- **Repositories** — работа с базой данных
- **Middlewares** — промежуточное ПО (аутентификация и т.д.)

### Добавление новых endpoints

1. Создайте новый пакет в `internal/endpoints/`
2. Реализуйте handler, service и repository
3. Зарегистрируйте роут в `internal/app/server.go`

## Логирование

Приложение использует `logrus` для логирования. Уровень логирования можно настроить через переменные окружения.

## Авторы

- Afanasev — разработка и поддержка


