# Библиотека задач

В платформе поддерживаются набор готовых задач для использования в сервисах.

!!! warning

    Библиотека задач находится в процессе разработки, отладки и тестирования.

    API вызова может измениться.
    Не забывайте сверяться с документацией.

## lint

Запуск линтера `golangci-lint` на исходном коде сервиса.

Определение в манифесте сервиса:

<!-- @formatter:off -->
```yaml
task:
  - name: lint
    image: platform/task/golangci-lint
```
<!-- @formatter:on -->

Пример вызова:

```shell
platformctl task lint
```

## db migrate master

Запуск миграций `golang-migrate` для главной базы данных сервиса.

Определение в манифесте сервиса:

<!-- @formatter:off -->
```yaml
task:
  - name: db migrate master
    image: platform/task/golang-migrate
    argument:
      workdir: /service/db/changelog/postgres/master
      command: >
        -path task-providing-service
        -database postgres://${PG_USER}:${PG_PASSWORD}@localhost:5432/${PG_DATABASE}?x-migrations-table="changelog"&x-migrations-table-quoted=true&sslmode=disable
        up
```
<!-- @formatter:on -->

!!! note

    Из-за топорности в обработке ошибок и негибкости интерфейса `golang-migrate`
    требуется определять аргументы команды.

    Переход на единообразную модель миграций во всех сервисах и смена мигратора
    на "нормальный" решит эту проблему.

Пример вызова:

```shell
platformctl task db migrate master
```
