# 🚀 Effective Mobile Subscription Service

**REST API сервис** для управления онлайн-подписками пользователей на Go с **gRPC** и **gRPC-Gateway**.

## 📋 Описание

Сервис предоставляет API для создания, чтения, обновления и удаления подписок пользователей на различные онлайн-сервисы. Поддерживает расчет общей стоимости подписок за выбранный период.

## 🛠️ Технологии

- **Go 1.24** + **gRPC** + **gRPC-Gateway**
- **PostgreSQL 16** + **pgx**
- **Docker Compose** для развертывания
- **Swagger UI** для документации
- **protoc-gen-validate** для валидации

## 🚀 Быстрый старт

### 1. Клонирование
```bash
git clone <repository>
cd effective-mobile-subscription-service
```

### 2. Запуск через Docker Compose
```bash
docker-compose up -d
```

Сервис запустится на:
- **HTTP API**: `http://localhost:8081`
- **gRPC**: `localhost:50052`
- **PostgreSQL**: `localhost:5432`

### 3. Документация API
- **Swagger UI**: http://localhost:8081/swagger-ui.html

## 🔌 HTTP API Endpoints

| Метод | Endpoint | Описание |
|-------|----------|----------|
| `POST` | `/api/v1/subscriptions` | ➕ Создание подписки |
| `GET` | `/api/v1/subscriptions/{user_id}` | 📋 Подписки пользователя |
| `GET` | `/api/v1/subscriptions` | 📋 Все подписки |
| `PATCH` | `/api/v1/subscriptions/{id}` | ✏️ Обновление подписки |
| `DELETE` | `/api/v1/subscriptions/{id}` | 🗑️ Удаление подписки |
| `GET` | `/api/v1/subscriptions/total-cost` | 💰 Расчет стоимости |

## 📊 Модель данных

```sql
subscriptions (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    service_name text NOT NULL,
    price bigint,
    start_date timestamp NOT NULL,
    end_date timestamp,
    deleted_at timestamp
)
```

## 💡 Примеры запросов

### Создание подписки
```bash
curl -X POST http://localhost:8081/api/v1/subscriptions \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
    "service_name": "Yandex Plus",
    "price": 400,
    "start_date": "2025-01-01T00:00:00Z"
  }'
```

### Расчет стоимости
```bash
curl "http://localhost:8081/api/v1/subscriptions/total-cost?start_date=2025-01-01T00:00:00Z&end_date=2025-12-31T23:59:59Z&user_ids=60601fee-2bf1-4721-ae6f-7636e79a0cba"
```

## 🔧 Конфигурация

Создайте `.env` файл:
```env
POSTGRES_DB=mydb
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_PORT=5432
HTTP_PORT=8081
GRPC_PORT=50052
```

## 🔒 Особенности

- ✅ **Валидация** через protoc-gen-validate
- ✅ **Мягкое удаление** записей
- ✅ **UUID** для идентификаторов
- ✅ **Connection pooling** для PostgreSQL
- ✅ **Graceful shutdown** сервера
- ✅ **Swagger документация** автоматически

---

**Примечание**: Сервис не проверяет существование пользователей, так как управление пользователями находится вне зоны ответственности.
