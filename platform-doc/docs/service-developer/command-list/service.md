# Работа с сервисом

## service create

Создание нового сервиса.

```shell
platformctl service create service_name
```

## service start

Запуск сервиса в локальном окружении.

!!! bug

    В настоящий момент команда работает нестабильно.
    Вместо неё предполагается использование команды `service debug`.

```shell
platformctl service start
```

## service debug

Запуск зависимостей сервиса в локальном окружении. Сам сервис предполагается
запускать на хост-машине с помощью отладчика.

```shell
platformctl service debug
```

## service stop

Остановка сервиса в локальном окружении.

Команда останавливает и удаляет все зависимости текущего сервиса из локального
окружения.

```shell
platformctl service stop
```

## service log

Просмотр логов сервиса, запущенного в локальном окружении.

!!! note

    Эта команда не предназначена для просмотра логов сервисов,
    запущенных в режиме отладки.

```shell
platformctl service log
```

## service component

Список компонентов сервиса и платформы с адресами доступа.

```shell
platformctl service component
```

## service env

Вывод переменных окружения сервиса для просмотра и экспорта.

```shell
platformctl service env
platformctl service env --service-env local
platformctl service env --service-env staging
platformctl service env --service-env production

eval $(platformctl service env)
eval $(platformctl service env --service-env local)
eval $(platformctl service env --service-env staging)
eval $(platformctl service env --service-env production)
```

## service doc

Запуск документации сервиса для локального просмотра и редактирования.

!!! warning

    На данный момент поддерживается только документация в формате
    `mkdocs-material`.

После успешного запуска контейнера с документацией, она будет открыта в
браузере.

```shell
platformctl service doc
```

Документация должна быть размещена в каталоге `doc` в соответствии с шаблоном
сервиса (см. паттерн микросервисной архитектуры `Service Template`).
