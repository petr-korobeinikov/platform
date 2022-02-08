# Roadmap

## `platform`

* `platform-service-registry`
    * Хранит в себе информацию о всех сервисах.
    * Предоставляет список `endpoint`-ов для сбора `prometheus`-метрик.
* `platform-service-asset-registry`
    * Хранит в себе необходимые артефакты для сборки сервисов:
        * `Dockerfile`
* Продуман механизм межсервисного взаимодействия:
    * Единое хранилище `proto`-файлов.
    * Генератор кода `service-code-generator`.
        * Название (?):
            * `service-protoc` :thinking:
            * `service-api-generator` :thinking:
            * `service-grpc-generator` :white_check_mark:
        * Генерирует фиксированной версией `protoc` код по спеке `grpc`.
* Общий линтер для всех сервисов с общей конфигурацией.
* Общий сканер безопасности с общей настройкой.
    * SAST
    * DAST
    * Материалы:
        * https://docs.gitlab.com/ee/development/integrations/secure.html
* Общий сканер лицензий зависимых библиотек.
* Вывод информации о сервисе:
    * `platformctl service info`

## `platformctl`

### `platformctl service create`

Команда создания нового сервиса.

На первом этапе разработки:

* :white_check_mark: Создаёт папку с проектом.
* :white_check_mark: Создаёт манифест нового сервиса (`platform.yaml`).
* :white_check_mark: Создаёт точку входа в сервис (`cmd/service/main.go`).
* :white_check_mark: Создаёт `.gitignore`.
* :white_check_mark: Инициализирует `go modules` сервиса.

По мере развития проекта:

* Регистрирует сервис в реестре сервисов.
* Создаёт репозиторий в системе контроля версий (`git`/`gitlab`,
  возможно, `github`).
* Создаёт сервис из шаблона (`boilerplate`).
* Учитывает выбранный язык программирования (`go`/`python`).

## `platform-showcase`

### `wordcounter`

...

### `reference-observable-service`

На первом этапе разработки:

* :white_check_mark: Сервис создан.
* :white_check_mark: В качестве http-сервера использован `labstack/echo`.
* :white_check_mark: Настроено эталонное логгирование.
    * :white_check_mark: В качестве `middleware`
      использован [brpaz/echozap](https://github.com/brpaz/echozap).
* :white_check_mark: Настроена эталонная отправка трейсов.
    * :white_check_mark: В качестве основы использован `jaegertracing`.
    * :white_check_mark: В качестве `middleware`
      использован https://echo.labstack.com/middleware/jaegertracing/.
* :white_check_mark: Продемонстрирована группировка ошибок в `sentry`.
* :white_check_mark: Продемонстрирован экспорт метрик `prometheus`.

По мере развития проекта:

* Настроен `opentelemetry`.
