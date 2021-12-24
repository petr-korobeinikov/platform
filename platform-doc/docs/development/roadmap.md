# Roadmap

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
