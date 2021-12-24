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

## `platformctl`

### `platformctl service create`

Команда создания нового сервиса.

На первом этапе разработки:

* Создаёт папку с проектом.
* Создаёт манифест нового сервиса (`platform.yaml`).
* Создаёт точку входа в сервис (`cmd/service/main.go`).
* Инициализирует `go modules` сервиса.

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

* Сервис создан.
* В качестве http-сервера использован `labstack/echo`.
* Настроено эталонное логгирование.
    * В качестве `middleware`
      использован [brpaz/echozap](https://github.com/brpaz/echozap).
* Настроена эталонная отправка трейсов.
    * В качестве основы использован `jaegertracing`.
    * В качестве `middleware`
      использован https://echo.labstack.com/middleware/jaegertracing/

По мере развития проекта:

* Настроен `opentelemetry`.
