# Lolzteam Go API Client

[![CI](https://github.com/teracotaCode/lolzteam-go/actions/workflows/ci.yml/badge.svg)](https://github.com/teracotaCode/lolzteam-go/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/teracotaCode/lolzteam-go.svg)](https://pkg.go.dev/github.com/teracotaCode/lolzteam-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/teracotaCode/lolzteam-go)](https://goreportcard.com/report/github.com/teracotaCode/lolzteam-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Клиентская библиотека на Go для работы с API [Lolzteam](https://zelenka.guru) — Форум и Маркет.

## Установка

```bash
git clone https://github.com/teracotaCode/lolzteam-go.git
```

## Быстрый старт

### API Форума

```go
package main

import (
    "context"
    "fmt"
    "log"

    lolzteam "github.com/teracotaCode/lolzteam-go"
)

func main() {
    client, err := lolzteam.NewForumClient("ваш-api-токен")
    if err != nil {
        log.Fatal(err)
    }

    // Доступны: client.Threads, client.Posts, client.Users и др.
    _ = client
    fmt.Println("Клиент форума готов")
}
```

### API Маркета

```go
package main

import (
    "fmt"
    "log"
    "time"

    lolzteam "github.com/teracotaCode/lolzteam-go"
    "github.com/teracotaCode/lolzteam-go/runtime"
)

func main() {
    client, err := lolzteam.NewMarketClient("ваш-api-токен",
        lolzteam.WithTimeout(30*time.Second),
        lolzteam.WithProxy("socks5://proxy.example.com:1080"),
        lolzteam.WithRetry(runtime.RetryConfig{MaxRetries: 5}),
    )
    if err != nil {
        log.Fatal(err)
    }

    _ = client
    fmt.Println("Клиент маркета готов")
}
```

## Параметры конфигурации

| Опция | Описание |
|-------|----------|
| `WithTimeout(d)` | Таймаут HTTP-клиента (по умолчанию: 60 сек) |
| `WithProxy(url)` | URL прокси-сервера (http, https, socks5) |
| `WithRetry(cfg)` | Настройки повторных попыток |
| `WithRateLimit(cfg)` | Настройки ограничения частоты запросов |
| `WithBaseURL(url)` | Переопределение базового URL API |

## Возможности

- **Автоматическое ограничение частоты запросов** — Token bucket с плавным пополнением (300 запросов/мин общие, 20 запросов/мин поиск)
- **Повторные попытки с откатом** — Экспоненциальный откат с джиттером, поддержка заголовка Retry-After
- **Поддержка прокси** — HTTP, HTTPS и SOCKS5 прокси
- **Поддержка контекстов** — Все операции принимают `context.Context` для отмены
- **Потокобезопасность** — Безопасное использование из нескольких горутин одновременно
- **Типизированные ошибки** — Детализированные типы ошибок для HTTP-ошибок, сетевых ошибок, превышения лимитов и др.

## Обработка ошибок

```go
import (
    "errors"
    "github.com/teracotaCode/lolzteam-go/runtime"
)

_, err := client.SomeMethod(ctx)
if err != nil {
    var rateLimitErr *runtime.RateLimitError
    if errors.As(err, &rateLimitErr) {
        // Обработка превышения лимита запросов
    }

    var authErr *runtime.AuthError
    if errors.As(err, &authErr) {
        // Обработка ошибки аутентификации
    }
}
```

## Разработка

```bash

# Запуск линтера
make lint

# Сборка
make build

# Перегенерация API-клиентов из схем
make generate
```

## Лицензия

[MIT](LICENSE)
