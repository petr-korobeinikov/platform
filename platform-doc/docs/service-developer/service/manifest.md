# Манифест сервиса

Манифест сервиса `platform.yaml` является основным источником информации о
сервисе.

Он находится в корне репозитория с сервисом и может быть изменен вручную
разработчиком.

## Пример манифеста `platform.yaml`

```yaml
name: wordcounter # (1)

environment: # (2)
  _:
    WORKER_BATCH_SIZE: 10
    WORKER_NAP_DURATION: 1s
  staging:
    WORKER_NAP_DURATION: 10s # (4)
  prod:
    WORKER_NAP_DURATION: 30s # (5)

component: # (3)
  - type: postgres
    name: postgres
    enabled: true

  - type: vault
    name: vault
    enabled: true

  - type: minio
    name: minio
    enabled: true
```

1. Имя сервиса
2. Спецификация переменных окружения
3. Список компонентов сервиса
4. Переопределение значения `WORKER_NAP_DURATION` для окружения `staging`.
5. Переопределение значения `WORKER_NAP_DURATION` для окружения `prod`.

## Структура манифеста

### name

Имя сервиса уникально в рамках инсталляции платформы.

```yaml
name: wordcounter
```

### environment

Определение переменных окружения сервиса.

Блок `_` определяет переменные окружения, доступные во всех средах.

Переменная окружения может быть переопределена для определённого окружения.

Согласно "12 факторам"[^1] вся конфигурация сервиса происходит только через
переменные окружения.

```yaml
environment:
  _:
    WORKER_BATCH_SIZE: 10
    WORKER_NAP_DURATION: 1s
  staging:
    WORKER_NAP_DURATION: 10s
  prod:
    WORKER_NAP_DURATION: 30s
```

### component

Компоненты сервиса, необходимые для его эксплуатации, например, базы данных,
брокеры очередей сообщений, системы кэширования и прочее.

Типы компонентов определяются платформой.

Имена компонентов определяются разработчиками сервиса. К сервису может быть
подключено несколько компонентов одного типа. Например, это могут быть `OLTP`-
и `OLAP`-база одного типа.

```yaml
component:
  - type: postgres
    name: postgres
    enabled: true
```

[^1]: [III. Конфигурация](https://12factor.net/ru/config)
